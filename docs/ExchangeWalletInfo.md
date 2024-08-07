# ExchangeWalletInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The wallet ID. | 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**Name** | **string** | The wallet name. | 
**OrgId** | **string** | The ID of the owning organization. | 
**Apikey** | **string** | The API key of your exchange account. | 
**ExchangeId** | [**ExchangeId**](ExchangeId.md) |  | 
**MainWalletId** | Pointer to **string** | The wallet ID of the Main Account associated with the Sub Account. This property is returned only if you are creating or querying an Exchange Wallet (Sub Account). | [optional] 

## Methods

### NewExchangeWalletInfo

`func NewExchangeWalletInfo(walletId string, walletType WalletType, walletSubtype WalletSubtype, name string, orgId string, apikey string, exchangeId ExchangeId, ) *ExchangeWalletInfo`

NewExchangeWalletInfo instantiates a new ExchangeWalletInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeWalletInfoWithDefaults

`func NewExchangeWalletInfoWithDefaults() *ExchangeWalletInfo`

NewExchangeWalletInfoWithDefaults instantiates a new ExchangeWalletInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *ExchangeWalletInfo) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *ExchangeWalletInfo) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *ExchangeWalletInfo) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetWalletType

`func (o *ExchangeWalletInfo) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *ExchangeWalletInfo) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *ExchangeWalletInfo) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *ExchangeWalletInfo) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *ExchangeWalletInfo) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *ExchangeWalletInfo) SetWalletSubtype(v WalletSubtype)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetName

`func (o *ExchangeWalletInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExchangeWalletInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExchangeWalletInfo) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *ExchangeWalletInfo) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ExchangeWalletInfo) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ExchangeWalletInfo) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetApikey

`func (o *ExchangeWalletInfo) GetApikey() string`

GetApikey returns the Apikey field if non-nil, zero value otherwise.

### GetApikeyOk

`func (o *ExchangeWalletInfo) GetApikeyOk() (*string, bool)`

GetApikeyOk returns a tuple with the Apikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApikey

`func (o *ExchangeWalletInfo) SetApikey(v string)`

SetApikey sets Apikey field to given value.


### GetExchangeId

`func (o *ExchangeWalletInfo) GetExchangeId() ExchangeId`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *ExchangeWalletInfo) GetExchangeIdOk() (*ExchangeId, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *ExchangeWalletInfo) SetExchangeId(v ExchangeId)`

SetExchangeId sets ExchangeId field to given value.


### GetMainWalletId

`func (o *ExchangeWalletInfo) GetMainWalletId() string`

GetMainWalletId returns the MainWalletId field if non-nil, zero value otherwise.

### GetMainWalletIdOk

`func (o *ExchangeWalletInfo) GetMainWalletIdOk() (*string, bool)`

GetMainWalletIdOk returns a tuple with the MainWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainWalletId

`func (o *ExchangeWalletInfo) SetMainWalletId(v string)`

SetMainWalletId sets MainWalletId field to given value.

### HasMainWalletId

`func (o *ExchangeWalletInfo) HasMainWalletId() bool`

HasMainWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


