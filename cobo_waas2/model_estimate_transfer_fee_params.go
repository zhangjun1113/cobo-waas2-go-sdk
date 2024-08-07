/*
Cobo Wallet as a Service 2.0

API version: 1.1.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EstimateTransferFeeParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EstimateTransferFeeParams{}

// EstimateTransferFeeParams The information about a token transfer.
type EstimateTransferFeeParams struct {
	// The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization.
	RequestId string `json:"request_id"`
	RequestType EstimateFeeRequestType `json:"request_type"`
	Source TransferSource `json:"source"`
	// The token ID of the transferred token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](/v2/api-references/wallets/list-enabled-tokens).
	TokenId string `json:"token_id"`
	Destination TransferDestination `json:"destination"`
	FeeType *FeeType `json:"fee_type,omitempty"`
}

type _EstimateTransferFeeParams EstimateTransferFeeParams

// NewEstimateTransferFeeParams instantiates a new EstimateTransferFeeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimateTransferFeeParams(requestId string, requestType EstimateFeeRequestType, source TransferSource, tokenId string, destination TransferDestination) *EstimateTransferFeeParams {
	this := EstimateTransferFeeParams{}
	this.RequestId = requestId
	this.RequestType = requestType
	this.Source = source
	this.TokenId = tokenId
	this.Destination = destination
	var feeType FeeType = FEETYPE_EVM_EIP_1559
	this.FeeType = &feeType
	return &this
}

// NewEstimateTransferFeeParamsWithDefaults instantiates a new EstimateTransferFeeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimateTransferFeeParamsWithDefaults() *EstimateTransferFeeParams {
	this := EstimateTransferFeeParams{}
	var feeType FeeType = FEETYPE_EVM_EIP_1559
	this.FeeType = &feeType
	return &this
}

// GetRequestId returns the RequestId field value
func (o *EstimateTransferFeeParams) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *EstimateTransferFeeParams) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *EstimateTransferFeeParams) SetRequestId(v string) {
	o.RequestId = v
}

// GetRequestType returns the RequestType field value
func (o *EstimateTransferFeeParams) GetRequestType() EstimateFeeRequestType {
	if o == nil {
		var ret EstimateFeeRequestType
		return ret
	}

	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *EstimateTransferFeeParams) GetRequestTypeOk() (*EstimateFeeRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value
func (o *EstimateTransferFeeParams) SetRequestType(v EstimateFeeRequestType) {
	o.RequestType = v
}

// GetSource returns the Source field value
func (o *EstimateTransferFeeParams) GetSource() TransferSource {
	if o == nil {
		var ret TransferSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *EstimateTransferFeeParams) GetSourceOk() (*TransferSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *EstimateTransferFeeParams) SetSource(v TransferSource) {
	o.Source = v
}

// GetTokenId returns the TokenId field value
func (o *EstimateTransferFeeParams) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *EstimateTransferFeeParams) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *EstimateTransferFeeParams) SetTokenId(v string) {
	o.TokenId = v
}

// GetDestination returns the Destination field value
func (o *EstimateTransferFeeParams) GetDestination() TransferDestination {
	if o == nil {
		var ret TransferDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *EstimateTransferFeeParams) GetDestinationOk() (*TransferDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *EstimateTransferFeeParams) SetDestination(v TransferDestination) {
	o.Destination = v
}

// GetFeeType returns the FeeType field value if set, zero value otherwise.
func (o *EstimateTransferFeeParams) GetFeeType() FeeType {
	if o == nil || IsNil(o.FeeType) {
		var ret FeeType
		return ret
	}
	return *o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimateTransferFeeParams) GetFeeTypeOk() (*FeeType, bool) {
	if o == nil || IsNil(o.FeeType) {
		return nil, false
	}
	return o.FeeType, true
}

// HasFeeType returns a boolean if a field has been set.
func (o *EstimateTransferFeeParams) HasFeeType() bool {
	if o != nil && !IsNil(o.FeeType) {
		return true
	}

	return false
}

// SetFeeType gets a reference to the given FeeType and assigns it to the FeeType field.
func (o *EstimateTransferFeeParams) SetFeeType(v FeeType) {
	o.FeeType = &v
}

func (o EstimateTransferFeeParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EstimateTransferFeeParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_id"] = o.RequestId
	toSerialize["request_type"] = o.RequestType
	toSerialize["source"] = o.Source
	toSerialize["token_id"] = o.TokenId
	toSerialize["destination"] = o.Destination
	if !IsNil(o.FeeType) {
		toSerialize["fee_type"] = o.FeeType
	}
	return toSerialize, nil
}

func (o *EstimateTransferFeeParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"request_id",
		"request_type",
		"source",
		"token_id",
		"destination",
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

	varEstimateTransferFeeParams := _EstimateTransferFeeParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEstimateTransferFeeParams)

	if err != nil {
		return err
	}

	*o = EstimateTransferFeeParams(varEstimateTransferFeeParams)

	return err
}

type NullableEstimateTransferFeeParams struct {
	value *EstimateTransferFeeParams
	isSet bool
}

func (v NullableEstimateTransferFeeParams) Get() *EstimateTransferFeeParams {
	return v.value
}

func (v *NullableEstimateTransferFeeParams) Set(val *EstimateTransferFeeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimateTransferFeeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimateTransferFeeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimateTransferFeeParams(val *EstimateTransferFeeParams) *NullableEstimateTransferFeeParams {
	return &NullableEstimateTransferFeeParams{value: val, isSet: true}
}

func (v NullableEstimateTransferFeeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimateTransferFeeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


