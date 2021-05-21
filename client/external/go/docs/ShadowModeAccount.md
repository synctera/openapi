# ShadowModeAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Account ID | [optional] [readonly] 
**AccountNumber** | **string** | Account number | 
**Currency** | Pointer to **string** | Account currency or account settlement currency. ISO 4217 alphabetic currency code. Default USD | [optional] 
**Status** | [**Status**](Status.md) |  | 
**IsAccountPool** | Pointer to **bool** | Account is investment (variable balance) account or a multi-balance account pool. Default false | [optional] 
**Relationships** | Pointer to [**[]SchemasRelationship**](SchemasRelationship.md) | List of the relationship for this account to the parties | [optional] 
**Aliases** | Pointer to [**[]Alias**](Alias.md) | A list of aliases for account. Account alias is the account number of different balance types to link to the same account ID | [optional] 
**Balances** | Pointer to [**[]Balance**](Balance.md) | A list of balances for account based on different type | [optional] [readonly] 
**RecentTransactions** | Pointer to [**[]Transaction**](Transaction.md) | The most recent 10 transactions of the account | [optional] [readonly] 

## Methods

### NewShadowModeAccount

`func NewShadowModeAccount(accountNumber string, status Status, ) *ShadowModeAccount`

NewShadowModeAccount instantiates a new ShadowModeAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShadowModeAccountWithDefaults

`func NewShadowModeAccountWithDefaults() *ShadowModeAccount`

NewShadowModeAccountWithDefaults instantiates a new ShadowModeAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShadowModeAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShadowModeAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShadowModeAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShadowModeAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountNumber

`func (o *ShadowModeAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *ShadowModeAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *ShadowModeAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetCurrency

`func (o *ShadowModeAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ShadowModeAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ShadowModeAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ShadowModeAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetStatus

`func (o *ShadowModeAccount) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShadowModeAccount) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShadowModeAccount) SetStatus(v Status)`

SetStatus sets Status field to given value.


### GetIsAccountPool

`func (o *ShadowModeAccount) GetIsAccountPool() bool`

GetIsAccountPool returns the IsAccountPool field if non-nil, zero value otherwise.

### GetIsAccountPoolOk

`func (o *ShadowModeAccount) GetIsAccountPoolOk() (*bool, bool)`

GetIsAccountPoolOk returns a tuple with the IsAccountPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountPool

`func (o *ShadowModeAccount) SetIsAccountPool(v bool)`

SetIsAccountPool sets IsAccountPool field to given value.

### HasIsAccountPool

`func (o *ShadowModeAccount) HasIsAccountPool() bool`

HasIsAccountPool returns a boolean if a field has been set.

### GetRelationships

`func (o *ShadowModeAccount) GetRelationships() []SchemasRelationship`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ShadowModeAccount) GetRelationshipsOk() (*[]SchemasRelationship, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ShadowModeAccount) SetRelationships(v []SchemasRelationship)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ShadowModeAccount) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetAliases

`func (o *ShadowModeAccount) GetAliases() []Alias`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *ShadowModeAccount) GetAliasesOk() (*[]Alias, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *ShadowModeAccount) SetAliases(v []Alias)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *ShadowModeAccount) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetBalances

`func (o *ShadowModeAccount) GetBalances() []Balance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *ShadowModeAccount) GetBalancesOk() (*[]Balance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *ShadowModeAccount) SetBalances(v []Balance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *ShadowModeAccount) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetRecentTransactions

`func (o *ShadowModeAccount) GetRecentTransactions() []Transaction`

GetRecentTransactions returns the RecentTransactions field if non-nil, zero value otherwise.

### GetRecentTransactionsOk

`func (o *ShadowModeAccount) GetRecentTransactionsOk() (*[]Transaction, bool)`

GetRecentTransactionsOk returns a tuple with the RecentTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentTransactions

`func (o *ShadowModeAccount) SetRecentTransactions(v []Transaction)`

SetRecentTransactions sets RecentTransactions field to given value.

### HasRecentTransactions

`func (o *ShadowModeAccount) HasRecentTransactions() bool`

HasRecentTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


