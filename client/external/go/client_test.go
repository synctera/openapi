package synctera

import (
	"bytes"
	"context"
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

	customerAddress := Address{
		Id:                nil,
		DefaultAddressFlg: true,
		Type:              "home",
		AddressLine1:      "abc",
		AddressLine2:      nil,
		City:              "whoville",
		State:             "xx",
		PostalCode:        "abc123",
		CountryCode:       "US",
	}
	const firstName = "GoClient"
	const lastName = "TestCustomer"
	testCustomer := NewCustomer(firstName, lastName, customerAddress, customerAddress, "1900-01-01", "", "+19178675309")
	customerResponse, httpResponse, err := client.CustomersApi.CreateCustomer(ctx).Customer(*testCustomer).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if httpResponse.StatusCode != http.StatusCreated {
		t.Error("Wrong status for create customer:", httpResponse.StatusCode)
	}

	if customerResponse.FirstName != firstName || customerResponse.LastName != lastName {
		t.Error("Wrong name for create customer")
	}

	list, httpResponse, err := client.CustomersApi.ListCustomers(ctx).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Error("Wrong status for create customer:", httpResponse.StatusCode)
	}
	customerList := list["customers"].([]interface{})
	newCustomer := customerList[len(customerList)-1].(map[string]interface{})
	if newCustomer["first_name"] != firstName || newCustomer["last_name"] != lastName {
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

	list, httpResponse, err  := client.DisclosuresApi.ListDisclosures(ctx, customerID).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Error("Wrong status for create customer:", httpResponse.StatusCode)
	}
	if list.Disclosures == nil || len(*list.Disclosures) == 0{
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
