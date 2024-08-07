# CreateMpcProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The project name. | 
**Participants** | **int32** | The number of key share holders in the project.  **Notes:** 1. Currently, the available [Threshold Signature Schemes (TSS)](https://manuals.cobo.com/en/portal/mpc-wallets/introduction#threshold-signature-scheme-tss) are 2-2, 2-3, and 3-3 schemes (in the \&quot;threshold - participants\&quot; format), so you can only set &#x60;participants&#x60; to 2 or 3.   2. &#x60;threshold&#x60; must be less than or equal to &#x60;participants&#x60;.  | 
**Threshold** | **int32** | The number of key share holders required to sign an operation in the project.  **Notes:** 1. Currently, the available [Threshold Signature Schemes (TSS)](https://manuals.cobo.com/en/portal/mpc-wallets/introduction#threshold-signature-scheme-tss) are 2-2, 2-3, and 3-3 schemes (in the \&quot;threshold - participants\&quot; format), so you can only set &#x60;threshold&#x60; to 2 or 3.   2. &#x60;threshold&#x60; must be less than or equal to &#x60;participants&#x60;.  | 

## Methods

### NewCreateMpcProjectRequest

`func NewCreateMpcProjectRequest(name string, participants int32, threshold int32, ) *CreateMpcProjectRequest`

NewCreateMpcProjectRequest instantiates a new CreateMpcProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMpcProjectRequestWithDefaults

`func NewCreateMpcProjectRequestWithDefaults() *CreateMpcProjectRequest`

NewCreateMpcProjectRequestWithDefaults instantiates a new CreateMpcProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateMpcProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMpcProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMpcProjectRequest) SetName(v string)`

SetName sets Name field to given value.


### GetParticipants

`func (o *CreateMpcProjectRequest) GetParticipants() int32`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *CreateMpcProjectRequest) GetParticipantsOk() (*int32, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *CreateMpcProjectRequest) SetParticipants(v int32)`

SetParticipants sets Participants field to given value.


### GetThreshold

`func (o *CreateMpcProjectRequest) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *CreateMpcProjectRequest) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *CreateMpcProjectRequest) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


