package customers

import (
	"bytes"
	"context"
	"gitlab.com/synctera/openapi/client/go/customers"
	"net/http"
	"os"
	"os/exec"
	"testing"
)

const sandbox = false

func TestCreateAndListCustomer(t *testing.T) {
	configuration := customers.NewConfiguration()
	configuration.Debug = false
	ctx := context.Background()



	if sandbox {
		ctx = context.WithValue(ctx, customers.ContextServerIndex, 1)
		apiKey := os.Getenv("SYNCTERA_API_KEY")
		if len(apiKey) == 0 {
			t.Fatal("Missing SYNCTERA_API_KEY environment variable")
		}
		configuration.AddDefaultHeader("Authorization", "Bearer "+apiKey)
	} else {
		configuration.Servers = append(configuration.Servers, customers.ServerConfiguration{URL: "http://localhost:5010/v0", Description: "local"})
		ctx = context.WithValue(ctx, customers.ContextServerIndex, len(configuration.Servers)-1)
		jwt, err := exec.Command("sc", "jwt").Output()
		if err != nil {
			t.Fatal("Failed to get SyncTera JWT")
		}
		jwt = bytes.TrimSuffix(jwt, []byte("\n"))
		configuration.AddDefaultHeader("Synctera-Identity", string(jwt))
	}

	client := customers.NewAPIClient(configuration)

	customerAddress := customers.Address{
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
	customer := customers.NewCustomer(firstName, lastName, customerAddress, customerAddress, "1900-01-01", "", "+19178675309")
	customerResponse, httpResponse, err := client.CustomersApi.CreateCustomer(ctx).Customer(*customer).Execute()
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
	lastCustomer := customerList[len(customerList)-1].(map[string]interface{})
	if lastCustomer["first_name"] != firstName || lastCustomer["last_name"] != lastName {
		t.Error("Wrong name in list customers")
	}
}
