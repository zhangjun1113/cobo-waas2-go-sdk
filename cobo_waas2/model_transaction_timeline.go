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

// checks if the TransactionTimeline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionTimeline{}

// TransactionTimeline The information about transaction timeline, which lists all statuses that the transaction passes through with timestamps indicating when each status is completed.
type TransactionTimeline struct {
	Status *TransactionStatus `json:"status,omitempty"`
	// Whether the transaction status is completed:   - `true`: The transaction status is completed.   - `false`: The transaction is currently in the status. 
	Finished *bool `json:"finished,omitempty"`
	// The time when the transaction status is completed in Unix timestamp format, measured in milliseconds.
	FinishedTimestamp *int64 `json:"finished_timestamp,omitempty"`
}

// NewTransactionTimeline instantiates a new TransactionTimeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionTimeline() *TransactionTimeline {
	this := TransactionTimeline{}
	return &this
}

// NewTransactionTimelineWithDefaults instantiates a new TransactionTimeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionTimelineWithDefaults() *TransactionTimeline {
	this := TransactionTimeline{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TransactionTimeline) GetStatus() TransactionStatus {
	if o == nil || IsNil(o.Status) {
		var ret TransactionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTimeline) GetStatusOk() (*TransactionStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TransactionTimeline) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TransactionStatus and assigns it to the Status field.
func (o *TransactionTimeline) SetStatus(v TransactionStatus) {
	o.Status = &v
}

// GetFinished returns the Finished field value if set, zero value otherwise.
func (o *TransactionTimeline) GetFinished() bool {
	if o == nil || IsNil(o.Finished) {
		var ret bool
		return ret
	}
	return *o.Finished
}

// GetFinishedOk returns a tuple with the Finished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTimeline) GetFinishedOk() (*bool, bool) {
	if o == nil || IsNil(o.Finished) {
		return nil, false
	}
	return o.Finished, true
}

// HasFinished returns a boolean if a field has been set.
func (o *TransactionTimeline) HasFinished() bool {
	if o != nil && !IsNil(o.Finished) {
		return true
	}

	return false
}

// SetFinished gets a reference to the given bool and assigns it to the Finished field.
func (o *TransactionTimeline) SetFinished(v bool) {
	o.Finished = &v
}

// GetFinishedTimestamp returns the FinishedTimestamp field value if set, zero value otherwise.
func (o *TransactionTimeline) GetFinishedTimestamp() int64 {
	if o == nil || IsNil(o.FinishedTimestamp) {
		var ret int64
		return ret
	}
	return *o.FinishedTimestamp
}

// GetFinishedTimestampOk returns a tuple with the FinishedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTimeline) GetFinishedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.FinishedTimestamp) {
		return nil, false
	}
	return o.FinishedTimestamp, true
}

// HasFinishedTimestamp returns a boolean if a field has been set.
func (o *TransactionTimeline) HasFinishedTimestamp() bool {
	if o != nil && !IsNil(o.FinishedTimestamp) {
		return true
	}

	return false
}

// SetFinishedTimestamp gets a reference to the given int64 and assigns it to the FinishedTimestamp field.
func (o *TransactionTimeline) SetFinishedTimestamp(v int64) {
	o.FinishedTimestamp = &v
}

func (o TransactionTimeline) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionTimeline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Finished) {
		toSerialize["finished"] = o.Finished
	}
	if !IsNil(o.FinishedTimestamp) {
		toSerialize["finished_timestamp"] = o.FinishedTimestamp
	}
	return toSerialize, nil
}

type NullableTransactionTimeline struct {
	value *TransactionTimeline
	isSet bool
}

func (v NullableTransactionTimeline) Get() *TransactionTimeline {
	return v.value
}

func (v *NullableTransactionTimeline) Set(val *TransactionTimeline) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionTimeline) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionTimeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionTimeline(val *TransactionTimeline) *NullableTransactionTimeline {
	return &NullableTransactionTimeline{value: val, isSet: true}
}

func (v NullableTransactionTimeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionTimeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


