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

// EstimatedFee - struct for EstimatedFee
type EstimatedFee struct {
	EstimatedEvmEip1559Fee *EstimatedEvmEip1559Fee
	EstimatedEvmLegacyFee *EstimatedEvmLegacyFee
	EstimatedFixedFee *EstimatedFixedFee
	EstimatedUtxoFee *EstimatedUtxoFee
}

// EstimatedEvmEip1559FeeAsEstimatedFee is a convenience function that returns EstimatedEvmEip1559Fee wrapped in EstimatedFee
func EstimatedEvmEip1559FeeAsEstimatedFee(v *EstimatedEvmEip1559Fee) EstimatedFee {
	return EstimatedFee{
		EstimatedEvmEip1559Fee: v,
	}
}

// EstimatedEvmLegacyFeeAsEstimatedFee is a convenience function that returns EstimatedEvmLegacyFee wrapped in EstimatedFee
func EstimatedEvmLegacyFeeAsEstimatedFee(v *EstimatedEvmLegacyFee) EstimatedFee {
	return EstimatedFee{
		EstimatedEvmLegacyFee: v,
	}
}

// EstimatedFixedFeeAsEstimatedFee is a convenience function that returns EstimatedFixedFee wrapped in EstimatedFee
func EstimatedFixedFeeAsEstimatedFee(v *EstimatedFixedFee) EstimatedFee {
	return EstimatedFee{
		EstimatedFixedFee: v,
	}
}

// EstimatedUtxoFeeAsEstimatedFee is a convenience function that returns EstimatedUtxoFee wrapped in EstimatedFee
func EstimatedUtxoFeeAsEstimatedFee(v *EstimatedUtxoFee) EstimatedFee {
	return EstimatedFee{
		EstimatedUtxoFee: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *EstimatedFee) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'EVM_EIP_1559'
	if jsonDict["fee_type"] == "EVM_EIP_1559" {
		// try to unmarshal JSON data into EstimatedEvmEip1559Fee
		err = json.Unmarshal(data, &dst.EstimatedEvmEip1559Fee)
		if err == nil {
			return nil // data stored in dst.EstimatedEvmEip1559Fee, return on the first match
		} else {
			dst.EstimatedEvmEip1559Fee = nil
			return fmt.Errorf("failed to unmarshal EstimatedFee as EstimatedEvmEip1559Fee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EVM_Legacy'
	if jsonDict["fee_type"] == "EVM_Legacy" {
		// try to unmarshal JSON data into EstimatedEvmLegacyFee
		err = json.Unmarshal(data, &dst.EstimatedEvmLegacyFee)
		if err == nil {
			return nil // data stored in dst.EstimatedEvmLegacyFee, return on the first match
		} else {
			dst.EstimatedEvmLegacyFee = nil
			return fmt.Errorf("failed to unmarshal EstimatedFee as EstimatedEvmLegacyFee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Fixed'
	if jsonDict["fee_type"] == "Fixed" {
		// try to unmarshal JSON data into EstimatedFixedFee
		err = json.Unmarshal(data, &dst.EstimatedFixedFee)
		if err == nil {
			return nil // data stored in dst.EstimatedFixedFee, return on the first match
		} else {
			dst.EstimatedFixedFee = nil
			return fmt.Errorf("failed to unmarshal EstimatedFee as EstimatedFixedFee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UTXO'
	if jsonDict["fee_type"] == "UTXO" {
		// try to unmarshal JSON data into EstimatedUtxoFee
		err = json.Unmarshal(data, &dst.EstimatedUtxoFee)
		if err == nil {
			return nil // data stored in dst.EstimatedUtxoFee, return on the first match
		} else {
			dst.EstimatedUtxoFee = nil
			return fmt.Errorf("failed to unmarshal EstimatedFee as EstimatedUtxoFee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EstimatedEvmEip1559Fee'
	if jsonDict["fee_type"] == "EstimatedEvmEip1559Fee" {
		// try to unmarshal JSON data into EstimatedEvmEip1559Fee
		err = json.Unmarshal(data, &dst.EstimatedEvmEip1559Fee)
		if err == nil {
			return nil // data stored in dst.EstimatedEvmEip1559Fee, return on the first match
		} else {
			dst.EstimatedEvmEip1559Fee = nil
			return fmt.Errorf("failed to unmarshal EstimatedFee as EstimatedEvmEip1559Fee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EstimatedEvmLegacyFee'
	if jsonDict["fee_type"] == "EstimatedEvmLegacyFee" {
		// try to unmarshal JSON data into EstimatedEvmLegacyFee
		err = json.Unmarshal(data, &dst.EstimatedEvmLegacyFee)
		if err == nil {
			return nil // data stored in dst.EstimatedEvmLegacyFee, return on the first match
		} else {
			dst.EstimatedEvmLegacyFee = nil
			return fmt.Errorf("failed to unmarshal EstimatedFee as EstimatedEvmLegacyFee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EstimatedFixedFee'
	if jsonDict["fee_type"] == "EstimatedFixedFee" {
		// try to unmarshal JSON data into EstimatedFixedFee
		err = json.Unmarshal(data, &dst.EstimatedFixedFee)
		if err == nil {
			return nil // data stored in dst.EstimatedFixedFee, return on the first match
		} else {
			dst.EstimatedFixedFee = nil
			return fmt.Errorf("failed to unmarshal EstimatedFee as EstimatedFixedFee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EstimatedUtxoFee'
	if jsonDict["fee_type"] == "EstimatedUtxoFee" {
		// try to unmarshal JSON data into EstimatedUtxoFee
		err = json.Unmarshal(data, &dst.EstimatedUtxoFee)
		if err == nil {
			return nil // data stored in dst.EstimatedUtxoFee, return on the first match
		} else {
			dst.EstimatedUtxoFee = nil
			return fmt.Errorf("failed to unmarshal EstimatedFee as EstimatedUtxoFee: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EstimatedFee) MarshalJSON() ([]byte, error) {
	if src.EstimatedEvmEip1559Fee != nil {
		return json.Marshal(&src.EstimatedEvmEip1559Fee)
	}

	if src.EstimatedEvmLegacyFee != nil {
		return json.Marshal(&src.EstimatedEvmLegacyFee)
	}

	if src.EstimatedFixedFee != nil {
		return json.Marshal(&src.EstimatedFixedFee)
	}

	if src.EstimatedUtxoFee != nil {
		return json.Marshal(&src.EstimatedUtxoFee)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EstimatedFee) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.EstimatedEvmEip1559Fee != nil {
		return obj.EstimatedEvmEip1559Fee
	}

	if obj.EstimatedEvmLegacyFee != nil {
		return obj.EstimatedEvmLegacyFee
	}

	if obj.EstimatedFixedFee != nil {
		return obj.EstimatedFixedFee
	}

	if obj.EstimatedUtxoFee != nil {
		return obj.EstimatedUtxoFee
	}

	// all schemas are nil
	return nil
}

type NullableEstimatedFee struct {
	value *EstimatedFee
	isSet bool
}

func (v NullableEstimatedFee) Get() *EstimatedFee {
	return v.value
}

func (v *NullableEstimatedFee) Set(val *EstimatedFee) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimatedFee) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimatedFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimatedFee(val *EstimatedFee) *NullableEstimatedFee {
	return &NullableEstimatedFee{value: val, isSet: true}
}

func (v NullableEstimatedFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimatedFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


