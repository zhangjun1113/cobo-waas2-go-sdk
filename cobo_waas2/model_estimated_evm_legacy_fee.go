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

// checks if the EstimatedEvmLegacyFee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EstimatedEvmLegacyFee{}

// EstimatedEvmLegacyFee The estimated transaction fee based on the legacy fee model.
type EstimatedEvmLegacyFee struct {
	FeeType FeeType `json:"fee_type"`
	// The token ID of the transaction fee.
	TokenId string `json:"token_id"`
	Slow *EstimatedEvmLegacyFeeSlow `json:"slow,omitempty"`
	Recommended EstimatedEvmLegacyFeeSlow `json:"recommended"`
	Fast *EstimatedEvmLegacyFeeSlow `json:"fast,omitempty"`
}

type _EstimatedEvmLegacyFee EstimatedEvmLegacyFee

// NewEstimatedEvmLegacyFee instantiates a new EstimatedEvmLegacyFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimatedEvmLegacyFee(feeType FeeType, tokenId string, recommended EstimatedEvmLegacyFeeSlow) *EstimatedEvmLegacyFee {
	this := EstimatedEvmLegacyFee{}
	this.FeeType = feeType
	this.TokenId = tokenId
	this.Recommended = recommended
	return &this
}

// NewEstimatedEvmLegacyFeeWithDefaults instantiates a new EstimatedEvmLegacyFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimatedEvmLegacyFeeWithDefaults() *EstimatedEvmLegacyFee {
	this := EstimatedEvmLegacyFee{}
	var feeType FeeType = FEETYPE_EVM_EIP_1559
	this.FeeType = feeType
	return &this
}

// GetFeeType returns the FeeType field value
func (o *EstimatedEvmLegacyFee) GetFeeType() FeeType {
	if o == nil {
		var ret FeeType
		return ret
	}

	return o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value
// and a boolean to check if the value has been set.
func (o *EstimatedEvmLegacyFee) GetFeeTypeOk() (*FeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeType, true
}

// SetFeeType sets field value
func (o *EstimatedEvmLegacyFee) SetFeeType(v FeeType) {
	o.FeeType = v
}

// GetTokenId returns the TokenId field value
func (o *EstimatedEvmLegacyFee) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *EstimatedEvmLegacyFee) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *EstimatedEvmLegacyFee) SetTokenId(v string) {
	o.TokenId = v
}

// GetSlow returns the Slow field value if set, zero value otherwise.
func (o *EstimatedEvmLegacyFee) GetSlow() EstimatedEvmLegacyFeeSlow {
	if o == nil || IsNil(o.Slow) {
		var ret EstimatedEvmLegacyFeeSlow
		return ret
	}
	return *o.Slow
}

// GetSlowOk returns a tuple with the Slow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedEvmLegacyFee) GetSlowOk() (*EstimatedEvmLegacyFeeSlow, bool) {
	if o == nil || IsNil(o.Slow) {
		return nil, false
	}
	return o.Slow, true
}

// HasSlow returns a boolean if a field has been set.
func (o *EstimatedEvmLegacyFee) HasSlow() bool {
	if o != nil && !IsNil(o.Slow) {
		return true
	}

	return false
}

// SetSlow gets a reference to the given EstimatedEvmLegacyFeeSlow and assigns it to the Slow field.
func (o *EstimatedEvmLegacyFee) SetSlow(v EstimatedEvmLegacyFeeSlow) {
	o.Slow = &v
}

// GetRecommended returns the Recommended field value
func (o *EstimatedEvmLegacyFee) GetRecommended() EstimatedEvmLegacyFeeSlow {
	if o == nil {
		var ret EstimatedEvmLegacyFeeSlow
		return ret
	}

	return o.Recommended
}

// GetRecommendedOk returns a tuple with the Recommended field value
// and a boolean to check if the value has been set.
func (o *EstimatedEvmLegacyFee) GetRecommendedOk() (*EstimatedEvmLegacyFeeSlow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recommended, true
}

// SetRecommended sets field value
func (o *EstimatedEvmLegacyFee) SetRecommended(v EstimatedEvmLegacyFeeSlow) {
	o.Recommended = v
}

// GetFast returns the Fast field value if set, zero value otherwise.
func (o *EstimatedEvmLegacyFee) GetFast() EstimatedEvmLegacyFeeSlow {
	if o == nil || IsNil(o.Fast) {
		var ret EstimatedEvmLegacyFeeSlow
		return ret
	}
	return *o.Fast
}

// GetFastOk returns a tuple with the Fast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedEvmLegacyFee) GetFastOk() (*EstimatedEvmLegacyFeeSlow, bool) {
	if o == nil || IsNil(o.Fast) {
		return nil, false
	}
	return o.Fast, true
}

// HasFast returns a boolean if a field has been set.
func (o *EstimatedEvmLegacyFee) HasFast() bool {
	if o != nil && !IsNil(o.Fast) {
		return true
	}

	return false
}

// SetFast gets a reference to the given EstimatedEvmLegacyFeeSlow and assigns it to the Fast field.
func (o *EstimatedEvmLegacyFee) SetFast(v EstimatedEvmLegacyFeeSlow) {
	o.Fast = &v
}

func (o EstimatedEvmLegacyFee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EstimatedEvmLegacyFee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fee_type"] = o.FeeType
	toSerialize["token_id"] = o.TokenId
	if !IsNil(o.Slow) {
		toSerialize["slow"] = o.Slow
	}
	toSerialize["recommended"] = o.Recommended
	if !IsNil(o.Fast) {
		toSerialize["fast"] = o.Fast
	}
	return toSerialize, nil
}

func (o *EstimatedEvmLegacyFee) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fee_type",
		"token_id",
		"recommended",
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

	varEstimatedEvmLegacyFee := _EstimatedEvmLegacyFee{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEstimatedEvmLegacyFee)

	if err != nil {
		return err
	}

	*o = EstimatedEvmLegacyFee(varEstimatedEvmLegacyFee)

	return err
}

type NullableEstimatedEvmLegacyFee struct {
	value *EstimatedEvmLegacyFee
	isSet bool
}

func (v NullableEstimatedEvmLegacyFee) Get() *EstimatedEvmLegacyFee {
	return v.value
}

func (v *NullableEstimatedEvmLegacyFee) Set(val *EstimatedEvmLegacyFee) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimatedEvmLegacyFee) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimatedEvmLegacyFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimatedEvmLegacyFee(val *EstimatedEvmLegacyFee) *NullableEstimatedEvmLegacyFee {
	return &NullableEstimatedEvmLegacyFee{value: val, isSet: true}
}

func (v NullableEstimatedEvmLegacyFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimatedEvmLegacyFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


