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

// checks if the WebhookEventDataType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookEventDataType{}

// WebhookEventDataType The data type of the event.
type WebhookEventDataType struct {
	// The data type of the event. When `data_type` is `Transaction`, it means the event uses the `transaction` schema as its data type.
	DataType string `json:"data_type"`
}

type _WebhookEventDataType WebhookEventDataType

// NewWebhookEventDataType instantiates a new WebhookEventDataType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEventDataType(dataType string) *WebhookEventDataType {
	this := WebhookEventDataType{}
	this.DataType = dataType
	return &this
}

// NewWebhookEventDataTypeWithDefaults instantiates a new WebhookEventDataType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEventDataTypeWithDefaults() *WebhookEventDataType {
	this := WebhookEventDataType{}
	return &this
}

// GetDataType returns the DataType field value
func (o *WebhookEventDataType) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *WebhookEventDataType) GetDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *WebhookEventDataType) SetDataType(v string) {
	o.DataType = v
}

func (o WebhookEventDataType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookEventDataType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data_type"] = o.DataType
	return toSerialize, nil
}

func (o *WebhookEventDataType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data_type",
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

	varWebhookEventDataType := _WebhookEventDataType{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookEventDataType)

	if err != nil {
		return err
	}

	*o = WebhookEventDataType(varWebhookEventDataType)

	return err
}

type NullableWebhookEventDataType struct {
	value *WebhookEventDataType
	isSet bool
}

func (v NullableWebhookEventDataType) Get() *WebhookEventDataType {
	return v.value
}

func (v *NullableWebhookEventDataType) Set(val *WebhookEventDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEventDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEventDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEventDataType(val *WebhookEventDataType) *NullableWebhookEventDataType {
	return &NullableWebhookEventDataType{value: val, isSet: true}
}

func (v NullableWebhookEventDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEventDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


