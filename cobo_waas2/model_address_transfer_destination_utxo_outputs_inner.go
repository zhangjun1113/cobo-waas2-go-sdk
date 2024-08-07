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

// checks if the AddressTransferDestinationUtxoOutputsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressTransferDestinationUtxoOutputsInner{}

// AddressTransferDestinationUtxoOutputsInner struct for AddressTransferDestinationUtxoOutputsInner
type AddressTransferDestinationUtxoOutputsInner struct {
	// The destination address.
	Address string `json:"address"`
	// The quantity of the token in the transaction. For example, if you trade 1.5 ETH, then the value is `1.5`. 
	Amount *string `json:"amount,omitempty"`
	// The script of the output. It is a programmable code fragment that defines the conditions under which the UTXO can be spent.
	Script *string `json:"script,omitempty"`
}

type _AddressTransferDestinationUtxoOutputsInner AddressTransferDestinationUtxoOutputsInner

// NewAddressTransferDestinationUtxoOutputsInner instantiates a new AddressTransferDestinationUtxoOutputsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransferDestinationUtxoOutputsInner(address string) *AddressTransferDestinationUtxoOutputsInner {
	this := AddressTransferDestinationUtxoOutputsInner{}
	this.Address = address
	return &this
}

// NewAddressTransferDestinationUtxoOutputsInnerWithDefaults instantiates a new AddressTransferDestinationUtxoOutputsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransferDestinationUtxoOutputsInnerWithDefaults() *AddressTransferDestinationUtxoOutputsInner {
	this := AddressTransferDestinationUtxoOutputsInner{}
	return &this
}

// GetAddress returns the Address field value
func (o *AddressTransferDestinationUtxoOutputsInner) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *AddressTransferDestinationUtxoOutputsInner) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *AddressTransferDestinationUtxoOutputsInner) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AddressTransferDestinationUtxoOutputsInner) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransferDestinationUtxoOutputsInner) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AddressTransferDestinationUtxoOutputsInner) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AddressTransferDestinationUtxoOutputsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *AddressTransferDestinationUtxoOutputsInner) GetScript() string {
	if o == nil || IsNil(o.Script) {
		var ret string
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransferDestinationUtxoOutputsInner) GetScriptOk() (*string, bool) {
	if o == nil || IsNil(o.Script) {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *AddressTransferDestinationUtxoOutputsInner) HasScript() bool {
	if o != nil && !IsNil(o.Script) {
		return true
	}

	return false
}

// SetScript gets a reference to the given string and assigns it to the Script field.
func (o *AddressTransferDestinationUtxoOutputsInner) SetScript(v string) {
	o.Script = &v
}

func (o AddressTransferDestinationUtxoOutputsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransferDestinationUtxoOutputsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Script) {
		toSerialize["script"] = o.Script
	}
	return toSerialize, nil
}

func (o *AddressTransferDestinationUtxoOutputsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varAddressTransferDestinationUtxoOutputsInner := _AddressTransferDestinationUtxoOutputsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransferDestinationUtxoOutputsInner)

	if err != nil {
		return err
	}

	*o = AddressTransferDestinationUtxoOutputsInner(varAddressTransferDestinationUtxoOutputsInner)

	return err
}

type NullableAddressTransferDestinationUtxoOutputsInner struct {
	value *AddressTransferDestinationUtxoOutputsInner
	isSet bool
}

func (v NullableAddressTransferDestinationUtxoOutputsInner) Get() *AddressTransferDestinationUtxoOutputsInner {
	return v.value
}

func (v *NullableAddressTransferDestinationUtxoOutputsInner) Set(val *AddressTransferDestinationUtxoOutputsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransferDestinationUtxoOutputsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransferDestinationUtxoOutputsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransferDestinationUtxoOutputsInner(val *AddressTransferDestinationUtxoOutputsInner) *NullableAddressTransferDestinationUtxoOutputsInner {
	return &NullableAddressTransferDestinationUtxoOutputsInner{value: val, isSet: true}
}

func (v NullableAddressTransferDestinationUtxoOutputsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransferDestinationUtxoOutputsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


