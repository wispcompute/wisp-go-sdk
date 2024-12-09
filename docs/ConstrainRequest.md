# ConstrainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accelerators** | Pointer to **map[string]interface{}** |  | [optional] 
**Storage** | Pointer to **int32** |  | [optional] 
**Memory** | Pointer to **int32** |  | [optional] 
**VCPUs** | Pointer to **int32** |  | [optional] 
**Regions** | Pointer to **[]string** |  | [optional] 
**Clouds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewConstrainRequest

`func NewConstrainRequest() *ConstrainRequest`

NewConstrainRequest instantiates a new ConstrainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstrainRequestWithDefaults

`func NewConstrainRequestWithDefaults() *ConstrainRequest`

NewConstrainRequestWithDefaults instantiates a new ConstrainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccelerators

`func (o *ConstrainRequest) GetAccelerators() map[string]interface{}`

GetAccelerators returns the Accelerators field if non-nil, zero value otherwise.

### GetAcceleratorsOk

`func (o *ConstrainRequest) GetAcceleratorsOk() (*map[string]interface{}, bool)`

GetAcceleratorsOk returns a tuple with the Accelerators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerators

`func (o *ConstrainRequest) SetAccelerators(v map[string]interface{})`

SetAccelerators sets Accelerators field to given value.

### HasAccelerators

`func (o *ConstrainRequest) HasAccelerators() bool`

HasAccelerators returns a boolean if a field has been set.

### SetAcceleratorsNil

`func (o *ConstrainRequest) SetAcceleratorsNil(b bool)`

 SetAcceleratorsNil sets the value for Accelerators to be an explicit nil

### UnsetAccelerators
`func (o *ConstrainRequest) UnsetAccelerators()`

UnsetAccelerators ensures that no value is present for Accelerators, not even an explicit nil
### GetStorage

`func (o *ConstrainRequest) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ConstrainRequest) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ConstrainRequest) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ConstrainRequest) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetMemory

`func (o *ConstrainRequest) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ConstrainRequest) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ConstrainRequest) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ConstrainRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetVCPUs

`func (o *ConstrainRequest) GetVCPUs() int32`

GetVCPUs returns the VCPUs field if non-nil, zero value otherwise.

### GetVCPUsOk

`func (o *ConstrainRequest) GetVCPUsOk() (*int32, bool)`

GetVCPUsOk returns a tuple with the VCPUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCPUs

`func (o *ConstrainRequest) SetVCPUs(v int32)`

SetVCPUs sets VCPUs field to given value.

### HasVCPUs

`func (o *ConstrainRequest) HasVCPUs() bool`

HasVCPUs returns a boolean if a field has been set.

### GetRegions

`func (o *ConstrainRequest) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ConstrainRequest) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ConstrainRequest) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *ConstrainRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetClouds

`func (o *ConstrainRequest) GetClouds() []string`

GetClouds returns the Clouds field if non-nil, zero value otherwise.

### GetCloudsOk

`func (o *ConstrainRequest) GetCloudsOk() (*[]string, bool)`

GetCloudsOk returns a tuple with the Clouds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClouds

`func (o *ConstrainRequest) SetClouds(v []string)`

SetClouds sets Clouds field to given value.

### HasClouds

`func (o *ConstrainRequest) HasClouds() bool`

HasClouds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


