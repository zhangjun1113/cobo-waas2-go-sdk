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

// UpdateWalletParams - struct for UpdateWalletParams
type UpdateWalletParams struct {
	UpdateCustodialWalletParams *UpdateCustodialWalletParams
	UpdateExchangeWalletParams *UpdateExchangeWalletParams
	UpdateMpcWalletParams *UpdateMpcWalletParams
	UpdateSmartContractWalletParams *UpdateSmartContractWalletParams
}

// UpdateCustodialWalletParamsAsUpdateWalletParams is a convenience function that returns UpdateCustodialWalletParams wrapped in UpdateWalletParams
func UpdateCustodialWalletParamsAsUpdateWalletParams(v *UpdateCustodialWalletParams) UpdateWalletParams {
	return UpdateWalletParams{
		UpdateCustodialWalletParams: v,
	}
}

// UpdateExchangeWalletParamsAsUpdateWalletParams is a convenience function that returns UpdateExchangeWalletParams wrapped in UpdateWalletParams
func UpdateExchangeWalletParamsAsUpdateWalletParams(v *UpdateExchangeWalletParams) UpdateWalletParams {
	return UpdateWalletParams{
		UpdateExchangeWalletParams: v,
	}
}

// UpdateMpcWalletParamsAsUpdateWalletParams is a convenience function that returns UpdateMpcWalletParams wrapped in UpdateWalletParams
func UpdateMpcWalletParamsAsUpdateWalletParams(v *UpdateMpcWalletParams) UpdateWalletParams {
	return UpdateWalletParams{
		UpdateMpcWalletParams: v,
	}
}

// UpdateSmartContractWalletParamsAsUpdateWalletParams is a convenience function that returns UpdateSmartContractWalletParams wrapped in UpdateWalletParams
func UpdateSmartContractWalletParamsAsUpdateWalletParams(v *UpdateSmartContractWalletParams) UpdateWalletParams {
	return UpdateWalletParams{
		UpdateSmartContractWalletParams: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateWalletParams) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'Custodial'
	if jsonDict["wallet_type"] == "Custodial" {
		// try to unmarshal JSON data into UpdateCustodialWalletParams
		err = json.Unmarshal(data, &dst.UpdateCustodialWalletParams)
		if err == nil {
			return nil // data stored in dst.UpdateCustodialWalletParams, return on the first match
		} else {
			dst.UpdateCustodialWalletParams = nil
			return fmt.Errorf("failed to unmarshal UpdateWalletParams as UpdateCustodialWalletParams: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Exchange'
	if jsonDict["wallet_type"] == "Exchange" {
		// try to unmarshal JSON data into UpdateExchangeWalletParams
		err = json.Unmarshal(data, &dst.UpdateExchangeWalletParams)
		if err == nil {
			return nil // data stored in dst.UpdateExchangeWalletParams, return on the first match
		} else {
			dst.UpdateExchangeWalletParams = nil
			return fmt.Errorf("failed to unmarshal UpdateWalletParams as UpdateExchangeWalletParams: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MPC'
	if jsonDict["wallet_type"] == "MPC" {
		// try to unmarshal JSON data into UpdateMpcWalletParams
		err = json.Unmarshal(data, &dst.UpdateMpcWalletParams)
		if err == nil {
			return nil // data stored in dst.UpdateMpcWalletParams, return on the first match
		} else {
			dst.UpdateMpcWalletParams = nil
			return fmt.Errorf("failed to unmarshal UpdateWalletParams as UpdateMpcWalletParams: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SmartContract'
	if jsonDict["wallet_type"] == "SmartContract" {
		// try to unmarshal JSON data into UpdateSmartContractWalletParams
		err = json.Unmarshal(data, &dst.UpdateSmartContractWalletParams)
		if err == nil {
			return nil // data stored in dst.UpdateSmartContractWalletParams, return on the first match
		} else {
			dst.UpdateSmartContractWalletParams = nil
			return fmt.Errorf("failed to unmarshal UpdateWalletParams as UpdateSmartContractWalletParams: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UpdateCustodialWalletParams'
	if jsonDict["wallet_type"] == "UpdateCustodialWalletParams" {
		// try to unmarshal JSON data into UpdateCustodialWalletParams
		err = json.Unmarshal(data, &dst.UpdateCustodialWalletParams)
		if err == nil {
			return nil // data stored in dst.UpdateCustodialWalletParams, return on the first match
		} else {
			dst.UpdateCustodialWalletParams = nil
			return fmt.Errorf("failed to unmarshal UpdateWalletParams as UpdateCustodialWalletParams: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UpdateExchangeWalletParams'
	if jsonDict["wallet_type"] == "UpdateExchangeWalletParams" {
		// try to unmarshal JSON data into UpdateExchangeWalletParams
		err = json.Unmarshal(data, &dst.UpdateExchangeWalletParams)
		if err == nil {
			return nil // data stored in dst.UpdateExchangeWalletParams, return on the first match
		} else {
			dst.UpdateExchangeWalletParams = nil
			return fmt.Errorf("failed to unmarshal UpdateWalletParams as UpdateExchangeWalletParams: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UpdateMpcWalletParams'
	if jsonDict["wallet_type"] == "UpdateMpcWalletParams" {
		// try to unmarshal JSON data into UpdateMpcWalletParams
		err = json.Unmarshal(data, &dst.UpdateMpcWalletParams)
		if err == nil {
			return nil // data stored in dst.UpdateMpcWalletParams, return on the first match
		} else {
			dst.UpdateMpcWalletParams = nil
			return fmt.Errorf("failed to unmarshal UpdateWalletParams as UpdateMpcWalletParams: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UpdateSmartContractWalletParams'
	if jsonDict["wallet_type"] == "UpdateSmartContractWalletParams" {
		// try to unmarshal JSON data into UpdateSmartContractWalletParams
		err = json.Unmarshal(data, &dst.UpdateSmartContractWalletParams)
		if err == nil {
			return nil // data stored in dst.UpdateSmartContractWalletParams, return on the first match
		} else {
			dst.UpdateSmartContractWalletParams = nil
			return fmt.Errorf("failed to unmarshal UpdateWalletParams as UpdateSmartContractWalletParams: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateWalletParams) MarshalJSON() ([]byte, error) {
	if src.UpdateCustodialWalletParams != nil {
		return json.Marshal(&src.UpdateCustodialWalletParams)
	}

	if src.UpdateExchangeWalletParams != nil {
		return json.Marshal(&src.UpdateExchangeWalletParams)
	}

	if src.UpdateMpcWalletParams != nil {
		return json.Marshal(&src.UpdateMpcWalletParams)
	}

	if src.UpdateSmartContractWalletParams != nil {
		return json.Marshal(&src.UpdateSmartContractWalletParams)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateWalletParams) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.UpdateCustodialWalletParams != nil {
		return obj.UpdateCustodialWalletParams
	}

	if obj.UpdateExchangeWalletParams != nil {
		return obj.UpdateExchangeWalletParams
	}

	if obj.UpdateMpcWalletParams != nil {
		return obj.UpdateMpcWalletParams
	}

	if obj.UpdateSmartContractWalletParams != nil {
		return obj.UpdateSmartContractWalletParams
	}

	// all schemas are nil
	return nil
}

type NullableUpdateWalletParams struct {
	value *UpdateWalletParams
	isSet bool
}

func (v NullableUpdateWalletParams) Get() *UpdateWalletParams {
	return v.value
}

func (v *NullableUpdateWalletParams) Set(val *UpdateWalletParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWalletParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWalletParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWalletParams(val *UpdateWalletParams) *NullableUpdateWalletParams {
	return &NullableUpdateWalletParams{value: val, isSet: true}
}

func (v NullableUpdateWalletParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWalletParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


