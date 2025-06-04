# Resources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clouds** | Pointer to **[]string** |  | [optional] 
**InstanceTypes** | Pointer to **[]string** |  | [optional] 
**DockerImage** | Pointer to **NullableString** |  | [optional] 
**Regions** | Pointer to **[]string** |  | [optional] 
**Areas** | Pointer to **[]string** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**Cpus** | Pointer to **NullableInt32** |  | [optional] 
**Storage** | Pointer to **NullableInt32** |  | [optional] 
**PersistentDisk** | Pointer to **NullableInt32** |  | [optional] 
**Accelerators** | **[]string** |  | 
**ComputeCapability** | Pointer to **NullableString** |  | [optional] 
**Vram** | Pointer to **NullableInt32** |  | [optional] 
**AcceleratorCount** | Pointer to **NullableInt32** |  | [optional] 
**Platform** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewResources

`func NewResources(accelerators []*string, ) *Resources`

NewResources instantiates a new Resources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesWithDefaults

`func NewResourcesWithDefaults() *Resources`

NewResourcesWithDefaults instantiates a new Resources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClouds

`func (o *Resources) GetClouds() []*string`

GetClouds returns the Clouds field if non-nil, zero value otherwise.

### GetCloudsOk

`func (o *Resources) GetCloudsOk() (*[]*string, bool)`

GetCloudsOk returns a tuple with the Clouds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClouds

`func (o *Resources) SetClouds(v []*string)`

SetClouds sets Clouds field to given value.

### HasClouds

`func (o *Resources) HasClouds() bool`

HasClouds returns a boolean if a field has been set.

### GetInstanceTypes

`func (o *Resources) GetInstanceTypes() []*string`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *Resources) GetInstanceTypesOk() (*[]*string, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *Resources) SetInstanceTypes(v []*string)`

SetInstanceTypes sets InstanceTypes field to given value.

### HasInstanceTypes

`func (o *Resources) HasInstanceTypes() bool`

HasInstanceTypes returns a boolean if a field has been set.

### GetDockerImage

`func (o *Resources) GetDockerImage() string`

GetDockerImage returns the DockerImage field if non-nil, zero value otherwise.

### GetDockerImageOk

`func (o *Resources) GetDockerImageOk() (*string, bool)`

GetDockerImageOk returns a tuple with the DockerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImage

`func (o *Resources) SetDockerImage(v string)`

SetDockerImage sets DockerImage field to given value.

### HasDockerImage

`func (o *Resources) HasDockerImage() bool`

HasDockerImage returns a boolean if a field has been set.

### SetDockerImageNil

`func (o *Resources) SetDockerImageNil(b bool)`

 SetDockerImageNil sets the value for DockerImage to be an explicit nil

### UnsetDockerImage
`func (o *Resources) UnsetDockerImage()`

UnsetDockerImage ensures that no value is present for DockerImage, not even an explicit nil
### GetRegions

`func (o *Resources) GetRegions() []*string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Resources) GetRegionsOk() (*[]*string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Resources) SetRegions(v []*string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Resources) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetAreas

`func (o *Resources) GetAreas() []*string`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *Resources) GetAreasOk() (*[]*string, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *Resources) SetAreas(v []*string)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *Resources) HasAreas() bool`

HasAreas returns a boolean if a field has been set.

### GetMemory

`func (o *Resources) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Resources) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Resources) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Resources) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *Resources) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *Resources) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetCpus

`func (o *Resources) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *Resources) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *Resources) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *Resources) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### SetCpusNil

`func (o *Resources) SetCpusNil(b bool)`

 SetCpusNil sets the value for Cpus to be an explicit nil

### UnsetCpus
`func (o *Resources) UnsetCpus()`

UnsetCpus ensures that no value is present for Cpus, not even an explicit nil
### GetStorage

`func (o *Resources) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *Resources) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *Resources) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *Resources) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *Resources) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *Resources) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetPersistentDisk

`func (o *Resources) GetPersistentDisk() int32`

GetPersistentDisk returns the PersistentDisk field if non-nil, zero value otherwise.

### GetPersistentDiskOk

`func (o *Resources) GetPersistentDiskOk() (*int32, bool)`

GetPersistentDiskOk returns a tuple with the PersistentDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentDisk

`func (o *Resources) SetPersistentDisk(v int32)`

SetPersistentDisk sets PersistentDisk field to given value.

### HasPersistentDisk

`func (o *Resources) HasPersistentDisk() bool`

HasPersistentDisk returns a boolean if a field has been set.

### SetPersistentDiskNil

`func (o *Resources) SetPersistentDiskNil(b bool)`

 SetPersistentDiskNil sets the value for PersistentDisk to be an explicit nil

### UnsetPersistentDisk
`func (o *Resources) UnsetPersistentDisk()`

UnsetPersistentDisk ensures that no value is present for PersistentDisk, not even an explicit nil
### GetAccelerators

`func (o *Resources) GetAccelerators() []*string`

GetAccelerators returns the Accelerators field if non-nil, zero value otherwise.

### GetAcceleratorsOk

`func (o *Resources) GetAcceleratorsOk() (*[]*string, bool)`

GetAcceleratorsOk returns a tuple with the Accelerators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerators

`func (o *Resources) SetAccelerators(v []*string)`

SetAccelerators sets Accelerators field to given value.


### SetAcceleratorsNil

`func (o *Resources) SetAcceleratorsNil(b bool)`

 SetAcceleratorsNil sets the value for Accelerators to be an explicit nil

### UnsetAccelerators
`func (o *Resources) UnsetAccelerators()`

UnsetAccelerators ensures that no value is present for Accelerators, not even an explicit nil
### GetComputeCapability

`func (o *Resources) GetComputeCapability() string`

GetComputeCapability returns the ComputeCapability field if non-nil, zero value otherwise.

### GetComputeCapabilityOk

`func (o *Resources) GetComputeCapabilityOk() (*string, bool)`

GetComputeCapabilityOk returns a tuple with the ComputeCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeCapability

`func (o *Resources) SetComputeCapability(v string)`

SetComputeCapability sets ComputeCapability field to given value.

### HasComputeCapability

`func (o *Resources) HasComputeCapability() bool`

HasComputeCapability returns a boolean if a field has been set.

### SetComputeCapabilityNil

`func (o *Resources) SetComputeCapabilityNil(b bool)`

 SetComputeCapabilityNil sets the value for ComputeCapability to be an explicit nil

### UnsetComputeCapability
`func (o *Resources) UnsetComputeCapability()`

UnsetComputeCapability ensures that no value is present for ComputeCapability, not even an explicit nil
### GetVram

`func (o *Resources) GetVram() int32`

GetVram returns the Vram field if non-nil, zero value otherwise.

### GetVramOk

`func (o *Resources) GetVramOk() (*int32, bool)`

GetVramOk returns a tuple with the Vram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVram

`func (o *Resources) SetVram(v int32)`

SetVram sets Vram field to given value.

### HasVram

`func (o *Resources) HasVram() bool`

HasVram returns a boolean if a field has been set.

### SetVramNil

`func (o *Resources) SetVramNil(b bool)`

 SetVramNil sets the value for Vram to be an explicit nil

### UnsetVram
`func (o *Resources) UnsetVram()`

UnsetVram ensures that no value is present for Vram, not even an explicit nil
### GetAcceleratorCount

`func (o *Resources) GetAcceleratorCount() int32`

GetAcceleratorCount returns the AcceleratorCount field if non-nil, zero value otherwise.

### GetAcceleratorCountOk

`func (o *Resources) GetAcceleratorCountOk() (*int32, bool)`

GetAcceleratorCountOk returns a tuple with the AcceleratorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleratorCount

`func (o *Resources) SetAcceleratorCount(v int32)`

SetAcceleratorCount sets AcceleratorCount field to given value.

### HasAcceleratorCount

`func (o *Resources) HasAcceleratorCount() bool`

HasAcceleratorCount returns a boolean if a field has been set.

### SetAcceleratorCountNil

`func (o *Resources) SetAcceleratorCountNil(b bool)`

 SetAcceleratorCountNil sets the value for AcceleratorCount to be an explicit nil

### UnsetAcceleratorCount
`func (o *Resources) UnsetAcceleratorCount()`

UnsetAcceleratorCount ensures that no value is present for AcceleratorCount, not even an explicit nil
### GetPlatform

`func (o *Resources) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Resources) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Resources) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *Resources) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *Resources) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *Resources) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


