/*
Wisp API

Wisp API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package wisp

import (
	"encoding/json"
)

// Resources Resource configuration serializer.
type Resources struct {
	Clouds *[]*string `json:"clouds,omitempty"`
	Regions *[]*string `json:"regions,omitempty"`
	Areas *[]*string `json:"areas,omitempty"`
	Memory NullableInt32 `json:"memory,omitempty"`
	Cpus NullableInt32 `json:"cpus,omitempty"`
	Storage NullableInt32 `json:"storage,omitempty"`
	PersistentDisk NullableInt32 `json:"persistent_disk,omitempty"`
	Accelerators []*string `json:"accelerators"`
	ComputeCapability NullableString `json:"compute_capability,omitempty"`
	Vram NullableInt32 `json:"vram,omitempty"`
	AcceleratorCount NullableInt32 `json:"accelerator_count,omitempty"`
	Platform NullableString `json:"platform,omitempty"`
}

// NewResources instantiates a new Resources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResources(accelerators []*string) *Resources {
	this := Resources{}
	this.Accelerators = accelerators
	return &this
}

// NewResourcesWithDefaults instantiates a new Resources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcesWithDefaults() *Resources {
	this := Resources{}
	return &this
}

// GetClouds returns the Clouds field value if set, zero value otherwise.
func (o *Resources) GetClouds() []*string {
	if o == nil || o.Clouds == nil {
		var ret []*string
		return ret
	}
	return *o.Clouds
}

// GetCloudsOk returns a tuple with the Clouds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resources) GetCloudsOk() (*[]*string, bool) {
	if o == nil || o.Clouds == nil {
		return nil, false
	}
	return o.Clouds, true
}

// HasClouds returns a boolean if a field has been set.
func (o *Resources) HasClouds() bool {
	if o != nil && o.Clouds != nil {
		return true
	}

	return false
}

// SetClouds gets a reference to the given []*string and assigns it to the Clouds field.
func (o *Resources) SetClouds(v []*string) {
	o.Clouds = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *Resources) GetRegions() []*string {
	if o == nil || o.Regions == nil {
		var ret []*string
		return ret
	}
	return *o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resources) GetRegionsOk() (*[]*string, bool) {
	if o == nil || o.Regions == nil {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *Resources) HasRegions() bool {
	if o != nil && o.Regions != nil {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []*string and assigns it to the Regions field.
func (o *Resources) SetRegions(v []*string) {
	o.Regions = &v
}

// GetAreas returns the Areas field value if set, zero value otherwise.
func (o *Resources) GetAreas() []*string {
	if o == nil || o.Areas == nil {
		var ret []*string
		return ret
	}
	return *o.Areas
}

// GetAreasOk returns a tuple with the Areas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resources) GetAreasOk() (*[]*string, bool) {
	if o == nil || o.Areas == nil {
		return nil, false
	}
	return o.Areas, true
}

// HasAreas returns a boolean if a field has been set.
func (o *Resources) HasAreas() bool {
	if o != nil && o.Areas != nil {
		return true
	}

	return false
}

// SetAreas gets a reference to the given []*string and assigns it to the Areas field.
func (o *Resources) SetAreas(v []*string) {
	o.Areas = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resources) GetMemory() int32 {
	if o == nil || o.Memory.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resources) GetMemoryOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *Resources) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableInt32 and assigns it to the Memory field.
func (o *Resources) SetMemory(v int32) {
	o.Memory.Set(&v)
}
// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *Resources) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *Resources) UnsetMemory() {
	o.Memory.Unset()
}

// GetCpus returns the Cpus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resources) GetCpus() int32 {
	if o == nil || o.Cpus.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Cpus.Get()
}

// GetCpusOk returns a tuple with the Cpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resources) GetCpusOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Cpus.Get(), o.Cpus.IsSet()
}

// HasCpus returns a boolean if a field has been set.
func (o *Resources) HasCpus() bool {
	if o != nil && o.Cpus.IsSet() {
		return true
	}

	return false
}

// SetCpus gets a reference to the given NullableInt32 and assigns it to the Cpus field.
func (o *Resources) SetCpus(v int32) {
	o.Cpus.Set(&v)
}
// SetCpusNil sets the value for Cpus to be an explicit nil
func (o *Resources) SetCpusNil() {
	o.Cpus.Set(nil)
}

// UnsetCpus ensures that no value is present for Cpus, not even an explicit nil
func (o *Resources) UnsetCpus() {
	o.Cpus.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resources) GetStorage() int32 {
	if o == nil || o.Storage.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resources) GetStorageOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *Resources) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableInt32 and assigns it to the Storage field.
func (o *Resources) SetStorage(v int32) {
	o.Storage.Set(&v)
}
// SetStorageNil sets the value for Storage to be an explicit nil
func (o *Resources) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *Resources) UnsetStorage() {
	o.Storage.Unset()
}

// GetPersistentDisk returns the PersistentDisk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resources) GetPersistentDisk() int32 {
	if o == nil || o.PersistentDisk.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PersistentDisk.Get()
}

// GetPersistentDiskOk returns a tuple with the PersistentDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resources) GetPersistentDiskOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PersistentDisk.Get(), o.PersistentDisk.IsSet()
}

// HasPersistentDisk returns a boolean if a field has been set.
func (o *Resources) HasPersistentDisk() bool {
	if o != nil && o.PersistentDisk.IsSet() {
		return true
	}

	return false
}

// SetPersistentDisk gets a reference to the given NullableInt32 and assigns it to the PersistentDisk field.
func (o *Resources) SetPersistentDisk(v int32) {
	o.PersistentDisk.Set(&v)
}
// SetPersistentDiskNil sets the value for PersistentDisk to be an explicit nil
func (o *Resources) SetPersistentDiskNil() {
	o.PersistentDisk.Set(nil)
}

// UnsetPersistentDisk ensures that no value is present for PersistentDisk, not even an explicit nil
func (o *Resources) UnsetPersistentDisk() {
	o.PersistentDisk.Unset()
}

// GetAccelerators returns the Accelerators field value
// If the value is explicit nil, the zero value for []*string will be returned
func (o *Resources) GetAccelerators() []*string {
	if o == nil {
		var ret []*string
		return ret
	}

	return o.Accelerators
}

// GetAcceleratorsOk returns a tuple with the Accelerators field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resources) GetAcceleratorsOk() (*[]*string, bool) {
	if o == nil || o.Accelerators == nil {
		return nil, false
	}
	return &o.Accelerators, true
}

// SetAccelerators sets field value
func (o *Resources) SetAccelerators(v []*string) {
	o.Accelerators = v
}

// GetComputeCapability returns the ComputeCapability field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resources) GetComputeCapability() string {
	if o == nil || o.ComputeCapability.Get() == nil {
		var ret string
		return ret
	}
	return *o.ComputeCapability.Get()
}

// GetComputeCapabilityOk returns a tuple with the ComputeCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resources) GetComputeCapabilityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ComputeCapability.Get(), o.ComputeCapability.IsSet()
}

// HasComputeCapability returns a boolean if a field has been set.
func (o *Resources) HasComputeCapability() bool {
	if o != nil && o.ComputeCapability.IsSet() {
		return true
	}

	return false
}

// SetComputeCapability gets a reference to the given NullableString and assigns it to the ComputeCapability field.
func (o *Resources) SetComputeCapability(v string) {
	o.ComputeCapability.Set(&v)
}
// SetComputeCapabilityNil sets the value for ComputeCapability to be an explicit nil
func (o *Resources) SetComputeCapabilityNil() {
	o.ComputeCapability.Set(nil)
}

// UnsetComputeCapability ensures that no value is present for ComputeCapability, not even an explicit nil
func (o *Resources) UnsetComputeCapability() {
	o.ComputeCapability.Unset()
}

// GetVram returns the Vram field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resources) GetVram() int32 {
	if o == nil || o.Vram.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Vram.Get()
}

// GetVramOk returns a tuple with the Vram field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resources) GetVramOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Vram.Get(), o.Vram.IsSet()
}

// HasVram returns a boolean if a field has been set.
func (o *Resources) HasVram() bool {
	if o != nil && o.Vram.IsSet() {
		return true
	}

	return false
}

// SetVram gets a reference to the given NullableInt32 and assigns it to the Vram field.
func (o *Resources) SetVram(v int32) {
	o.Vram.Set(&v)
}
// SetVramNil sets the value for Vram to be an explicit nil
func (o *Resources) SetVramNil() {
	o.Vram.Set(nil)
}

// UnsetVram ensures that no value is present for Vram, not even an explicit nil
func (o *Resources) UnsetVram() {
	o.Vram.Unset()
}

// GetAcceleratorCount returns the AcceleratorCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resources) GetAcceleratorCount() int32 {
	if o == nil || o.AcceleratorCount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.AcceleratorCount.Get()
}

// GetAcceleratorCountOk returns a tuple with the AcceleratorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resources) GetAcceleratorCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AcceleratorCount.Get(), o.AcceleratorCount.IsSet()
}

// HasAcceleratorCount returns a boolean if a field has been set.
func (o *Resources) HasAcceleratorCount() bool {
	if o != nil && o.AcceleratorCount.IsSet() {
		return true
	}

	return false
}

// SetAcceleratorCount gets a reference to the given NullableInt32 and assigns it to the AcceleratorCount field.
func (o *Resources) SetAcceleratorCount(v int32) {
	o.AcceleratorCount.Set(&v)
}
// SetAcceleratorCountNil sets the value for AcceleratorCount to be an explicit nil
func (o *Resources) SetAcceleratorCountNil() {
	o.AcceleratorCount.Set(nil)
}

// UnsetAcceleratorCount ensures that no value is present for AcceleratorCount, not even an explicit nil
func (o *Resources) UnsetAcceleratorCount() {
	o.AcceleratorCount.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resources) GetPlatform() string {
	if o == nil || o.Platform.Get() == nil {
		var ret string
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resources) GetPlatformOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *Resources) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableString and assigns it to the Platform field.
func (o *Resources) SetPlatform(v string) {
	o.Platform.Set(&v)
}
// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *Resources) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *Resources) UnsetPlatform() {
	o.Platform.Unset()
}

func (o Resources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Clouds != nil {
		toSerialize["clouds"] = o.Clouds
	}
	if o.Regions != nil {
		toSerialize["regions"] = o.Regions
	}
	if o.Areas != nil {
		toSerialize["areas"] = o.Areas
	}
	if o.Memory.IsSet() {
		toSerialize["memory"] = o.Memory.Get()
	}
	if o.Cpus.IsSet() {
		toSerialize["cpus"] = o.Cpus.Get()
	}
	if o.Storage.IsSet() {
		toSerialize["storage"] = o.Storage.Get()
	}
	if o.PersistentDisk.IsSet() {
		toSerialize["persistent_disk"] = o.PersistentDisk.Get()
	}
	if o.Accelerators != nil {
		toSerialize["accelerators"] = o.Accelerators
	}
	if o.ComputeCapability.IsSet() {
		toSerialize["compute_capability"] = o.ComputeCapability.Get()
	}
	if o.Vram.IsSet() {
		toSerialize["vram"] = o.Vram.Get()
	}
	if o.AcceleratorCount.IsSet() {
		toSerialize["accelerator_count"] = o.AcceleratorCount.Get()
	}
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableResources struct {
	value *Resources
	isSet bool
}

func (v NullableResources) Get() *Resources {
	return v.value
}

func (v *NullableResources) Set(val *Resources) {
	v.value = val
	v.isSet = true
}

func (v NullableResources) IsSet() bool {
	return v.isSet
}

func (v *NullableResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResources(val *Resources) *NullableResources {
	return &NullableResources{value: val, isSet: true}
}

func (v NullableResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


