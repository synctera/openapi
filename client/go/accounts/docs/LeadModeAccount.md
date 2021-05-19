# LeadModeAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Account ID | [optional] [readonly] 
**AccountNumber** | Pointer to **string** | Account number | [optional] [readonly] 
**Currency** | Pointer to **string** | Account currency or account settlement currency. ISO 4217 alphabetic currency code. Default USD | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**ExchangeRateType** | Pointer to **string** | Exchange rate type | [optional] 
**Iban** | Pointer to **string** | International bank account number | [optional] 
**SwiftCode** | Pointer to **string** | SWIFT code | [optional] 
**IsAccountPool** | Pointer to **bool** | Account is investment (variable balance) account or a multi-balance account pool. Default false | [optional] 
**AccountTemplateCode** | Pointer to **string** | Account template code | [optional] 
**AccountTemplateVersion** | Pointer to **float32** | Account template version | [optional] 
**Relationships** | [**[]Relationship**](Relationship.md) | List of the relationship for this account to the parties | 
**Aliases** | Pointer to [**[]Alias**](Alias.md) | A list of the aliases for account. Account alias is the account number of different balance types to link to the same account ID | [optional] 
**Balances** | Pointer to [**[]Balance**](Balance.md) | A list of balances for account based on different type | [optional] [readonly] 
**RecentTransactions** | Pointer to [**[]Transaction**](Transaction.md) | The most recent 10 transactions of the account | [optional] [readonly] 

## Methods

### NewLeadModeAccount

`func NewLeadModeAccount(relationships []Relationship, ) *LeadModeAccount`

NewLeadModeAccount instantiates a new LeadModeAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeadModeAccountWithDefaults

`func NewLeadModeAccountWithDefaults() *LeadModeAccount`

NewLeadModeAccountWithDefaults instantiates a new LeadModeAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LeadModeAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LeadModeAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LeadModeAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LeadModeAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountNumber

`func (o *LeadModeAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *LeadModeAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *LeadModeAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *LeadModeAccount) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetCurrency

`func (o *LeadModeAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *LeadModeAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *LeadModeAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *LeadModeAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetStatus

`func (o *LeadModeAccount) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LeadModeAccount) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LeadModeAccount) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LeadModeAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExchangeRateType

`func (o *LeadModeAccount) GetExchangeRateType() string`

GetExchangeRateType returns the ExchangeRateType field if non-nil, zero value otherwise.

### GetExchangeRateTypeOk

`func (o *LeadModeAccount) GetExchangeRateTypeOk() (*string, bool)`

GetExchangeRateTypeOk returns a tuple with the ExchangeRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateType

`func (o *LeadModeAccount) SetExchangeRateType(v string)`

SetExchangeRateType sets ExchangeRateType field to given value.

### HasExchangeRateType

`func (o *LeadModeAccount) HasExchangeRateType() bool`

HasExchangeRateType returns a boolean if a field has been set.

### GetIban

`func (o *LeadModeAccount) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *LeadModeAccount) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *LeadModeAccount) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *LeadModeAccount) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetSwiftCode

`func (o *LeadModeAccount) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *LeadModeAccount) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *LeadModeAccount) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.

### HasSwiftCode

`func (o *LeadModeAccount) HasSwiftCode() bool`

HasSwiftCode returns a boolean if a field has been set.

### GetIsAccountPool

`func (o *LeadModeAccount) GetIsAccountPool() bool`

GetIsAccountPool returns the IsAccountPool field if non-nil, zero value otherwise.

### GetIsAccountPoolOk

`func (o *LeadModeAccount) GetIsAccountPoolOk() (*bool, bool)`

GetIsAccountPoolOk returns a tuple with the IsAccountPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountPool

`func (o *LeadModeAccount) SetIsAccountPool(v bool)`

SetIsAccountPool sets IsAccountPool field to given value.

### HasIsAccountPool

`func (o *LeadModeAccount) HasIsAccountPool() bool`

HasIsAccountPool returns a boolean if a field has been set.

### GetAccountTemplateCode

`func (o *LeadModeAccount) GetAccountTemplateCode() string`

GetAccountTemplateCode returns the AccountTemplateCode field if non-nil, zero value otherwise.

### GetAccountTemplateCodeOk

`func (o *LeadModeAccount) GetAccountTemplateCodeOk() (*string, bool)`

GetAccountTemplateCodeOk returns a tuple with the AccountTemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTemplateCode

`func (o *LeadModeAccount) SetAccountTemplateCode(v string)`

SetAccountTemplateCode sets AccountTemplateCode field to given value.

### HasAccountTemplateCode

`func (o *LeadModeAccount) HasAccountTemplateCode() bool`

HasAccountTemplateCode returns a boolean if a field has been set.

### GetAccountTemplateVersion

`func (o *LeadModeAccount) GetAccountTemplateVersion() float32`

GetAccountTemplateVersion returns the AccountTemplateVersion field if non-nil, zero value otherwise.

### GetAccountTemplateVersionOk

`func (o *LeadModeAccount) GetAccountTemplateVersionOk() (*float32, bool)`

GetAccountTemplateVersionOk returns a tuple with the AccountTemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTemplateVersion

`func (o *LeadModeAccount) SetAccountTemplateVersion(v float32)`

SetAccountTemplateVersion sets AccountTemplateVersion field to given value.

### HasAccountTemplateVersion

`func (o *LeadModeAccount) HasAccountTemplateVersion() bool`

HasAccountTemplateVersion returns a boolean if a field has been set.

### GetRelationships

`func (o *LeadModeAccount) GetRelationships() []Relationship`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *LeadModeAccount) GetRelationshipsOk() (*[]Relationship, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *LeadModeAccount) SetRelationships(v []Relationship)`

SetRelationships sets Relationships field to given value.


### GetAliases

`func (o *LeadModeAccount) GetAliases() []Alias`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *LeadModeAccount) GetAliasesOk() (*[]Alias, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *LeadModeAccount) SetAliases(v []Alias)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *LeadModeAccount) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetBalances

`func (o *LeadModeAccount) GetBalances() []Balance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *LeadModeAccount) GetBalancesOk() (*[]Balance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *LeadModeAccount) SetBalances(v []Balance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *LeadModeAccount) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetRecentTransactions

`func (o *LeadModeAccount) GetRecentTransactions() []Transaction`

GetRecentTransactions returns the RecentTransactions field if non-nil, zero value otherwise.

### GetRecentTransactionsOk

`func (o *LeadModeAccount) GetRecentTransactionsOk() (*[]Transaction, bool)`

GetRecentTransactionsOk returns a tuple with the RecentTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentTransactions

`func (o *LeadModeAccount) SetRecentTransactions(v []Transaction)`

SetRecentTransactions sets RecentTransactions field to given value.

### HasRecentTransactions

`func (o *LeadModeAccount) HasRecentTransactions() bool`

HasRecentTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


