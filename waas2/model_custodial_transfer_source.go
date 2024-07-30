/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CustodialTransferSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustodialTransferSource{}

// CustodialTransferSource The information about the transaction source.
type CustodialTransferSource struct {
	SourceType WalletSubtype `json:"source_type"`
	// The wallet ID.
	WalletId string `json:"wallet_id"`
}

type _CustodialTransferSource CustodialTransferSource

// NewCustodialTransferSource instantiates a new CustodialTransferSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustodialTransferSource(sourceType WalletSubtype, walletId string) *CustodialTransferSource {
	this := CustodialTransferSource{}
	this.SourceType = sourceType
	this.WalletId = walletId
	return &this
}

// NewCustodialTransferSourceWithDefaults instantiates a new CustodialTransferSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustodialTransferSourceWithDefaults() *CustodialTransferSource {
	this := CustodialTransferSource{}
	return &this
}

// GetSourceType returns the SourceType field value
func (o *CustodialTransferSource) GetSourceType() WalletSubtype {
	if o == nil {
		var ret WalletSubtype
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *CustodialTransferSource) GetSourceTypeOk() (*WalletSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *CustodialTransferSource) SetSourceType(v WalletSubtype) {
	o.SourceType = v
}

// GetWalletId returns the WalletId field value
func (o *CustodialTransferSource) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *CustodialTransferSource) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *CustodialTransferSource) SetWalletId(v string) {
	o.WalletId = v
}

func (o CustodialTransferSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustodialTransferSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_type"] = o.SourceType
	toSerialize["wallet_id"] = o.WalletId
	return toSerialize, nil
}

func (o *CustodialTransferSource) UnmarshalJSON(data []byte) (err error) {
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

	varCustodialTransferSource := _CustodialTransferSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustodialTransferSource)

	if err != nil {
		return err
	}

	*o = CustodialTransferSource(varCustodialTransferSource)

	return err
}

type NullableCustodialTransferSource struct {
	value *CustodialTransferSource
	isSet bool
}

func (v NullableCustodialTransferSource) Get() *CustodialTransferSource {
	return v.value
}

func (v *NullableCustodialTransferSource) Set(val *CustodialTransferSource) {
	v.value = val
	v.isSet = true
}

func (v NullableCustodialTransferSource) IsSet() bool {
	return v.isSet
}

func (v *NullableCustodialTransferSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustodialTransferSource(val *CustodialTransferSource) *NullableCustodialTransferSource {
	return &NullableCustodialTransferSource{value: val, isSet: true}
}

func (v NullableCustodialTransferSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustodialTransferSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

