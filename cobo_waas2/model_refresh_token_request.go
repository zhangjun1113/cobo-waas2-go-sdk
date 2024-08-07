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

// checks if the RefreshTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshTokenRequest{}

// RefreshTokenRequest struct for RefreshTokenRequest
type RefreshTokenRequest struct {
	// The App ID, a unique identifier to distinguish Cobo Portal Apps. You can get the App ID by retrieving the Manifest file after receiving the notification of app launch approval.
	ClientId *string `json:"client_id,omitempty"`
	// The type of the permission granting. To refresh an access token, you need to set the value as `refresh_token`.
	GrantType *string `json:"grant_type,omitempty"`
	// The refresh token of the expired or expiring access token.
	RefreshToken *string `json:"refresh_token,omitempty"`
}

// NewRefreshTokenRequest instantiates a new RefreshTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshTokenRequest() *RefreshTokenRequest {
	this := RefreshTokenRequest{}
	return &this
}

// NewRefreshTokenRequestWithDefaults instantiates a new RefreshTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshTokenRequestWithDefaults() *RefreshTokenRequest {
	this := RefreshTokenRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *RefreshTokenRequest) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenRequest) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *RefreshTokenRequest) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *RefreshTokenRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetGrantType returns the GrantType field value if set, zero value otherwise.
func (o *RefreshTokenRequest) GetGrantType() string {
	if o == nil || IsNil(o.GrantType) {
		var ret string
		return ret
	}
	return *o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenRequest) GetGrantTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GrantType) {
		return nil, false
	}
	return o.GrantType, true
}

// HasGrantType returns a boolean if a field has been set.
func (o *RefreshTokenRequest) HasGrantType() bool {
	if o != nil && !IsNil(o.GrantType) {
		return true
	}

	return false
}

// SetGrantType gets a reference to the given string and assigns it to the GrantType field.
func (o *RefreshTokenRequest) SetGrantType(v string) {
	o.GrantType = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *RefreshTokenRequest) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenRequest) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *RefreshTokenRequest) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *RefreshTokenRequest) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

func (o RefreshTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.GrantType) {
		toSerialize["grant_type"] = o.GrantType
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	return toSerialize, nil
}

type NullableRefreshTokenRequest struct {
	value *RefreshTokenRequest
	isSet bool
}

func (v NullableRefreshTokenRequest) Get() *RefreshTokenRequest {
	return v.value
}

func (v *NullableRefreshTokenRequest) Set(val *RefreshTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshTokenRequest(val *RefreshTokenRequest) *NullableRefreshTokenRequest {
	return &NullableRefreshTokenRequest{value: val, isSet: true}
}

func (v NullableRefreshTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


