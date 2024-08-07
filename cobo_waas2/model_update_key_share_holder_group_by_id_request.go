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

// checks if the UpdateKeyShareHolderGroupByIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateKeyShareHolderGroupByIdRequest{}

// UpdateKeyShareHolderGroupByIdRequest struct for UpdateKeyShareHolderGroupByIdRequest
type UpdateKeyShareHolderGroupByIdRequest struct {
	UpdateKeyShareHolderGroupAction UpdateGroupAction `json:"update_key_share_holder_group_action"`
}

type _UpdateKeyShareHolderGroupByIdRequest UpdateKeyShareHolderGroupByIdRequest

// NewUpdateKeyShareHolderGroupByIdRequest instantiates a new UpdateKeyShareHolderGroupByIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateKeyShareHolderGroupByIdRequest(updateKeyShareHolderGroupAction UpdateGroupAction) *UpdateKeyShareHolderGroupByIdRequest {
	this := UpdateKeyShareHolderGroupByIdRequest{}
	this.UpdateKeyShareHolderGroupAction = updateKeyShareHolderGroupAction
	return &this
}

// NewUpdateKeyShareHolderGroupByIdRequestWithDefaults instantiates a new UpdateKeyShareHolderGroupByIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateKeyShareHolderGroupByIdRequestWithDefaults() *UpdateKeyShareHolderGroupByIdRequest {
	this := UpdateKeyShareHolderGroupByIdRequest{}
	return &this
}

// GetUpdateKeyShareHolderGroupAction returns the UpdateKeyShareHolderGroupAction field value
func (o *UpdateKeyShareHolderGroupByIdRequest) GetUpdateKeyShareHolderGroupAction() UpdateGroupAction {
	if o == nil {
		var ret UpdateGroupAction
		return ret
	}

	return o.UpdateKeyShareHolderGroupAction
}

// GetUpdateKeyShareHolderGroupActionOk returns a tuple with the UpdateKeyShareHolderGroupAction field value
// and a boolean to check if the value has been set.
func (o *UpdateKeyShareHolderGroupByIdRequest) GetUpdateKeyShareHolderGroupActionOk() (*UpdateGroupAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateKeyShareHolderGroupAction, true
}

// SetUpdateKeyShareHolderGroupAction sets field value
func (o *UpdateKeyShareHolderGroupByIdRequest) SetUpdateKeyShareHolderGroupAction(v UpdateGroupAction) {
	o.UpdateKeyShareHolderGroupAction = v
}

func (o UpdateKeyShareHolderGroupByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateKeyShareHolderGroupByIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["update_key_share_holder_group_action"] = o.UpdateKeyShareHolderGroupAction
	return toSerialize, nil
}

func (o *UpdateKeyShareHolderGroupByIdRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"update_key_share_holder_group_action",
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

	varUpdateKeyShareHolderGroupByIdRequest := _UpdateKeyShareHolderGroupByIdRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateKeyShareHolderGroupByIdRequest)

	if err != nil {
		return err
	}

	*o = UpdateKeyShareHolderGroupByIdRequest(varUpdateKeyShareHolderGroupByIdRequest)

	return err
}

type NullableUpdateKeyShareHolderGroupByIdRequest struct {
	value *UpdateKeyShareHolderGroupByIdRequest
	isSet bool
}

func (v NullableUpdateKeyShareHolderGroupByIdRequest) Get() *UpdateKeyShareHolderGroupByIdRequest {
	return v.value
}

func (v *NullableUpdateKeyShareHolderGroupByIdRequest) Set(val *UpdateKeyShareHolderGroupByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateKeyShareHolderGroupByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateKeyShareHolderGroupByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateKeyShareHolderGroupByIdRequest(val *UpdateKeyShareHolderGroupByIdRequest) *NullableUpdateKeyShareHolderGroupByIdRequest {
	return &NullableUpdateKeyShareHolderGroupByIdRequest{value: val, isSet: true}
}

func (v NullableUpdateKeyShareHolderGroupByIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateKeyShareHolderGroupByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


