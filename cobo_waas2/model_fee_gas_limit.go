/*
Cobo Wallet as a Service 2.0

Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the FeeGasLimit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeeGasLimit{}

// FeeGasLimit struct for FeeGasLimit
type FeeGasLimit struct {
	// The gas limit. It represents the maximum number of gas units that you are willing to pay for the execution of a transaction or Ethereum Virtual Machine (EVM) operation. The gas unit cost of each operation varies.
	GasLimit *string `json:"gas_limit,omitempty"`
}

// NewFeeGasLimit instantiates a new FeeGasLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeeGasLimit() *FeeGasLimit {
	this := FeeGasLimit{}
	var gasLimit string = "21000"
	this.GasLimit = &gasLimit
	return &this
}

// NewFeeGasLimitWithDefaults instantiates a new FeeGasLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeGasLimitWithDefaults() *FeeGasLimit {
	this := FeeGasLimit{}
	var gasLimit string = "21000"
	this.GasLimit = &gasLimit
	return &this
}

// GetGasLimit returns the GasLimit field value if set, zero value otherwise.
func (o *FeeGasLimit) GetGasLimit() string {
	if o == nil || IsNil(o.GasLimit) {
		var ret string
		return ret
	}
	return *o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeGasLimit) GetGasLimitOk() (*string, bool) {
	if o == nil || IsNil(o.GasLimit) {
		return nil, false
	}
	return o.GasLimit, true
}

// HasGasLimit returns a boolean if a field has been set.
func (o *FeeGasLimit) HasGasLimit() bool {
	if o != nil && !IsNil(o.GasLimit) {
		return true
	}

	return false
}

// SetGasLimit gets a reference to the given string and assigns it to the GasLimit field.
func (o *FeeGasLimit) SetGasLimit(v string) {
	o.GasLimit = &v
}

func (o FeeGasLimit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeeGasLimit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GasLimit) {
		toSerialize["gas_limit"] = o.GasLimit
	}
	return toSerialize, nil
}

type NullableFeeGasLimit struct {
	value *FeeGasLimit
	isSet bool
}

func (v NullableFeeGasLimit) Get() *FeeGasLimit {
	return v.value
}

func (v *NullableFeeGasLimit) Set(val *FeeGasLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeGasLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeGasLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeGasLimit(val *FeeGasLimit) *NullableFeeGasLimit {
	return &NullableFeeGasLimit{value: val, isSet: true}
}

func (v NullableFeeGasLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeeGasLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


