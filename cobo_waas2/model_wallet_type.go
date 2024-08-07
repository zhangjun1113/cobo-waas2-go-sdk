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

// WalletType The wallet type.  - `Custodial`: [Custodial Wallets](https://manuals.cobo.com/en/portal/custodial-wallets/introduction)  - `MPC`: [MPC Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/introduction)  - `SmartContract`: [Smart Contract Wallets](https://manuals.cobo.com/en/portal/smart-contract-wallets/introduction)  - `Exchange`: [Exchange Wallets](https://manuals.cobo.com/en/portal/exchange-wallets/introduction) 
type WalletType string

// List of WalletType
const (
	WALLETTYPE_CUSTODIAL WalletType = "Custodial"
	WALLETTYPE_MPC WalletType = "MPC"
	WALLETTYPE_SMART_CONTRACT WalletType = "SmartContract"
	WALLETTYPE_EXCHANGE WalletType = "Exchange"
)

// All allowed values of WalletType enum
var AllowedWalletTypeEnumValues = []WalletType{
	"Custodial",
	"MPC",
	"SmartContract",
	"Exchange",
}

func (v *WalletType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WalletType(value)
	for _, existing := range AllowedWalletTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = WalletType("unknown")
    return nil
}

// NewWalletTypeFromValue returns a pointer to a valid WalletType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWalletTypeFromValue(v string) (*WalletType, error) {
	ev := WalletType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WalletType: valid values are %v", v, AllowedWalletTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WalletType) IsValid() bool {
	for _, existing := range AllowedWalletTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WalletType value
func (v WalletType) Ptr() *WalletType {
	return &v
}

type NullableWalletType struct {
	value *WalletType
	isSet bool
}

func (v NullableWalletType) Get() *WalletType {
	return v.value
}

func (v *NullableWalletType) Set(val *WalletType) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletType) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletType(val *WalletType) *NullableWalletType {
	return &NullableWalletType{value: val, isSet: true}
}

func (v NullableWalletType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

