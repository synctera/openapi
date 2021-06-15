package synctera

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"
)

const sandbox = false

func createClient(t testing.TB, ctx context.Context) (*APIClient, context.Context) {
	configuration := NewConfiguration()
	//configuration.Debug = true

	if sandbox {
		ctx = context.WithValue(ctx, ContextServerIndex, 1)
		apiKey := os.Getenv("SYNCTERA_API_KEY")
		if len(apiKey) == 0 {
			t.Fatal("Missing SYNCTERA_API_KEY environment variable")
		}
		configuration.AddDefaultHeader("Authorization", "Bearer "+apiKey)
	} else {
		configuration.Servers = append(configuration.Servers, ServerConfiguration{URL: "http://localhost:5010/v0", Description: "local"})
		ctx = context.WithValue(ctx, ContextServerIndex, len(configuration.Servers)-1)
		jwt, err := exec.Command("sc", "jwt").Output()
		if err != nil {
			t.Fatal("Failed to get SyncTera JWT")
		}
		jwt = bytes.TrimSuffix(jwt, []byte("\n"))
		configuration.AddDefaultHeader("Synctera-Identity", string(jwt))
	}

	return NewAPIClient(configuration), ctx
}

func TestCreateCustomer(t *testing.T) {
	ctx := context.Background()
	client, ctx := createClient(t, ctx)

	const firstName = "GoClient"
	const lastName = "TestCreateCustomer"
	customerAddress := Address{
		Id:                nil,
		DefaultAddressFlg: true,
		Type:              PtrString("home"),
		AddressLine1:      "abc",
		AddressLine2:      nil,
		City:              "whoville",
		State:             "xx",
		PostalCode:        "abc123",
		CountryCode:       "US",
	}
	testCustomer := Customer{
		FirstName:       PtrString(firstName),
		LastName:        PtrString(lastName),
		LegalAddress:    &customerAddress,
		ShippingAddress: &customerAddress,
		Dob:             PtrString("1900-01-01"),
		PhoneNumber:     PtrString("+19178675309"),
	}
	createCustomerResponse, httpResponse, err := client.CustomersApi.CreateCustomer(ctx).Customer(testCustomer).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if httpResponse.StatusCode != http.StatusCreated {
		t.Error("Wrong status for create customer:", httpResponse.StatusCode)
	}

	if createCustomerResponse.GetFirstName() != firstName || createCustomerResponse.GetLastName() != lastName {
		t.Error("Wrong name for create customer")
	}

	getCustomerResponse, httpResponse, err := client.CustomersApi.GetCustomer(ctx, createCustomerResponse.GetId()).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Error("Wrong status for get customer:", httpResponse.StatusCode)
	}
	if getCustomerResponse.GetFirstName() != firstName || getCustomerResponse.GetLastName() != lastName {
		t.Error("Wrong name in list customers")
	}
}

func TestCreateDisclosure(t *testing.T) {
	ctx := context.Background()
	client, ctx := createClient(t, ctx)

	const disclosureType = "REG_DD"
	const disclosureVersion = "v1.0"
	disclosure := NewDisclosure(disclosureType, disclosureVersion, time.Now(), "VIEWED")

	customerID := uuid.New().String()
	result, httpResponse, err := client.DisclosuresApi.CreateDisclosure(ctx, customerID).Disclosure(*disclosure).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if httpResponse.StatusCode != http.StatusCreated {
		t.Error("Wrong status for create disclosure:", httpResponse.StatusCode)
	}
	if result.Type != disclosureType {
		t.Error("Wrong disclosure type")
	}
	if result.Version != disclosureVersion {
		t.Error("Wrong disclosure version")
	}

	list, httpResponse, err := client.DisclosuresApi.ListDisclosures(ctx, customerID).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Error("Wrong status for create customer:", httpResponse.StatusCode)
	}
	if list.Disclosures == nil || len(*list.Disclosures) == 0 {
		t.Fatal("Got empty list of disclosures")
	}
	lastDisclosure := (*list.Disclosures)[0]
	if lastDisclosure.Type != disclosureType {
		t.Error("Wrong disclosure type")
	}
	if lastDisclosure.Version != disclosureVersion {
		t.Error("Wrong disclosure version")
	}
}

func TestCreateCustomerVerification(t *testing.T) {
	ctx := context.Background()
	client, ctx := createClient(t, ctx)

	const firstName = "GoClient"
	const lastName = "TestCreateCustomerVerification"
	customerAddress := Address{
		Id:                nil,
		DefaultAddressFlg: true,
		Type:              PtrString("home"),
		AddressLine1:      "abc",
		AddressLine2:      nil,
		City:              "whoville",
		State:             "xx",
		PostalCode:        "abc123",
		CountryCode:       "US",
	}
	testCustomer := Customer{
		FirstName:         PtrString(firstName),
		LastName:          PtrString(lastName),
		LegalAddress:      &customerAddress,
		ShippingAddress:   &customerAddress,
		Dob:               PtrString("1900-01-01"),
		PhoneNumber: PtrString("+19178675309"),
	}

	customerResponse, httpResponse, err := client.CustomersApi.CreateCustomer(ctx).Customer(testCustomer).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if httpResponse.StatusCode != http.StatusCreated {
		t.Error("Wrong status for create customer:", httpResponse.StatusCode)
	}
	if customerResponse.Id == nil {
		t.Fatal("Missing constumer ID")
	}
	customerID := *customerResponse.Id

	const rawResponseData = `{"foo":"bar"}`
	year, month, day := time.Now().Date()
	customerVerificationResult := CustomerVerificationResult{
		VerificationsRun: []VerificationType{VERIFICATIONTYPE_KYC},
		Result:           "review",
		VerificationDate: fmt.Sprintf("%04d-%02d-%02d", year, month, day),
		RawResponse:      &RawResponse{Provider: PROVIDERTYPE_SOCURE.Ptr(), RawData: PtrString(rawResponseData)},
	}
	createResult, httpResponse, err := client.KYCVerificationApi.CreateCustomerVerificationResult(ctx, customerID).CustomerVerificationResult(customerVerificationResult).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if httpResponse.StatusCode != http.StatusCreated {
		t.Error("Wrong status for create customer verification:", httpResponse.StatusCode)
	}
	verificationID, hasVerificationID := createResult.GetIdOk()
	if !hasVerificationID {
		t.Error("ID missing")
	}

	getResult, httpResponse, err := client.KYCVerificationApi.GetVerification(ctx, customerID, *verificationID).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Error("Wrong status for create customer verification:", httpResponse.StatusCode)
	}
	if run := getResult.GetVerificationsRun(); len(run) != 1 || run[0] != VERIFICATIONTYPE_KYC {
		t.Error("Wrong verifications run")
	}
	rawResponse, rawResponseOk := getResult.GetRawResponseOk()
	if !rawResponseOk {
		t.Fatal("Missing raw response")
	}
	if rawResponse.GetProvider() != PROVIDERTYPE_SOCURE {
		t.Error("Wrong provider")
	}
	if rawResponse.GetRawData() != rawResponseData {
		t.Error("Wrong raw data")
	}
}

func TestCreateAccount(t *testing.T) {
	ctx := context.Background()
	client, ctx := createClient(t, ctx)

	const accountNumber = "123"
	newAccount := Account{
		AccountNumber: PtrString(accountNumber),
		Status:        STATUS_CLOSED.Ptr(),
	}
	account, httpResponse, err := client.AccountsApi.CreateAccount(ctx).Account(newAccount).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if httpResponse.StatusCode != http.StatusCreated {
		t.Error("Wrong status for create account:", httpResponse.StatusCode)
	}
	if !account.HasId() {
		t.Fatal("Account is missing ID")
	}
	accountID := account.GetId()

	account, httpResponse, err = client.AccountsApi.GetAccount(ctx, accountID).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Error("Wrong status for create account:", httpResponse.StatusCode)
	}
	if account.GetAccountNumber() != accountNumber {
		t.Error("Wrong account number")
	}
	if account.GetStatus() != STATUS_CLOSED {
		t.Error("Wrong account status")
	}
}
