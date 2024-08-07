/*
Cobo Wallet as a Service 2.0

API version: 1.1.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the ListWebhookEvents200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWebhookEvents200Response{}

// ListWebhookEvents200Response struct for ListWebhookEvents200Response
type ListWebhookEvents200Response struct {
	Data []WebhookEvent `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// NewListWebhookEvents200Response instantiates a new ListWebhookEvents200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWebhookEvents200Response() *ListWebhookEvents200Response {
	this := ListWebhookEvents200Response{}
	return &this
}

// NewListWebhookEvents200ResponseWithDefaults instantiates a new ListWebhookEvents200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWebhookEvents200ResponseWithDefaults() *ListWebhookEvents200Response {
	this := ListWebhookEvents200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListWebhookEvents200Response) GetData() []WebhookEvent {
	if o == nil || IsNil(o.Data) {
		var ret []WebhookEvent
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWebhookEvents200Response) GetDataOk() ([]WebhookEvent, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListWebhookEvents200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []WebhookEvent and assigns it to the Data field.
func (o *ListWebhookEvents200Response) SetData(v []WebhookEvent) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListWebhookEvents200Response) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWebhookEvents200Response) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListWebhookEvents200Response) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListWebhookEvents200Response) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListWebhookEvents200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWebhookEvents200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListWebhookEvents200Response struct {
	value *ListWebhookEvents200Response
	isSet bool
}

func (v NullableListWebhookEvents200Response) Get() *ListWebhookEvents200Response {
	return v.value
}

func (v *NullableListWebhookEvents200Response) Set(val *ListWebhookEvents200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListWebhookEvents200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListWebhookEvents200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWebhookEvents200Response(val *ListWebhookEvents200Response) *NullableListWebhookEvents200Response {
	return &NullableListWebhookEvents200Response{value: val, isSet: true}
}

func (v NullableListWebhookEvents200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWebhookEvents200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


