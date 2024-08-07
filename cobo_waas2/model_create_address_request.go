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

// checks if the CreateAddressRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAddressRequest{}

// CreateAddressRequest struct for CreateAddressRequest
type CreateAddressRequest struct {
	// The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](/v2/api-references/wallets/list-enabled-chains).
	ChainId string `json:"chain_id"`
	// The number of addresses to create.
	Count int32 `json:"count"`
	Encoding *AddressEncoding `json:"encoding,omitempty"`
}

type _CreateAddressRequest CreateAddressRequest

// NewCreateAddressRequest instantiates a new CreateAddressRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAddressRequest(chainId string, count int32) *CreateAddressRequest {
	this := CreateAddressRequest{}
	this.ChainId = chainId
	this.Count = count
	return &this
}

// NewCreateAddressRequestWithDefaults instantiates a new CreateAddressRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAddressRequestWithDefaults() *CreateAddressRequest {
	this := CreateAddressRequest{}
	var count int32 = 1
	this.Count = count
	return &this
}

// GetChainId returns the ChainId field value
func (o *CreateAddressRequest) GetChainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *CreateAddressRequest) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *CreateAddressRequest) SetChainId(v string) {
	o.ChainId = v
}

// GetCount returns the Count field value
func (o *CreateAddressRequest) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *CreateAddressRequest) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *CreateAddressRequest) SetCount(v int32) {
	o.Count = v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *CreateAddressRequest) GetEncoding() AddressEncoding {
	if o == nil || IsNil(o.Encoding) {
		var ret AddressEncoding
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAddressRequest) GetEncodingOk() (*AddressEncoding, bool) {
	if o == nil || IsNil(o.Encoding) {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *CreateAddressRequest) HasEncoding() bool {
	if o != nil && !IsNil(o.Encoding) {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given AddressEncoding and assigns it to the Encoding field.
func (o *CreateAddressRequest) SetEncoding(v AddressEncoding) {
	o.Encoding = &v
}

func (o CreateAddressRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAddressRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chain_id"] = o.ChainId
	toSerialize["count"] = o.Count
	if !IsNil(o.Encoding) {
		toSerialize["encoding"] = o.Encoding
	}
	return toSerialize, nil
}

func (o *CreateAddressRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chain_id",
		"count",
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

	varCreateAddressRequest := _CreateAddressRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateAddressRequest)

	if err != nil {
		return err
	}

	*o = CreateAddressRequest(varCreateAddressRequest)

	return err
}

type NullableCreateAddressRequest struct {
	value *CreateAddressRequest
	isSet bool
}

func (v NullableCreateAddressRequest) Get() *CreateAddressRequest {
	return v.value
}

func (v *NullableCreateAddressRequest) Set(val *CreateAddressRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAddressRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAddressRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAddressRequest(val *CreateAddressRequest) *NullableCreateAddressRequest {
	return &NullableCreateAddressRequest{value: val, isSet: true}
}

func (v NullableCreateAddressRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAddressRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


