# PatchedConstrainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to [**NullableProject**](Project.md) |  | [optional] 
**Setup** | Pointer to [**NullableScript**](Script.md) |  | [optional] 
**Run** | Pointer to [**NullableScript**](Script.md) |  | [optional] 
**Teardown** | Pointer to [**NullableScript**](Script.md) |  | [optional] 
**Resources** | Pointer to [**NullableResources**](Resources.md) |  | [optional] 
**Io** | Pointer to [**NullableIO**](IO.md) |  | [optional] 

## Methods

### NewPatchedConstrainRequest

`func NewPatchedConstrainRequest() *PatchedConstrainRequest`

NewPatchedConstrainRequest instantiates a new PatchedConstrainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedConstrainRequestWithDefaults

`func NewPatchedConstrainRequestWithDefaults() *PatchedConstrainRequest`

NewPatchedConstrainRequestWithDefaults instantiates a new PatchedConstrainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *PatchedConstrainRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PatchedConstrainRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PatchedConstrainRequest) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *PatchedConstrainRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *PatchedConstrainRequest) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *PatchedConstrainRequest) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetSetup

`func (o *PatchedConstrainRequest) GetSetup() Script`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *PatchedConstrainRequest) GetSetupOk() (*Script, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *PatchedConstrainRequest) SetSetup(v Script)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *PatchedConstrainRequest) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### SetSetupNil

`func (o *PatchedConstrainRequest) SetSetupNil(b bool)`

 SetSetupNil sets the value for Setup to be an explicit nil

### UnsetSetup
`func (o *PatchedConstrainRequest) UnsetSetup()`

UnsetSetup ensures that no value is present for Setup, not even an explicit nil
### GetRun

`func (o *PatchedConstrainRequest) GetRun() Script`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *PatchedConstrainRequest) GetRunOk() (*Script, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *PatchedConstrainRequest) SetRun(v Script)`

SetRun sets Run field to given value.

### HasRun

`func (o *PatchedConstrainRequest) HasRun() bool`

HasRun returns a boolean if a field has been set.

### SetRunNil

`func (o *PatchedConstrainRequest) SetRunNil(b bool)`

 SetRunNil sets the value for Run to be an explicit nil

### UnsetRun
`func (o *PatchedConstrainRequest) UnsetRun()`

UnsetRun ensures that no value is present for Run, not even an explicit nil
### GetTeardown

`func (o *PatchedConstrainRequest) GetTeardown() Script`

GetTeardown returns the Teardown field if non-nil, zero value otherwise.

### GetTeardownOk

`func (o *PatchedConstrainRequest) GetTeardownOk() (*Script, bool)`

GetTeardownOk returns a tuple with the Teardown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardown

`func (o *PatchedConstrainRequest) SetTeardown(v Script)`

SetTeardown sets Teardown field to given value.

### HasTeardown

`func (o *PatchedConstrainRequest) HasTeardown() bool`

HasTeardown returns a boolean if a field has been set.

### SetTeardownNil

`func (o *PatchedConstrainRequest) SetTeardownNil(b bool)`

 SetTeardownNil sets the value for Teardown to be an explicit nil

### UnsetTeardown
`func (o *PatchedConstrainRequest) UnsetTeardown()`

UnsetTeardown ensures that no value is present for Teardown, not even an explicit nil
### GetResources

`func (o *PatchedConstrainRequest) GetResources() Resources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *PatchedConstrainRequest) GetResourcesOk() (*Resources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *PatchedConstrainRequest) SetResources(v Resources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *PatchedConstrainRequest) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *PatchedConstrainRequest) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *PatchedConstrainRequest) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil
### GetIo

`func (o *PatchedConstrainRequest) GetIo() IO`

GetIo returns the Io field if non-nil, zero value otherwise.

### GetIoOk

`func (o *PatchedConstrainRequest) GetIoOk() (*IO, bool)`

GetIoOk returns a tuple with the Io field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIo

`func (o *PatchedConstrainRequest) SetIo(v IO)`

SetIo sets Io field to given value.

### HasIo

`func (o *PatchedConstrainRequest) HasIo() bool`

HasIo returns a boolean if a field has been set.

### SetIoNil

`func (o *PatchedConstrainRequest) SetIoNil(b bool)`

 SetIoNil sets the value for Io to be an explicit nil

### UnsetIo
`func (o *PatchedConstrainRequest) UnsetIo()`

UnsetIo ensures that no value is present for Io, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


