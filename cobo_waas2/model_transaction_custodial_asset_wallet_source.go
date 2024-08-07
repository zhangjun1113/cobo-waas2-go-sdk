/*
Cobo Wallet as a Service 2.0

API version: 1.1.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TransactionCustodialAssetWalletSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionCustodialAssetWalletSource{}

// TransactionCustodialAssetWalletSource Information about the transaction source type `Asset`. 
type TransactionCustodialAssetWalletSource struct {
	SourceType TransactionSourceType `json:"source_type"`
	// The wallet ID.
	WalletId string `json:"wallet_id"`
}

type _TransactionCustodialAssetWalletSource TransactionCustodialAssetWalletSource

// NewTransactionCustodialAssetWalletSource instantiates a new TransactionCustodialAssetWalletSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionCustodialAssetWalletSource(sourceType TransactionSourceType, walletId string) *TransactionCustodialAssetWalletSource {
	this := TransactionCustodialAssetWalletSource{}
	this.SourceType = sourceType
	this.WalletId = walletId
	return &this
}

// NewTransactionCustodialAssetWalletSourceWithDefaults instantiates a new TransactionCustodialAssetWalletSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionCustodialAssetWalletSourceWithDefaults() *TransactionCustodialAssetWalletSource {
	this := TransactionCustodialAssetWalletSource{}
	return &this
}

// GetSourceType returns the SourceType field value
func (o *TransactionCustodialAssetWalletSource) GetSourceType() TransactionSourceType {
	if o == nil {
		var ret TransactionSourceType
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *TransactionCustodialAssetWalletSource) GetSourceTypeOk() (*TransactionSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *TransactionCustodialAssetWalletSource) SetSourceType(v TransactionSourceType) {
	o.SourceType = v
}

// GetWalletId returns the WalletId field value
func (o *TransactionCustodialAssetWalletSource) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *TransactionCustodialAssetWalletSource) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *TransactionCustodialAssetWalletSource) SetWalletId(v string) {
	o.WalletId = v
}

func (o TransactionCustodialAssetWalletSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionCustodialAssetWalletSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_type"] = o.SourceType
	toSerialize["wallet_id"] = o.WalletId
	return toSerialize, nil
}

func (o *TransactionCustodialAssetWalletSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source_type",
		"wallet_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTransactionCustodialAssetWalletSource := _TransactionCustodialAssetWalletSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionCustodialAssetWalletSource)

	if err != nil {
		return err
	}

	*o = TransactionCustodialAssetWalletSource(varTransactionCustodialAssetWalletSource)

	return err
}

type NullableTransactionCustodialAssetWalletSource struct {
	value *TransactionCustodialAssetWalletSource
	isSet bool
}

func (v NullableTransactionCustodialAssetWalletSource) Get() *TransactionCustodialAssetWalletSource {
	return v.value
}

func (v *NullableTransactionCustodialAssetWalletSource) Set(val *TransactionCustodialAssetWalletSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionCustodialAssetWalletSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionCustodialAssetWalletSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionCustodialAssetWalletSource(val *TransactionCustodialAssetWalletSource) *NullableTransactionCustodialAssetWalletSource {
	return &NullableTransactionCustodialAssetWalletSource{value: val, isSet: true}
}

func (v NullableTransactionCustodialAssetWalletSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionCustodialAssetWalletSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


