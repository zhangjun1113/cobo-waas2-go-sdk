# TransactionCustodialAssetWalletSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**TransactionSourceType**](TransactionSourceType.md) |  | 
**WalletId** | **string** | The wallet ID. | 

## Methods

### NewTransactionCustodialAssetWalletSource

`func NewTransactionCustodialAssetWalletSource(sourceType TransactionSourceType, walletId string, ) *TransactionCustodialAssetWalletSource`

NewTransactionCustodialAssetWalletSource instantiates a new TransactionCustodialAssetWalletSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionCustodialAssetWalletSourceWithDefaults

`func NewTransactionCustodialAssetWalletSourceWithDefaults() *TransactionCustodialAssetWalletSource`

NewTransactionCustodialAssetWalletSourceWithDefaults instantiates a new TransactionCustodialAssetWalletSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *TransactionCustodialAssetWalletSource) GetSourceType() TransactionSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TransactionCustodialAssetWalletSource) GetSourceTypeOk() (*TransactionSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TransactionCustodialAssetWalletSource) SetSourceType(v TransactionSourceType)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *TransactionCustodialAssetWalletSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransactionCustodialAssetWalletSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransactionCustodialAssetWalletSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


