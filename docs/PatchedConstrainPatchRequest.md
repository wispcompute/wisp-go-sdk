# PatchedConstrainPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**ConstrainRequest**](ConstrainRequest.md) |  | [optional] 

## Methods

### NewPatchedConstrainPatchRequest

`func NewPatchedConstrainPatchRequest() *PatchedConstrainPatchRequest`

NewPatchedConstrainPatchRequest instantiates a new PatchedConstrainPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedConstrainPatchRequestWithDefaults

`func NewPatchedConstrainPatchRequestWithDefaults() *PatchedConstrainPatchRequest`

NewPatchedConstrainPatchRequestWithDefaults instantiates a new PatchedConstrainPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *PatchedConstrainPatchRequest) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *PatchedConstrainPatchRequest) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *PatchedConstrainPatchRequest) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *PatchedConstrainPatchRequest) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetConfig

`func (o *PatchedConstrainPatchRequest) GetConfig() ConstrainRequest`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PatchedConstrainPatchRequest) GetConfigOk() (*ConstrainRequest, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PatchedConstrainPatchRequest) SetConfig(v ConstrainRequest)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *PatchedConstrainPatchRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


