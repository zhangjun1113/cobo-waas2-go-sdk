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

// checks if the BaseEstimateStakingFee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseEstimateStakingFee{}

// BaseEstimateStakingFee struct for BaseEstimateStakingFee
type BaseEstimateStakingFee struct {
	ActivityType ActivityType `json:"activity_type"`
}

type _BaseEstimateStakingFee BaseEstimateStakingFee

// NewBaseEstimateStakingFee instantiates a new BaseEstimateStakingFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEstimateStakingFee(activityType ActivityType) *BaseEstimateStakingFee {
	this := BaseEstimateStakingFee{}
	this.ActivityType = activityType
	return &this
}

// NewBaseEstimateStakingFeeWithDefaults instantiates a new BaseEstimateStakingFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEstimateStakingFeeWithDefaults() *BaseEstimateStakingFee {
	this := BaseEstimateStakingFee{}
	return &this
}

// GetActivityType returns the ActivityType field value
func (o *BaseEstimateStakingFee) GetActivityType() ActivityType {
	if o == nil {
		var ret ActivityType
		return ret
	}

	return o.ActivityType
}

// GetActivityTypeOk returns a tuple with the ActivityType field value
// and a boolean to check if the value has been set.
func (o *BaseEstimateStakingFee) GetActivityTypeOk() (*ActivityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActivityType, true
}

// SetActivityType sets field value
func (o *BaseEstimateStakingFee) SetActivityType(v ActivityType) {
	o.ActivityType = v
}

func (o BaseEstimateStakingFee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseEstimateStakingFee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["activity_type"] = o.ActivityType
	return toSerialize, nil
}

func (o *BaseEstimateStakingFee) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"activity_type",
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

	varBaseEstimateStakingFee := _BaseEstimateStakingFee{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBaseEstimateStakingFee)

	if err != nil {
		return err
	}

	*o = BaseEstimateStakingFee(varBaseEstimateStakingFee)

	return err
}

type NullableBaseEstimateStakingFee struct {
	value *BaseEstimateStakingFee
	isSet bool
}

func (v NullableBaseEstimateStakingFee) Get() *BaseEstimateStakingFee {
	return v.value
}

func (v *NullableBaseEstimateStakingFee) Set(val *BaseEstimateStakingFee) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEstimateStakingFee) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEstimateStakingFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEstimateStakingFee(val *BaseEstimateStakingFee) *NullableBaseEstimateStakingFee {
	return &NullableBaseEstimateStakingFee{value: val, isSet: true}
}

func (v NullableBaseEstimateStakingFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEstimateStakingFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

