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

// checks if the UpdateMpcVaultByIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMpcVaultByIdRequest{}

// UpdateMpcVaultByIdRequest struct for UpdateMpcVaultByIdRequest
type UpdateMpcVaultByIdRequest struct {
	// The new name of the vault.
	Name string `json:"name"`
}

type _UpdateMpcVaultByIdRequest UpdateMpcVaultByIdRequest

// NewUpdateMpcVaultByIdRequest instantiates a new UpdateMpcVaultByIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMpcVaultByIdRequest(name string) *UpdateMpcVaultByIdRequest {
	this := UpdateMpcVaultByIdRequest{}
	this.Name = name
	return &this
}

// NewUpdateMpcVaultByIdRequestWithDefaults instantiates a new UpdateMpcVaultByIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMpcVaultByIdRequestWithDefaults() *UpdateMpcVaultByIdRequest {
	this := UpdateMpcVaultByIdRequest{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateMpcVaultByIdRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateMpcVaultByIdRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateMpcVaultByIdRequest) SetName(v string) {
	o.Name = v
}

func (o UpdateMpcVaultByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMpcVaultByIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *UpdateMpcVaultByIdRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varUpdateMpcVaultByIdRequest := _UpdateMpcVaultByIdRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateMpcVaultByIdRequest)

	if err != nil {
		return err
	}

	*o = UpdateMpcVaultByIdRequest(varUpdateMpcVaultByIdRequest)

	return err
}

type NullableUpdateMpcVaultByIdRequest struct {
	value *UpdateMpcVaultByIdRequest
	isSet bool
}

func (v NullableUpdateMpcVaultByIdRequest) Get() *UpdateMpcVaultByIdRequest {
	return v.value
}

func (v *NullableUpdateMpcVaultByIdRequest) Set(val *UpdateMpcVaultByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMpcVaultByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMpcVaultByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMpcVaultByIdRequest(val *UpdateMpcVaultByIdRequest) *NullableUpdateMpcVaultByIdRequest {
	return &NullableUpdateMpcVaultByIdRequest{value: val, isSet: true}
}

func (v NullableUpdateMpcVaultByIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMpcVaultByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


