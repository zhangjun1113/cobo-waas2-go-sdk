/*
Cobo Wallet as a Service 2.0

Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"fmt"
)

// TransactionResult - struct for TransactionResult
type TransactionResult struct {
	TransactionSignatureResult *TransactionSignatureResult
}

// TransactionSignatureResultAsTransactionResult is a convenience function that returns TransactionSignatureResult wrapped in TransactionResult
func TransactionSignatureResultAsTransactionResult(v *TransactionSignatureResult) TransactionResult {
	return TransactionResult{
		TransactionSignatureResult: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TransactionResult) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'Signature'
	if jsonDict["result_type"] == "Signature" {
		// try to unmarshal JSON data into TransactionSignatureResult
		err = json.Unmarshal(data, &dst.TransactionSignatureResult)
		if err == nil {
			return nil // data stored in dst.TransactionSignatureResult, return on the first match
		} else {
			dst.TransactionSignatureResult = nil
			return fmt.Errorf("failed to unmarshal TransactionResult as TransactionSignatureResult: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionSignatureResult'
	if jsonDict["result_type"] == "TransactionSignatureResult" {
		// try to unmarshal JSON data into TransactionSignatureResult
		err = json.Unmarshal(data, &dst.TransactionSignatureResult)
		if err == nil {
			return nil // data stored in dst.TransactionSignatureResult, return on the first match
		} else {
			dst.TransactionSignatureResult = nil
			return fmt.Errorf("failed to unmarshal TransactionResult as TransactionSignatureResult: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransactionResult) MarshalJSON() ([]byte, error) {
	if src.TransactionSignatureResult != nil {
		return json.Marshal(&src.TransactionSignatureResult)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TransactionResult) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.TransactionSignatureResult != nil {
		return obj.TransactionSignatureResult
	}

	// all schemas are nil
	return nil
}

type NullableTransactionResult struct {
	value *TransactionResult
	isSet bool
}

func (v NullableTransactionResult) Get() *TransactionResult {
	return v.value
}

func (v *NullableTransactionResult) Set(val *TransactionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionResult(val *TransactionResult) *NullableTransactionResult {
	return &NullableTransactionResult{value: val, isSet: true}
}

func (v NullableTransactionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


