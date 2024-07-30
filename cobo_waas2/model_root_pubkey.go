/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the RootPubkey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RootPubkey{}

// RootPubkey The data for MPC Root Extended Public Key information.
type RootPubkey struct {
	// The vault's [root extended public key](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/tss-node-deployment#tss-node-on-cobo-portal-and-mpc-root-extended-public-key).
	Pubkey *string `json:"pubkey,omitempty"`
	Curve *CurveType `json:"curve,omitempty"`
}

// NewRootPubkey instantiates a new RootPubkey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRootPubkey() *RootPubkey {
	this := RootPubkey{}
	return &this
}

// NewRootPubkeyWithDefaults instantiates a new RootPubkey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRootPubkeyWithDefaults() *RootPubkey {
	this := RootPubkey{}
	return &this
}

// GetPubkey returns the Pubkey field value if set, zero value otherwise.
func (o *RootPubkey) GetPubkey() string {
	if o == nil || IsNil(o.Pubkey) {
		var ret string
		return ret
	}
	return *o.Pubkey
}

// GetPubkeyOk returns a tuple with the Pubkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootPubkey) GetPubkeyOk() (*string, bool) {
	if o == nil || IsNil(o.Pubkey) {
		return nil, false
	}
	return o.Pubkey, true
}

// HasPubkey returns a boolean if a field has been set.
func (o *RootPubkey) HasPubkey() bool {
	if o != nil && !IsNil(o.Pubkey) {
		return true
	}

	return false
}

// SetPubkey gets a reference to the given string and assigns it to the Pubkey field.
func (o *RootPubkey) SetPubkey(v string) {
	o.Pubkey = &v
}

// GetCurve returns the Curve field value if set, zero value otherwise.
func (o *RootPubkey) GetCurve() CurveType {
	if o == nil || IsNil(o.Curve) {
		var ret CurveType
		return ret
	}
	return *o.Curve
}

// GetCurveOk returns a tuple with the Curve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootPubkey) GetCurveOk() (*CurveType, bool) {
	if o == nil || IsNil(o.Curve) {
		return nil, false
	}
	return o.Curve, true
}

// HasCurve returns a boolean if a field has been set.
func (o *RootPubkey) HasCurve() bool {
	if o != nil && !IsNil(o.Curve) {
		return true
	}

	return false
}

// SetCurve gets a reference to the given CurveType and assigns it to the Curve field.
func (o *RootPubkey) SetCurve(v CurveType) {
	o.Curve = &v
}

func (o RootPubkey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RootPubkey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pubkey) {
		toSerialize["pubkey"] = o.Pubkey
	}
	if !IsNil(o.Curve) {
		toSerialize["curve"] = o.Curve
	}
	return toSerialize, nil
}

type NullableRootPubkey struct {
	value *RootPubkey
	isSet bool
}

func (v NullableRootPubkey) Get() *RootPubkey {
	return v.value
}

func (v *NullableRootPubkey) Set(val *RootPubkey) {
	v.value = val
	v.isSet = true
}

func (v NullableRootPubkey) IsSet() bool {
	return v.isSet
}

func (v *NullableRootPubkey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRootPubkey(val *RootPubkey) *NullableRootPubkey {
	return &NullableRootPubkey{value: val, isSet: true}
}

func (v NullableRootPubkey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRootPubkey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

