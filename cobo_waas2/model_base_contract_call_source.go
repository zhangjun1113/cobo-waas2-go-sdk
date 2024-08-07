/*
Cobo Wallet as a Service 2.0

Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BaseContractCallSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseContractCallSource{}

// BaseContractCallSource The information about the transaction source type `Org-Controlled` and `User-Controlled`.
type BaseContractCallSource struct {
	SourceType ContractCallSourceType `json:"source_type"`
	// The wallet ID.
	WalletId string `json:"wallet_id"`
	// The wallet address.
	Address string `json:"address"`
}

type _BaseContractCallSource BaseContractCallSource

// NewBaseContractCallSource instantiates a new BaseContractCallSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseContractCallSource(sourceType ContractCallSourceType, walletId string, address string) *BaseContractCallSource {
	this := BaseContractCallSource{}
	this.SourceType = sourceType
	this.WalletId = walletId
	this.Address = address
	return &this
}

// NewBaseContractCallSourceWithDefaults instantiates a new BaseContractCallSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseContractCallSourceWithDefaults() *BaseContractCallSource {
	this := BaseContractCallSource{}
	return &this
}

// GetSourceType returns the SourceType field value
func (o *BaseContractCallSource) GetSourceType() ContractCallSourceType {
	if o == nil {
		var ret ContractCallSourceType
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *BaseContractCallSource) GetSourceTypeOk() (*ContractCallSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *BaseContractCallSource) SetSourceType(v ContractCallSourceType) {
	o.SourceType = v
}

// GetWalletId returns the WalletId field value
func (o *BaseContractCallSource) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *BaseContractCallSource) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *BaseContractCallSource) SetWalletId(v string) {
	o.WalletId = v
}

// GetAddress returns the Address field value
func (o *BaseContractCallSource) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *BaseContractCallSource) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *BaseContractCallSource) SetAddress(v string) {
	o.Address = v
}

func (o BaseContractCallSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseContractCallSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_type"] = o.SourceType
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

func (o *BaseContractCallSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source_type",
		"wallet_id",
		"address",
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

	varBaseContractCallSource := _BaseContractCallSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBaseContractCallSource)

	if err != nil {
		return err
	}

	*o = BaseContractCallSource(varBaseContractCallSource)

	return err
}

type NullableBaseContractCallSource struct {
	value *BaseContractCallSource
	isSet bool
}

func (v NullableBaseContractCallSource) Get() *BaseContractCallSource {
	return v.value
}

func (v *NullableBaseContractCallSource) Set(val *BaseContractCallSource) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseContractCallSource) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseContractCallSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseContractCallSource(val *BaseContractCallSource) *NullableBaseContractCallSource {
	return &NullableBaseContractCallSource{value: val, isSet: true}
}

func (v NullableBaseContractCallSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseContractCallSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


