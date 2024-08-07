/*
Cobo Wallet as a Service 2.0

Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the MaxFeeAmount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaxFeeAmount{}

// MaxFeeAmount The maximum transaction fee.
type MaxFeeAmount struct {
	// The maximum fee that you are willing to pay for the transaction. The transaction will fail if the transaction fee exceeds the maximum fee.
	MaxFeeAmount *string `json:"max_fee_amount,omitempty"`
}

// NewMaxFeeAmount instantiates a new MaxFeeAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaxFeeAmount() *MaxFeeAmount {
	this := MaxFeeAmount{}
	return &this
}

// NewMaxFeeAmountWithDefaults instantiates a new MaxFeeAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaxFeeAmountWithDefaults() *MaxFeeAmount {
	this := MaxFeeAmount{}
	return &this
}

// GetMaxFeeAmount returns the MaxFeeAmount field value if set, zero value otherwise.
func (o *MaxFeeAmount) GetMaxFeeAmount() string {
	if o == nil || IsNil(o.MaxFeeAmount) {
		var ret string
		return ret
	}
	return *o.MaxFeeAmount
}

// GetMaxFeeAmountOk returns a tuple with the MaxFeeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxFeeAmount) GetMaxFeeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.MaxFeeAmount) {
		return nil, false
	}
	return o.MaxFeeAmount, true
}

// HasMaxFeeAmount returns a boolean if a field has been set.
func (o *MaxFeeAmount) HasMaxFeeAmount() bool {
	if o != nil && !IsNil(o.MaxFeeAmount) {
		return true
	}

	return false
}

// SetMaxFeeAmount gets a reference to the given string and assigns it to the MaxFeeAmount field.
func (o *MaxFeeAmount) SetMaxFeeAmount(v string) {
	o.MaxFeeAmount = &v
}

func (o MaxFeeAmount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaxFeeAmount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxFeeAmount) {
		toSerialize["max_fee_amount"] = o.MaxFeeAmount
	}
	return toSerialize, nil
}

type NullableMaxFeeAmount struct {
	value *MaxFeeAmount
	isSet bool
}

func (v NullableMaxFeeAmount) Get() *MaxFeeAmount {
	return v.value
}

func (v *NullableMaxFeeAmount) Set(val *MaxFeeAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableMaxFeeAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableMaxFeeAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaxFeeAmount(val *MaxFeeAmount) *NullableMaxFeeAmount {
	return &NullableMaxFeeAmount{value: val, isSet: true}
}

func (v NullableMaxFeeAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaxFeeAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


