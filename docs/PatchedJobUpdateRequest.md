# PatchedJobUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to **string** |  | [optional] 
**ExitCode** | Pointer to **int32** |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPatchedJobUpdateRequest

`func NewPatchedJobUpdateRequest() *PatchedJobUpdateRequest`

NewPatchedJobUpdateRequest instantiates a new PatchedJobUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedJobUpdateRequestWithDefaults

`func NewPatchedJobUpdateRequestWithDefaults() *PatchedJobUpdateRequest`

NewPatchedJobUpdateRequestWithDefaults instantiates a new PatchedJobUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *PatchedJobUpdateRequest) GetLogs() string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *PatchedJobUpdateRequest) GetLogsOk() (*string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *PatchedJobUpdateRequest) SetLogs(v string)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *PatchedJobUpdateRequest) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetExitCode

`func (o *PatchedJobUpdateRequest) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *PatchedJobUpdateRequest) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *PatchedJobUpdateRequest) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *PatchedJobUpdateRequest) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetFinishedAt

`func (o *PatchedJobUpdateRequest) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *PatchedJobUpdateRequest) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *PatchedJobUpdateRequest) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *PatchedJobUpdateRequest) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


