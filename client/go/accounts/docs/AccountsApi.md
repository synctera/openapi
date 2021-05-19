# \AccountsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](AccountsApi.md#CreateAccount) | **Post** /accounts | Create an account
[**CreateAccountAddress**](AccountsApi.md#CreateAccountAddress) | **Post** /accounts/{account_id}/addresses | Create address for an account
[**DeleteAccount**](AccountsApi.md#DeleteAccount) | **Delete** /accounts/{account_id} | Delete account
[**DeleteAccountAddress**](AccountsApi.md#DeleteAccountAddress) | **Delete** /accounts/{account_id}/addresses/{connect_id} | Delete address for an account
[**GetAccount**](AccountsApi.md#GetAccount) | **Get** /accounts/{account_id} | Get account
[**GetAccountAddress**](AccountsApi.md#GetAccountAddress) | **Get** /accounts/{account_id}/addresses/{connect_id} | Get address for an account
[**GetAccountBalance**](AccountsApi.md#GetAccountBalance) | **Get** /accounts/{account_id}/balance | Get account balance
[**GetAccountTransactions**](AccountsApi.md#GetAccountTransactions) | **Get** /accounts/{account_id}/transactions | Get account transactions
[**ListAccounts**](AccountsApi.md#ListAccounts) | **Get** /accounts | List accounts
[**UpdateAccount**](AccountsApi.md#UpdateAccount) | **Put** /accounts/{account_id} | Update account
[**UpdateAccountAddress**](AccountsApi.md#UpdateAccountAddress) | **Put** /accounts/{account_id}/addresses/{connect_id} | Update address for an account



## CreateAccount

> Account CreateAccount(ctx).Account(account).Execute()

Create an account



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
    account := openapiclient.account{LeadModeAccount: openapiclient.NewLeadModeAccount([]openapiclient.Relationship{*openapiclient.NewRelationship()})} // Account | Account to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.CreateAccount(context.Background()).Account(account).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.CreateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | [**Account**](Account.md) | Account to create | 

### Return type

[**Account**](Account.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountAddress

> AccountAddress CreateAccountAddress(ctx, accountId).AccountAddress(accountAddress).Execute()

Create address for an account



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
    accountId := TODO // string | Account ID
    accountAddress := *openapiclient.NewAccountAddress() // AccountAddress | Account address to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.CreateAccountAddress(context.Background(), accountId).AccountAddress(accountAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.CreateAccountAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccountAddress`: AccountAddress
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.CreateAccountAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountAddress** | [**AccountAddress**](AccountAddress.md) | Account address to create | 

### Return type

[**AccountAddress**](AccountAddress.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccount

> DeleteResponse DeleteAccount(ctx, accountId).Execute()

Delete account



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
    accountId := TODO // string | Account ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.DeleteAccount(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccount`: DeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountAddress

> DeleteResponse DeleteAccountAddress(ctx, accountId, connectId).Execute()

Delete address for an account



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
    accountId := TODO // string | Account ID
    connectId := TODO // string | Connect ID of the account associate with the account entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.DeleteAccountAddress(context.Background(), accountId, connectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DeleteAccountAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccountAddress`: DeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.DeleteAccountAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 
**connectId** | [**string**](.md) | Connect ID of the account associate with the account entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> Account GetAccount(ctx, accountId).Execute()

Get account



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
    accountId := TODO // string | Account ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetAccount(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAddress

> AccountAddress GetAccountAddress(ctx, accountId, connectId).InlineObject3(inlineObject3).Execute()

Get address for an account



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
    accountId := TODO // string | Account ID
    connectId := TODO // string | Connect ID of the account associate with the account entity
    inlineObject3 := *openapiclient.NewInlineObject3() // InlineObject3 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetAccountAddress(context.Background(), accountId, connectId).InlineObject3(inlineObject3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountAddress`: AccountAddress
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 
**connectId** | [**string**](.md) | Connect ID of the account associate with the account entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject3** | [**InlineObject3**](InlineObject3.md) |  | 

### Return type

[**AccountAddress**](AccountAddress.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountBalance

> Balance GetAccountBalance(ctx, accountId).InlineObject1(inlineObject1).Execute()

Get account balance



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
    accountId := TODO // string | Account ID
    inlineObject1 := *openapiclient.NewInlineObject1() // InlineObject1 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetAccountBalance(context.Background(), accountId).InlineObject1(inlineObject1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountBalance`: Balance
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject1** | [**InlineObject1**](InlineObject1.md) |  | 

### Return type

[**Balance**](Balance.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountTransactions

> map[string]interface{} GetAccountTransactions(ctx, accountId).InlineObject2(inlineObject2).Limit(limit).PageToken(pageToken).Execute()

Get account transactions



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
    accountId := TODO // string | Account ID
    inlineObject2 := *openapiclient.NewInlineObject2() // InlineObject2 | 
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "faker.random.alphaNumeric(10)" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetAccountTransactions(context.Background(), accountId).InlineObject2(inlineObject2).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountTransactions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject2** | [**InlineObject2**](InlineObject2.md) |  | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> map[string]interface{} ListAccounts(ctx).InlineObject(inlineObject).Limit(limit).PageToken(pageToken).Execute()

List accounts



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
    inlineObject := *openapiclient.NewInlineObject() // InlineObject | 
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "faker.random.alphaNumeric(10)" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ListAccounts(context.Background()).InlineObject(inlineObject).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccounts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ListAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> Account UpdateAccount(ctx, accountId).Account(account).Execute()

Update account



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
    accountId := TODO // string | Account ID
    account := openapiclient.account{LeadModeAccount: openapiclient.NewLeadModeAccount([]openapiclient.Relationship{*openapiclient.NewRelationship()})} // Account | Account to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.UpdateAccount(context.Background(), accountId).Account(account).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **account** | [**Account**](Account.md) | Account to update | 

### Return type

[**Account**](Account.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountAddress

> AccountAddress UpdateAccountAddress(ctx, accountId, connectId).AccountAddress(accountAddress).Execute()

Update address for an account



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
    accountId := TODO // string | Account ID
    connectId := TODO // string | Connect ID of the account associate with the account entity
    accountAddress := *openapiclient.NewAccountAddress() // AccountAddress | Account address to be updated

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.UpdateAccountAddress(context.Background(), accountId, connectId).AccountAddress(accountAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.UpdateAccountAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountAddress`: AccountAddress
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.UpdateAccountAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 
**connectId** | [**string**](.md) | Connect ID of the account associate with the account entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountAddress** | [**AccountAddress**](AccountAddress.md) | Account address to be updated | 

### Return type

[**AccountAddress**](AccountAddress.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

