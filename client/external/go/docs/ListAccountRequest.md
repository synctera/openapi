# ListAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | Return accounts belonging to this customer ID | 
**HasDetails** | Pointer to **bool** | Should the account include information for relationships, aliases, balances and recent transactions | [optional] [default to false]

## Methods

### NewListAccountRequest

`func NewListAccountRequest(customerId string, ) *ListAccountRequest`

NewListAccountRequest instantiates a new ListAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccountRequestWithDefaults

`func NewListAccountRequestWithDefaults() *ListAccountRequest`

NewListAccountRequestWithDefaults instantiates a new ListAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *ListAccountRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ListAccountRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ListAccountRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetHasDetails

`func (o *ListAccountRequest) GetHasDetails() bool`

GetHasDetails returns the HasDetails field if non-nil, zero value otherwise.

### GetHasDetailsOk

`func (o *ListAccountRequest) GetHasDetailsOk() (*bool, bool)`

GetHasDetailsOk returns a tuple with the HasDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDetails

`func (o *ListAccountRequest) SetHasDetails(v bool)`

SetHasDetails sets HasDetails field to given value.

### HasHasDetails

`func (o *ListAccountRequest) HasHasDetails() bool`

HasHasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


