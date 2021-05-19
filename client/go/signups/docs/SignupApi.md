# \SignupApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSignup**](SignupApi.md#CreateSignup) | **Post** /signups | Signup a new sandbox user



## CreateSignup

> SignupIds CreateSignup(ctx).SignupInput(signupInput).Execute()

Signup a new sandbox user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    signupInput := *openapiclient.NewSignupInput("BankName_example", "BankShortCode_example", "PartnerName_example", "UserId_example") // SignupInput | Signup input data (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SignupApi.CreateSignup(context.Background()).SignupInput(signupInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignupApi.CreateSignup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSignup`: SignupIds
    fmt.Fprintf(os.Stdout, "Response from `SignupApi.CreateSignup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signupInput** | [**SignupInput**](SignupInput.md) | Signup input data | 

### Return type

[**SignupIds**](SignupIds.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

