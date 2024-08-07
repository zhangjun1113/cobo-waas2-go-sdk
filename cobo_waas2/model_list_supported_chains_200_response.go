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

// checks if the ListSupportedChains200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSupportedChains200Response{}

// ListSupportedChains200Response struct for ListSupportedChains200Response
type ListSupportedChains200Response struct {
	Data []ChainInfo `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// NewListSupportedChains200Response instantiates a new ListSupportedChains200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSupportedChains200Response() *ListSupportedChains200Response {
	this := ListSupportedChains200Response{}
	return &this
}

// NewListSupportedChains200ResponseWithDefaults instantiates a new ListSupportedChains200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSupportedChains200ResponseWithDefaults() *ListSupportedChains200Response {
	this := ListSupportedChains200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListSupportedChains200Response) GetData() []ChainInfo {
	if o == nil || IsNil(o.Data) {
		var ret []ChainInfo
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSupportedChains200Response) GetDataOk() ([]ChainInfo, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListSupportedChains200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ChainInfo and assigns it to the Data field.
func (o *ListSupportedChains200Response) SetData(v []ChainInfo) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListSupportedChains200Response) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSupportedChains200Response) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListSupportedChains200Response) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListSupportedChains200Response) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListSupportedChains200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSupportedChains200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListSupportedChains200Response struct {
	value *ListSupportedChains200Response
	isSet bool
}

func (v NullableListSupportedChains200Response) Get() *ListSupportedChains200Response {
	return v.value
}

func (v *NullableListSupportedChains200Response) Set(val *ListSupportedChains200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListSupportedChains200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListSupportedChains200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSupportedChains200Response(val *ListSupportedChains200Response) *NullableListSupportedChains200Response {
	return &NullableListSupportedChains200Response{value: val, isSet: true}
}

func (v NullableListSupportedChains200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSupportedChains200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


