# ClusterOffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskSize** | Pointer to **NullableInt32** |  | [optional] 
**Cloud** | Pointer to **NullableString** |  | [optional] 
**JobRecovery** | Pointer to **NullableString** |  | [optional] 
**DiskTier** | Pointer to **NullableString** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**AcceleratorArgs** | Pointer to **interface{}** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **NullableFloat64** |  | [optional] 
**Storage** | Pointer to **NullableString** |  | [optional] 
**InstanceType** | Pointer to **NullableString** |  | [optional] 
**Cpus** | Pointer to **NullableInt32** |  | [optional] 
**Accelerators** | Pointer to **interface{}** |  | [optional] 
**UseSpot** | Pointer to **NullableBool** |  | [optional] 
**Zone** | Pointer to **NullableString** |  | [optional] 
**ImageId** | Pointer to **interface{}** |  | [optional] 
**Ports** | Pointer to **interface{}** |  | [optional] 
**Labels** | Pointer to **interface{}** |  | [optional] 
**AcceleratorCount** | Pointer to **interface{}** |  | [optional] 
**Regions** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewClusterOffer

`func NewClusterOffer() *ClusterOffer`

NewClusterOffer instantiates a new ClusterOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterOfferWithDefaults

`func NewClusterOfferWithDefaults() *ClusterOffer`

NewClusterOfferWithDefaults instantiates a new ClusterOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskSize

`func (o *ClusterOffer) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *ClusterOffer) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *ClusterOffer) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *ClusterOffer) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### SetDiskSizeNil

`func (o *ClusterOffer) SetDiskSizeNil(b bool)`

 SetDiskSizeNil sets the value for DiskSize to be an explicit nil

### UnsetDiskSize
`func (o *ClusterOffer) UnsetDiskSize()`

UnsetDiskSize ensures that no value is present for DiskSize, not even an explicit nil
### GetCloud

`func (o *ClusterOffer) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ClusterOffer) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ClusterOffer) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ClusterOffer) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### SetCloudNil

`func (o *ClusterOffer) SetCloudNil(b bool)`

 SetCloudNil sets the value for Cloud to be an explicit nil

### UnsetCloud
`func (o *ClusterOffer) UnsetCloud()`

UnsetCloud ensures that no value is present for Cloud, not even an explicit nil
### GetJobRecovery

`func (o *ClusterOffer) GetJobRecovery() string`

GetJobRecovery returns the JobRecovery field if non-nil, zero value otherwise.

### GetJobRecoveryOk

`func (o *ClusterOffer) GetJobRecoveryOk() (*string, bool)`

GetJobRecoveryOk returns a tuple with the JobRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRecovery

`func (o *ClusterOffer) SetJobRecovery(v string)`

SetJobRecovery sets JobRecovery field to given value.

### HasJobRecovery

`func (o *ClusterOffer) HasJobRecovery() bool`

HasJobRecovery returns a boolean if a field has been set.

### SetJobRecoveryNil

`func (o *ClusterOffer) SetJobRecoveryNil(b bool)`

 SetJobRecoveryNil sets the value for JobRecovery to be an explicit nil

### UnsetJobRecovery
`func (o *ClusterOffer) UnsetJobRecovery()`

UnsetJobRecovery ensures that no value is present for JobRecovery, not even an explicit nil
### GetDiskTier

`func (o *ClusterOffer) GetDiskTier() string`

GetDiskTier returns the DiskTier field if non-nil, zero value otherwise.

### GetDiskTierOk

`func (o *ClusterOffer) GetDiskTierOk() (*string, bool)`

GetDiskTierOk returns a tuple with the DiskTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskTier

`func (o *ClusterOffer) SetDiskTier(v string)`

SetDiskTier sets DiskTier field to given value.

### HasDiskTier

`func (o *ClusterOffer) HasDiskTier() bool`

HasDiskTier returns a boolean if a field has been set.

### SetDiskTierNil

`func (o *ClusterOffer) SetDiskTierNil(b bool)`

 SetDiskTierNil sets the value for DiskTier to be an explicit nil

### UnsetDiskTier
`func (o *ClusterOffer) UnsetDiskTier()`

UnsetDiskTier ensures that no value is present for DiskTier, not even an explicit nil
### GetMemory

`func (o *ClusterOffer) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ClusterOffer) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ClusterOffer) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ClusterOffer) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *ClusterOffer) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *ClusterOffer) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetAcceleratorArgs

`func (o *ClusterOffer) GetAcceleratorArgs() interface{}`

GetAcceleratorArgs returns the AcceleratorArgs field if non-nil, zero value otherwise.

### GetAcceleratorArgsOk

`func (o *ClusterOffer) GetAcceleratorArgsOk() (*interface{}, bool)`

GetAcceleratorArgsOk returns a tuple with the AcceleratorArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleratorArgs

`func (o *ClusterOffer) SetAcceleratorArgs(v interface{})`

SetAcceleratorArgs sets AcceleratorArgs field to given value.

### HasAcceleratorArgs

`func (o *ClusterOffer) HasAcceleratorArgs() bool`

HasAcceleratorArgs returns a boolean if a field has been set.

### SetAcceleratorArgsNil

`func (o *ClusterOffer) SetAcceleratorArgsNil(b bool)`

 SetAcceleratorArgsNil sets the value for AcceleratorArgs to be an explicit nil

### UnsetAcceleratorArgs
`func (o *ClusterOffer) UnsetAcceleratorArgs()`

UnsetAcceleratorArgs ensures that no value is present for AcceleratorArgs, not even an explicit nil
### GetRegion

`func (o *ClusterOffer) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClusterOffer) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClusterOffer) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ClusterOffer) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *ClusterOffer) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *ClusterOffer) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetPrice

`func (o *ClusterOffer) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ClusterOffer) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ClusterOffer) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ClusterOffer) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ClusterOffer) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ClusterOffer) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetStorage

`func (o *ClusterOffer) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ClusterOffer) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ClusterOffer) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ClusterOffer) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *ClusterOffer) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *ClusterOffer) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetInstanceType

`func (o *ClusterOffer) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ClusterOffer) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ClusterOffer) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ClusterOffer) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### SetInstanceTypeNil

`func (o *ClusterOffer) SetInstanceTypeNil(b bool)`

 SetInstanceTypeNil sets the value for InstanceType to be an explicit nil

### UnsetInstanceType
`func (o *ClusterOffer) UnsetInstanceType()`

UnsetInstanceType ensures that no value is present for InstanceType, not even an explicit nil
### GetCpus

`func (o *ClusterOffer) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *ClusterOffer) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *ClusterOffer) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *ClusterOffer) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### SetCpusNil

`func (o *ClusterOffer) SetCpusNil(b bool)`

 SetCpusNil sets the value for Cpus to be an explicit nil

### UnsetCpus
`func (o *ClusterOffer) UnsetCpus()`

UnsetCpus ensures that no value is present for Cpus, not even an explicit nil
### GetAccelerators

`func (o *ClusterOffer) GetAccelerators() interface{}`

GetAccelerators returns the Accelerators field if non-nil, zero value otherwise.

### GetAcceleratorsOk

`func (o *ClusterOffer) GetAcceleratorsOk() (*interface{}, bool)`

GetAcceleratorsOk returns a tuple with the Accelerators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerators

`func (o *ClusterOffer) SetAccelerators(v interface{})`

SetAccelerators sets Accelerators field to given value.

### HasAccelerators

`func (o *ClusterOffer) HasAccelerators() bool`

HasAccelerators returns a boolean if a field has been set.

### SetAcceleratorsNil

`func (o *ClusterOffer) SetAcceleratorsNil(b bool)`

 SetAcceleratorsNil sets the value for Accelerators to be an explicit nil

### UnsetAccelerators
`func (o *ClusterOffer) UnsetAccelerators()`

UnsetAccelerators ensures that no value is present for Accelerators, not even an explicit nil
### GetUseSpot

`func (o *ClusterOffer) GetUseSpot() bool`

GetUseSpot returns the UseSpot field if non-nil, zero value otherwise.

### GetUseSpotOk

`func (o *ClusterOffer) GetUseSpotOk() (*bool, bool)`

GetUseSpotOk returns a tuple with the UseSpot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSpot

`func (o *ClusterOffer) SetUseSpot(v bool)`

SetUseSpot sets UseSpot field to given value.

### HasUseSpot

`func (o *ClusterOffer) HasUseSpot() bool`

HasUseSpot returns a boolean if a field has been set.

### SetUseSpotNil

`func (o *ClusterOffer) SetUseSpotNil(b bool)`

 SetUseSpotNil sets the value for UseSpot to be an explicit nil

### UnsetUseSpot
`func (o *ClusterOffer) UnsetUseSpot()`

UnsetUseSpot ensures that no value is present for UseSpot, not even an explicit nil
### GetZone

`func (o *ClusterOffer) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ClusterOffer) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ClusterOffer) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ClusterOffer) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *ClusterOffer) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *ClusterOffer) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetImageId

`func (o *ClusterOffer) GetImageId() interface{}`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ClusterOffer) GetImageIdOk() (*interface{}, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ClusterOffer) SetImageId(v interface{})`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *ClusterOffer) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageIdNil

`func (o *ClusterOffer) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *ClusterOffer) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
### GetPorts

`func (o *ClusterOffer) GetPorts() interface{}`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ClusterOffer) GetPortsOk() (*interface{}, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ClusterOffer) SetPorts(v interface{})`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ClusterOffer) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *ClusterOffer) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *ClusterOffer) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetLabels

`func (o *ClusterOffer) GetLabels() interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClusterOffer) GetLabelsOk() (*interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClusterOffer) SetLabels(v interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ClusterOffer) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ClusterOffer) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ClusterOffer) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetAcceleratorCount

`func (o *ClusterOffer) GetAcceleratorCount() interface{}`

GetAcceleratorCount returns the AcceleratorCount field if non-nil, zero value otherwise.

### GetAcceleratorCountOk

`func (o *ClusterOffer) GetAcceleratorCountOk() (*interface{}, bool)`

GetAcceleratorCountOk returns a tuple with the AcceleratorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleratorCount

`func (o *ClusterOffer) SetAcceleratorCount(v interface{})`

SetAcceleratorCount sets AcceleratorCount field to given value.

### HasAcceleratorCount

`func (o *ClusterOffer) HasAcceleratorCount() bool`

HasAcceleratorCount returns a boolean if a field has been set.

### SetAcceleratorCountNil

`func (o *ClusterOffer) SetAcceleratorCountNil(b bool)`

 SetAcceleratorCountNil sets the value for AcceleratorCount to be an explicit nil

### UnsetAcceleratorCount
`func (o *ClusterOffer) UnsetAcceleratorCount()`

UnsetAcceleratorCount ensures that no value is present for AcceleratorCount, not even an explicit nil
### GetRegions

`func (o *ClusterOffer) GetRegions() interface{}`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ClusterOffer) GetRegionsOk() (*interface{}, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ClusterOffer) SetRegions(v interface{})`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *ClusterOffer) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### SetRegionsNil

`func (o *ClusterOffer) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *ClusterOffer) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


