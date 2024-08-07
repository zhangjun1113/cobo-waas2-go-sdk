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

// checks if the LockUtxosRequestUtxosInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LockUtxosRequestUtxosInner{}

// LockUtxosRequestUtxosInner struct for LockUtxosRequestUtxosInner
type LockUtxosRequestUtxosInner struct {
	// The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](/v2/api-references/wallets/list-enabled-tokens).
	TokenId string `json:"token_id"`
	// The transaction hash.
	TxHash string `json:"tx_hash"`
	// The output index of the UTXO.
	VoutN int32 `json:"vout_n"`
}

type _LockUtxosRequestUtxosInner LockUtxosRequestUtxosInner

// NewLockUtxosRequestUtxosInner instantiates a new LockUtxosRequestUtxosInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLockUtxosRequestUtxosInner(tokenId string, txHash string, voutN int32) *LockUtxosRequestUtxosInner {
	this := LockUtxosRequestUtxosInner{}
	this.TokenId = tokenId
	this.TxHash = txHash
	this.VoutN = voutN
	return &this
}

// NewLockUtxosRequestUtxosInnerWithDefaults instantiates a new LockUtxosRequestUtxosInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLockUtxosRequestUtxosInnerWithDefaults() *LockUtxosRequestUtxosInner {
	this := LockUtxosRequestUtxosInner{}
	return &this
}

// GetTokenId returns the TokenId field value
func (o *LockUtxosRequestUtxosInner) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *LockUtxosRequestUtxosInner) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *LockUtxosRequestUtxosInner) SetTokenId(v string) {
	o.TokenId = v
}

// GetTxHash returns the TxHash field value
func (o *LockUtxosRequestUtxosInner) GetTxHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value
// and a boolean to check if the value has been set.
func (o *LockUtxosRequestUtxosInner) GetTxHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxHash, true
}

// SetTxHash sets field value
func (o *LockUtxosRequestUtxosInner) SetTxHash(v string) {
	o.TxHash = v
}

// GetVoutN returns the VoutN field value
func (o *LockUtxosRequestUtxosInner) GetVoutN() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VoutN
}

// GetVoutNOk returns a tuple with the VoutN field value
// and a boolean to check if the value has been set.
func (o *LockUtxosRequestUtxosInner) GetVoutNOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VoutN, true
}

// SetVoutN sets field value
func (o *LockUtxosRequestUtxosInner) SetVoutN(v int32) {
	o.VoutN = v
}

func (o LockUtxosRequestUtxosInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LockUtxosRequestUtxosInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token_id"] = o.TokenId
	toSerialize["tx_hash"] = o.TxHash
	toSerialize["vout_n"] = o.VoutN
	return toSerialize, nil
}

func (o *LockUtxosRequestUtxosInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token_id",
		"tx_hash",
		"vout_n",
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

	varLockUtxosRequestUtxosInner := _LockUtxosRequestUtxosInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLockUtxosRequestUtxosInner)

	if err != nil {
		return err
	}

	*o = LockUtxosRequestUtxosInner(varLockUtxosRequestUtxosInner)

	return err
}

type NullableLockUtxosRequestUtxosInner struct {
	value *LockUtxosRequestUtxosInner
	isSet bool
}

func (v NullableLockUtxosRequestUtxosInner) Get() *LockUtxosRequestUtxosInner {
	return v.value
}

func (v *NullableLockUtxosRequestUtxosInner) Set(val *LockUtxosRequestUtxosInner) {
	v.value = val
	v.isSet = true
}

func (v NullableLockUtxosRequestUtxosInner) IsSet() bool {
	return v.isSet
}

func (v *NullableLockUtxosRequestUtxosInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLockUtxosRequestUtxosInner(val *LockUtxosRequestUtxosInner) *NullableLockUtxosRequestUtxosInner {
	return &NullableLockUtxosRequestUtxosInner{value: val, isSet: true}
}

func (v NullableLockUtxosRequestUtxosInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLockUtxosRequestUtxosInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


