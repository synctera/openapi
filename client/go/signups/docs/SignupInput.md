# SignupInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankName** | **string** | Bank Name | [readonly] 
**BankShortCode** | **string** | Bank Short Code | [readonly] 
**PartnerName** | **string** | Partner Name | [readonly] 
**UserId** | **string** | Bank&#39;s party&#39;s user_id | [readonly] 

## Methods

### NewSignupInput

`func NewSignupInput(bankName string, bankShortCode string, partnerName string, userId string, ) *SignupInput`

NewSignupInput instantiates a new SignupInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignupInputWithDefaults

`func NewSignupInputWithDefaults() *SignupInput`

NewSignupInputWithDefaults instantiates a new SignupInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankName

`func (o *SignupInput) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *SignupInput) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *SignupInput) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetBankShortCode

`func (o *SignupInput) GetBankShortCode() string`

GetBankShortCode returns the BankShortCode field if non-nil, zero value otherwise.

### GetBankShortCodeOk

`func (o *SignupInput) GetBankShortCodeOk() (*string, bool)`

GetBankShortCodeOk returns a tuple with the BankShortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankShortCode

`func (o *SignupInput) SetBankShortCode(v string)`

SetBankShortCode sets BankShortCode field to given value.


### GetPartnerName

`func (o *SignupInput) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *SignupInput) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *SignupInput) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.


### GetUserId

`func (o *SignupInput) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SignupInput) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SignupInput) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


