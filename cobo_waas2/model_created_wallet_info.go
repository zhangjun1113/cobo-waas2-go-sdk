/*
Cobo Wallet as a Service 2.0

API version: 1.1.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"fmt"
)

// CreatedWalletInfo - struct for CreatedWalletInfo
type CreatedWalletInfo struct {
	CustodialWalletInfo *CustodialWalletInfo
	ExchangeWalletInfo *ExchangeWalletInfo
	MPCWalletInfo *MPCWalletInfo
}

// CustodialWalletInfoAsCreatedWalletInfo is a convenience function that returns CustodialWalletInfo wrapped in CreatedWalletInfo
func CustodialWalletInfoAsCreatedWalletInfo(v *CustodialWalletInfo) CreatedWalletInfo {
	return CreatedWalletInfo{
		CustodialWalletInfo: v,
	}
}

// ExchangeWalletInfoAsCreatedWalletInfo is a convenience function that returns ExchangeWalletInfo wrapped in CreatedWalletInfo
func ExchangeWalletInfoAsCreatedWalletInfo(v *ExchangeWalletInfo) CreatedWalletInfo {
	return CreatedWalletInfo{
		ExchangeWalletInfo: v,
	}
}

// MPCWalletInfoAsCreatedWalletInfo is a convenience function that returns MPCWalletInfo wrapped in CreatedWalletInfo
func MPCWalletInfoAsCreatedWalletInfo(v *MPCWalletInfo) CreatedWalletInfo {
	return CreatedWalletInfo{
		MPCWalletInfo: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreatedWalletInfo) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'Custodial'
	if jsonDict["wallet_type"] == "Custodial" {
		// try to unmarshal JSON data into CustodialWalletInfo
		err = json.Unmarshal(data, &dst.CustodialWalletInfo)
		if err == nil {
			return nil // data stored in dst.CustodialWalletInfo, return on the first match
		} else {
			dst.CustodialWalletInfo = nil
			return fmt.Errorf("failed to unmarshal CreatedWalletInfo as CustodialWalletInfo: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Exchange'
	if jsonDict["wallet_type"] == "Exchange" {
		// try to unmarshal JSON data into ExchangeWalletInfo
		err = json.Unmarshal(data, &dst.ExchangeWalletInfo)
		if err == nil {
			return nil // data stored in dst.ExchangeWalletInfo, return on the first match
		} else {
			dst.ExchangeWalletInfo = nil
			return fmt.Errorf("failed to unmarshal CreatedWalletInfo as ExchangeWalletInfo: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MPC'
	if jsonDict["wallet_type"] == "MPC" {
		// try to unmarshal JSON data into MPCWalletInfo
		err = json.Unmarshal(data, &dst.MPCWalletInfo)
		if err == nil {
			return nil // data stored in dst.MPCWalletInfo, return on the first match
		} else {
			dst.MPCWalletInfo = nil
			return fmt.Errorf("failed to unmarshal CreatedWalletInfo as MPCWalletInfo: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CustodialWalletInfo'
	if jsonDict["wallet_type"] == "CustodialWalletInfo" {
		// try to unmarshal JSON data into CustodialWalletInfo
		err = json.Unmarshal(data, &dst.CustodialWalletInfo)
		if err == nil {
			return nil // data stored in dst.CustodialWalletInfo, return on the first match
		} else {
			dst.CustodialWalletInfo = nil
			return fmt.Errorf("failed to unmarshal CreatedWalletInfo as CustodialWalletInfo: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ExchangeWalletInfo'
	if jsonDict["wallet_type"] == "ExchangeWalletInfo" {
		// try to unmarshal JSON data into ExchangeWalletInfo
		err = json.Unmarshal(data, &dst.ExchangeWalletInfo)
		if err == nil {
			return nil // data stored in dst.ExchangeWalletInfo, return on the first match
		} else {
			dst.ExchangeWalletInfo = nil
			return fmt.Errorf("failed to unmarshal CreatedWalletInfo as ExchangeWalletInfo: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MPCWalletInfo'
	if jsonDict["wallet_type"] == "MPCWalletInfo" {
		// try to unmarshal JSON data into MPCWalletInfo
		err = json.Unmarshal(data, &dst.MPCWalletInfo)
		if err == nil {
			return nil // data stored in dst.MPCWalletInfo, return on the first match
		} else {
			dst.MPCWalletInfo = nil
			return fmt.Errorf("failed to unmarshal CreatedWalletInfo as MPCWalletInfo: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreatedWalletInfo) MarshalJSON() ([]byte, error) {
	if src.CustodialWalletInfo != nil {
		return json.Marshal(&src.CustodialWalletInfo)
	}

	if src.ExchangeWalletInfo != nil {
		return json.Marshal(&src.ExchangeWalletInfo)
	}

	if src.MPCWalletInfo != nil {
		return json.Marshal(&src.MPCWalletInfo)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreatedWalletInfo) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CustodialWalletInfo != nil {
		return obj.CustodialWalletInfo
	}

	if obj.ExchangeWalletInfo != nil {
		return obj.ExchangeWalletInfo
	}

	if obj.MPCWalletInfo != nil {
		return obj.MPCWalletInfo
	}

	// all schemas are nil
	return nil
}

type NullableCreatedWalletInfo struct {
	value *CreatedWalletInfo
	isSet bool
}

func (v NullableCreatedWalletInfo) Get() *CreatedWalletInfo {
	return v.value
}

func (v *NullableCreatedWalletInfo) Set(val *CreatedWalletInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatedWalletInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatedWalletInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatedWalletInfo(val *CreatedWalletInfo) *NullableCreatedWalletInfo {
	return &NullableCreatedWalletInfo{value: val, isSet: true}
}

func (v NullableCreatedWalletInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatedWalletInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


