# ConstrainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | [**Project**](Project.md) |  | [readonly] 
**Setup** | Pointer to [**NullableScript**](Script.md) |  | [optional] 
**Run** | Pointer to [**NullableScript**](Script.md) |  | [optional] 
**Teardown** | Pointer to [**NullableScript**](Script.md) |  | [optional] 
**Resources** | [**Resources**](Resources.md) |  | 
**Io** | [**NullableIO**](IO.md) |  | 

## Methods

### NewConstrainRequest

`func NewConstrainRequest(project Project, resources Resources, io NullableIO, ) *ConstrainRequest`

NewConstrainRequest instantiates a new ConstrainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstrainRequestWithDefaults

`func NewConstrainRequestWithDefaults() *ConstrainRequest`

NewConstrainRequestWithDefaults instantiates a new ConstrainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ConstrainRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ConstrainRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ConstrainRequest) SetProject(v Project)`

SetProject sets Project field to given value.


### GetSetup

`func (o *ConstrainRequest) GetSetup() Script`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *ConstrainRequest) GetSetupOk() (*Script, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *ConstrainRequest) SetSetup(v Script)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *ConstrainRequest) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### SetSetupNil

`func (o *ConstrainRequest) SetSetupNil(b bool)`

 SetSetupNil sets the value for Setup to be an explicit nil

### UnsetSetup
`func (o *ConstrainRequest) UnsetSetup()`

UnsetSetup ensures that no value is present for Setup, not even an explicit nil
### GetRun

`func (o *ConstrainRequest) GetRun() Script`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *ConstrainRequest) GetRunOk() (*Script, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *ConstrainRequest) SetRun(v Script)`

SetRun sets Run field to given value.

### HasRun

`func (o *ConstrainRequest) HasRun() bool`

HasRun returns a boolean if a field has been set.

### SetRunNil

`func (o *ConstrainRequest) SetRunNil(b bool)`

 SetRunNil sets the value for Run to be an explicit nil

### UnsetRun
`func (o *ConstrainRequest) UnsetRun()`

UnsetRun ensures that no value is present for Run, not even an explicit nil
### GetTeardown

`func (o *ConstrainRequest) GetTeardown() Script`

GetTeardown returns the Teardown field if non-nil, zero value otherwise.

### GetTeardownOk

`func (o *ConstrainRequest) GetTeardownOk() (*Script, bool)`

GetTeardownOk returns a tuple with the Teardown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardown

`func (o *ConstrainRequest) SetTeardown(v Script)`

SetTeardown sets Teardown field to given value.

### HasTeardown

`func (o *ConstrainRequest) HasTeardown() bool`

HasTeardown returns a boolean if a field has been set.

### SetTeardownNil

`func (o *ConstrainRequest) SetTeardownNil(b bool)`

 SetTeardownNil sets the value for Teardown to be an explicit nil

### UnsetTeardown
`func (o *ConstrainRequest) UnsetTeardown()`

UnsetTeardown ensures that no value is present for Teardown, not even an explicit nil
### GetResources

`func (o *ConstrainRequest) GetResources() Resources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ConstrainRequest) GetResourcesOk() (*Resources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ConstrainRequest) SetResources(v Resources)`

SetResources sets Resources field to given value.


### GetIo

`func (o *ConstrainRequest) GetIo() IO`

GetIo returns the Io field if non-nil, zero value otherwise.

### GetIoOk

`func (o *ConstrainRequest) GetIoOk() (*IO, bool)`

GetIoOk returns a tuple with the Io field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIo

`func (o *ConstrainRequest) SetIo(v IO)`

SetIo sets Io field to given value.


### SetIoNil

`func (o *ConstrainRequest) SetIoNil(b bool)`

 SetIoNil sets the value for Io to be an explicit nil

### UnsetIo
`func (o *ConstrainRequest) UnsetIo()`

UnsetIo ensures that no value is present for Io, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


