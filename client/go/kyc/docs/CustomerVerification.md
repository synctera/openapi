# CustomerVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationType** | [**[]VerificationType**](VerificationType.md) |  | 
**Ssn** | **string** | Customer&#39;s full tax ID eg SSN formatted with hyphens 123-45-6789 | 
**UserIpAddress** | Pointer to **string** | IP address | [optional] 

## Methods

### NewCustomerVerification

`func NewCustomerVerification(verificationType []VerificationType, ssn string, ) *CustomerVerification`

NewCustomerVerification instantiates a new CustomerVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerVerificationWithDefaults

`func NewCustomerVerificationWithDefaults() *CustomerVerification`

NewCustomerVerificationWithDefaults instantiates a new CustomerVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationType

`func (o *CustomerVerification) GetVerificationType() []VerificationType`

GetVerificationType returns the VerificationType field if non-nil, zero value otherwise.

### GetVerificationTypeOk

`func (o *CustomerVerification) GetVerificationTypeOk() (*[]VerificationType, bool)`

GetVerificationTypeOk returns a tuple with the VerificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationType

`func (o *CustomerVerification) SetVerificationType(v []VerificationType)`

SetVerificationType sets VerificationType field to given value.


### GetSsn

`func (o *CustomerVerification) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *CustomerVerification) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *CustomerVerification) SetSsn(v string)`

SetSsn sets Ssn field to given value.


### GetUserIpAddress

`func (o *CustomerVerification) GetUserIpAddress() string`

GetUserIpAddress returns the UserIpAddress field if non-nil, zero value otherwise.

### GetUserIpAddressOk

`func (o *CustomerVerification) GetUserIpAddressOk() (*string, bool)`

GetUserIpAddressOk returns a tuple with the UserIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIpAddress

`func (o *CustomerVerification) SetUserIpAddress(v string)`

SetUserIpAddress sets UserIpAddress field to given value.

### HasUserIpAddress

`func (o *CustomerVerification) HasUserIpAddress() bool`

HasUserIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


