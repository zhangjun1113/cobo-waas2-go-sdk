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

// checks if the TransactionTransferToAddressDestinationAccountOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionTransferToAddressDestinationAccountOutput{}

// TransactionTransferToAddressDestinationAccountOutput struct for TransactionTransferToAddressDestinationAccountOutput
type TransactionTransferToAddressDestinationAccountOutput struct {
	// The destination address.
	Address *string `json:"address,omitempty"`
	// The memo that identifies a transaction in order to credit the correct account. For transfers out of Cobo Portal, it is highly recommended to include a memo for the chains such as XRP, EOS, XLM, IOST, BNB_BNB, ATOM, LUNA, and TON.
	Memo *string `json:"memo,omitempty"`
	// The quantity of the token in the transaction. For example, if you trade 1.5 ETH, then the value is `1.5`. 
	Amount *string `json:"amount,omitempty"`
}

// NewTransactionTransferToAddressDestinationAccountOutput instantiates a new TransactionTransferToAddressDestinationAccountOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionTransferToAddressDestinationAccountOutput() *TransactionTransferToAddressDestinationAccountOutput {
	this := TransactionTransferToAddressDestinationAccountOutput{}
	return &this
}

// NewTransactionTransferToAddressDestinationAccountOutputWithDefaults instantiates a new TransactionTransferToAddressDestinationAccountOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionTransferToAddressDestinationAccountOutputWithDefaults() *TransactionTransferToAddressDestinationAccountOutput {
	this := TransactionTransferToAddressDestinationAccountOutput{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *TransactionTransferToAddressDestinationAccountOutput) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTransferToAddressDestinationAccountOutput) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *TransactionTransferToAddressDestinationAccountOutput) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *TransactionTransferToAddressDestinationAccountOutput) SetAddress(v string) {
	o.Address = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *TransactionTransferToAddressDestinationAccountOutput) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTransferToAddressDestinationAccountOutput) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *TransactionTransferToAddressDestinationAccountOutput) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *TransactionTransferToAddressDestinationAccountOutput) SetMemo(v string) {
	o.Memo = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TransactionTransferToAddressDestinationAccountOutput) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTransferToAddressDestinationAccountOutput) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TransactionTransferToAddressDestinationAccountOutput) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *TransactionTransferToAddressDestinationAccountOutput) SetAmount(v string) {
	o.Amount = &v
}

func (o TransactionTransferToAddressDestinationAccountOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionTransferToAddressDestinationAccountOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableTransactionTransferToAddressDestinationAccountOutput struct {
	value *TransactionTransferToAddressDestinationAccountOutput
	isSet bool
}

func (v NullableTransactionTransferToAddressDestinationAccountOutput) Get() *TransactionTransferToAddressDestinationAccountOutput {
	return v.value
}

func (v *NullableTransactionTransferToAddressDestinationAccountOutput) Set(val *TransactionTransferToAddressDestinationAccountOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionTransferToAddressDestinationAccountOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionTransferToAddressDestinationAccountOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionTransferToAddressDestinationAccountOutput(val *TransactionTransferToAddressDestinationAccountOutput) *NullableTransactionTransferToAddressDestinationAccountOutput {
	return &NullableTransactionTransferToAddressDestinationAccountOutput{value: val, isSet: true}
}

func (v NullableTransactionTransferToAddressDestinationAccountOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionTransferToAddressDestinationAccountOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


