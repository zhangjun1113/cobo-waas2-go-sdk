/*
Cobo Wallet as a Service 2.0

Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the StakingsValidatorInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StakingsValidatorInfo{}

// StakingsValidatorInfo The validator info of the stake.
type StakingsValidatorInfo struct {
	IconUrl *string `json:"icon_url,omitempty"`
	PublicKey *string `json:"public_key,omitempty"`
	Name *string `json:"name,omitempty"`
	Address *string `json:"address,omitempty"`
	CommissionRate *string `json:"commission_rate,omitempty"`
}

// NewStakingsValidatorInfo instantiates a new StakingsValidatorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStakingsValidatorInfo() *StakingsValidatorInfo {
	this := StakingsValidatorInfo{}
	return &this
}

// NewStakingsValidatorInfoWithDefaults instantiates a new StakingsValidatorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStakingsValidatorInfoWithDefaults() *StakingsValidatorInfo {
	this := StakingsValidatorInfo{}
	return &this
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *StakingsValidatorInfo) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl) {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakingsValidatorInfo) GetIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IconUrl) {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *StakingsValidatorInfo) HasIconUrl() bool {
	if o != nil && !IsNil(o.IconUrl) {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *StakingsValidatorInfo) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *StakingsValidatorInfo) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakingsValidatorInfo) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *StakingsValidatorInfo) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *StakingsValidatorInfo) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StakingsValidatorInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakingsValidatorInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StakingsValidatorInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StakingsValidatorInfo) SetName(v string) {
	o.Name = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *StakingsValidatorInfo) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakingsValidatorInfo) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *StakingsValidatorInfo) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *StakingsValidatorInfo) SetAddress(v string) {
	o.Address = &v
}

// GetCommissionRate returns the CommissionRate field value if set, zero value otherwise.
func (o *StakingsValidatorInfo) GetCommissionRate() string {
	if o == nil || IsNil(o.CommissionRate) {
		var ret string
		return ret
	}
	return *o.CommissionRate
}

// GetCommissionRateOk returns a tuple with the CommissionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakingsValidatorInfo) GetCommissionRateOk() (*string, bool) {
	if o == nil || IsNil(o.CommissionRate) {
		return nil, false
	}
	return o.CommissionRate, true
}

// HasCommissionRate returns a boolean if a field has been set.
func (o *StakingsValidatorInfo) HasCommissionRate() bool {
	if o != nil && !IsNil(o.CommissionRate) {
		return true
	}

	return false
}

// SetCommissionRate gets a reference to the given string and assigns it to the CommissionRate field.
func (o *StakingsValidatorInfo) SetCommissionRate(v string) {
	o.CommissionRate = &v
}

func (o StakingsValidatorInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StakingsValidatorInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IconUrl) {
		toSerialize["icon_url"] = o.IconUrl
	}
	if !IsNil(o.PublicKey) {
		toSerialize["public_key"] = o.PublicKey
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.CommissionRate) {
		toSerialize["commission_rate"] = o.CommissionRate
	}
	return toSerialize, nil
}

type NullableStakingsValidatorInfo struct {
	value *StakingsValidatorInfo
	isSet bool
}

func (v NullableStakingsValidatorInfo) Get() *StakingsValidatorInfo {
	return v.value
}

func (v *NullableStakingsValidatorInfo) Set(val *StakingsValidatorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStakingsValidatorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStakingsValidatorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStakingsValidatorInfo(val *StakingsValidatorInfo) *NullableStakingsValidatorInfo {
	return &NullableStakingsValidatorInfo{value: val, isSet: true}
}

func (v NullableStakingsValidatorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStakingsValidatorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


