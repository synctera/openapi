# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | Return accounts belonging to this customer ID | [optional] 
**HasDetails** | Pointer to **bool** | Should the account include information for relationships, aliases, balances and recent transactions | [optional] 

## Methods

### NewInlineObject

`func NewInlineObject() *InlineObject`

NewInlineObject instantiates a new InlineObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObjectWithDefaults

`func NewInlineObjectWithDefaults() *InlineObject`

NewInlineObjectWithDefaults instantiates a new InlineObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *InlineObject) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InlineObject) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InlineObject) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *InlineObject) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetHasDetails

`func (o *InlineObject) GetHasDetails() bool`

GetHasDetails returns the HasDetails field if non-nil, zero value otherwise.

### GetHasDetailsOk

`func (o *InlineObject) GetHasDetailsOk() (*bool, bool)`

GetHasDetailsOk returns a tuple with the HasDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDetails

`func (o *InlineObject) SetHasDetails(v bool)`

SetHasDetails sets HasDetails field to given value.

### HasHasDetails

`func (o *InlineObject) HasHasDetails() bool`

HasHasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


