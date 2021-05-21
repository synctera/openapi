package synctera

import (
	"bytes"
	"context"
	"net/http"
	"os"
	"os/exec"
	"testing"
)

const sandbox = false

func createConfiguration(t testing.TB, ctx context.Context) (*Configuration, context.Context) {
	configuration := NewConfiguration()

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

	return configuration, ctx
}

func TestCreateAndListCustomer(t *testing.T) {
	ctx := context.Background()
	customersConfiguration, ctx := createConfiguration(t, ctx)
	customersClient := NewAPIClient(customersConfiguration)

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
	customerResponse, httpResponse, err := customersClient.CustomersApi.CreateCustomer(ctx).Customer(*testCustomer).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if httpResponse.StatusCode != http.StatusCreated {
		t.Error("Wrong status for create customer:", httpResponse.StatusCode)
	}

	if customerResponse.FirstName != firstName || customerResponse.LastName != lastName {
		t.Error("Wrong name for create customer")
	}

	list, httpResponse, err := customersClient.CustomersApi.ListCustomers(ctx).Execute()
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

	//newCustomerID := newCustomer["id"].(string)
	//
	//var disclosuresConfiguration disclosures.Configuration
	//copier.Copy(disclosuresConfiguration, customersConfiguration)
	//disclosuresClient := disclosures.NewAPIClient(&disclosuresConfiguration)
	//
	//disclosure, httpResponse, err := disclosuresClient.DisclosuresApi.CreateDisclosure(ctx, newCustomerID).Execute()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log(httpResponse.StatusCode)
	//t.Log(disclosure.Id)

}
