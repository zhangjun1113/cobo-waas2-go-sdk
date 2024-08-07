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

// checks if the TransactionRequestFixedFee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionRequestFixedFee{}

// TransactionRequestFixedFee In the fixed fee model, the transaction fee is a fixed amount within a certain amount of period regardless of the transaction size or network congestion, which can vary between different chains.  You can specify the maximum fee amount to limit the transaction fee. The transaction will fail if the transaction fee exceeds the specified maximum fee amount. 
type TransactionRequestFixedFee struct {
	// The maximum fee that you are willing to pay for the transaction. The transaction will fail if the transaction fee exceeds the maximum fee.
	MaxFeeAmount *string `json:"max_fee_amount,omitempty"`
	FeeType FeeType `json:"fee_type"`
	// The token ID of the transaction fee.
	TokenId string `json:"token_id"`
}

type _TransactionRequestFixedFee TransactionRequestFixedFee

// NewTransactionRequestFixedFee instantiates a new TransactionRequestFixedFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRequestFixedFee(feeType FeeType, tokenId string) *TransactionRequestFixedFee {
	this := TransactionRequestFixedFee{}
	this.FeeType = feeType
	this.TokenId = tokenId
	return &this
}

// NewTransactionRequestFixedFeeWithDefaults instantiates a new TransactionRequestFixedFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRequestFixedFeeWithDefaults() *TransactionRequestFixedFee {
	this := TransactionRequestFixedFee{}
	var feeType FeeType = FEETYPE_EVM_EIP_1559
	this.FeeType = feeType
	return &this
}

// GetMaxFeeAmount returns the MaxFeeAmount field value if set, zero value otherwise.
func (o *TransactionRequestFixedFee) GetMaxFeeAmount() string {
	if o == nil || IsNil(o.MaxFeeAmount) {
		var ret string
		return ret
	}
	return *o.MaxFeeAmount
}

// GetMaxFeeAmountOk returns a tuple with the MaxFeeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRequestFixedFee) GetMaxFeeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.MaxFeeAmount) {
		return nil, false
	}
	return o.MaxFeeAmount, true
}

// HasMaxFeeAmount returns a boolean if a field has been set.
func (o *TransactionRequestFixedFee) HasMaxFeeAmount() bool {
	if o != nil && !IsNil(o.MaxFeeAmount) {
		return true
	}

	return false
}

// SetMaxFeeAmount gets a reference to the given string and assigns it to the MaxFeeAmount field.
func (o *TransactionRequestFixedFee) SetMaxFeeAmount(v string) {
	o.MaxFeeAmount = &v
}

// GetFeeType returns the FeeType field value
func (o *TransactionRequestFixedFee) GetFeeType() FeeType {
	if o == nil {
		var ret FeeType
		return ret
	}

	return o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestFixedFee) GetFeeTypeOk() (*FeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeType, true
}

// SetFeeType sets field value
func (o *TransactionRequestFixedFee) SetFeeType(v FeeType) {
	o.FeeType = v
}

// GetTokenId returns the TokenId field value
func (o *TransactionRequestFixedFee) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestFixedFee) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *TransactionRequestFixedFee) SetTokenId(v string) {
	o.TokenId = v
}

func (o TransactionRequestFixedFee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRequestFixedFee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxFeeAmount) {
		toSerialize["max_fee_amount"] = o.MaxFeeAmount
	}
	toSerialize["fee_type"] = o.FeeType
	toSerialize["token_id"] = o.TokenId
	return toSerialize, nil
}

func (o *TransactionRequestFixedFee) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fee_type",
		"token_id",
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

	varTransactionRequestFixedFee := _TransactionRequestFixedFee{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionRequestFixedFee)

	if err != nil {
		return err
	}

	*o = TransactionRequestFixedFee(varTransactionRequestFixedFee)

	return err
}

type NullableTransactionRequestFixedFee struct {
	value *TransactionRequestFixedFee
	isSet bool
}

func (v NullableTransactionRequestFixedFee) Get() *TransactionRequestFixedFee {
	return v.value
}

func (v *NullableTransactionRequestFixedFee) Set(val *TransactionRequestFixedFee) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRequestFixedFee) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRequestFixedFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRequestFixedFee(val *TransactionRequestFixedFee) *NullableTransactionRequestFixedFee {
	return &NullableTransactionRequestFixedFee{value: val, isSet: true}
}

func (v NullableTransactionRequestFixedFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRequestFixedFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


