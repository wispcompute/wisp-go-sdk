/*
Wisp API

Wisp API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ConstrainRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstrainRequest{}

// ConstrainRequest Serializer for cluster constraint request.  Fields:     Accelerators (dict): Accelerator specifications including name, count and VRAM     Storage (int): Required storage in GB     Memory (int): Required memory in GB     VCPUs (int): Required number of virtual CPUs     Regions (list): List of acceptable regions     Clouds (list): List of acceptable cloud providers
type ConstrainRequest struct {
	Accelerators map[string]interface{} `json:"Accelerators,omitempty"`
	Storage *int32 `json:"Storage,omitempty"`
	Memory *int32 `json:"Memory,omitempty"`
	VCPUs *int32 `json:"VCPUs,omitempty"`
	Regions []string `json:"Regions,omitempty"`
	Clouds []string `json:"Clouds,omitempty"`
}

// NewConstrainRequest instantiates a new ConstrainRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstrainRequest() *ConstrainRequest {
	this := ConstrainRequest{}
	return &this
}

// NewConstrainRequestWithDefaults instantiates a new ConstrainRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstrainRequestWithDefaults() *ConstrainRequest {
	this := ConstrainRequest{}
	return &this
}

// GetAccelerators returns the Accelerators field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConstrainRequest) GetAccelerators() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Accelerators
}

// GetAcceleratorsOk returns a tuple with the Accelerators field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConstrainRequest) GetAcceleratorsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Accelerators) {
		return map[string]interface{}{}, false
	}
	return o.Accelerators, true
}

// HasAccelerators returns a boolean if a field has been set.
func (o *ConstrainRequest) HasAccelerators() bool {
	if o != nil && !IsNil(o.Accelerators) {
		return true
	}

	return false
}

// SetAccelerators gets a reference to the given map[string]interface{} and assigns it to the Accelerators field.
func (o *ConstrainRequest) SetAccelerators(v map[string]interface{}) {
	o.Accelerators = v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ConstrainRequest) GetStorage() int32 {
	if o == nil || IsNil(o.Storage) {
		var ret int32
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstrainRequest) GetStorageOk() (*int32, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ConstrainRequest) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given int32 and assigns it to the Storage field.
func (o *ConstrainRequest) SetStorage(v int32) {
	o.Storage = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ConstrainRequest) GetMemory() int32 {
	if o == nil || IsNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstrainRequest) GetMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ConstrainRequest) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *ConstrainRequest) SetMemory(v int32) {
	o.Memory = &v
}

// GetVCPUs returns the VCPUs field value if set, zero value otherwise.
func (o *ConstrainRequest) GetVCPUs() int32 {
	if o == nil || IsNil(o.VCPUs) {
		var ret int32
		return ret
	}
	return *o.VCPUs
}

// GetVCPUsOk returns a tuple with the VCPUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstrainRequest) GetVCPUsOk() (*int32, bool) {
	if o == nil || IsNil(o.VCPUs) {
		return nil, false
	}
	return o.VCPUs, true
}

// HasVCPUs returns a boolean if a field has been set.
func (o *ConstrainRequest) HasVCPUs() bool {
	if o != nil && !IsNil(o.VCPUs) {
		return true
	}

	return false
}

// SetVCPUs gets a reference to the given int32 and assigns it to the VCPUs field.
func (o *ConstrainRequest) SetVCPUs(v int32) {
	o.VCPUs = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *ConstrainRequest) GetRegions() []string {
	if o == nil || IsNil(o.Regions) {
		var ret []string
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstrainRequest) GetRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *ConstrainRequest) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *ConstrainRequest) SetRegions(v []string) {
	o.Regions = v
}

// GetClouds returns the Clouds field value if set, zero value otherwise.
func (o *ConstrainRequest) GetClouds() []string {
	if o == nil || IsNil(o.Clouds) {
		var ret []string
		return ret
	}
	return o.Clouds
}

// GetCloudsOk returns a tuple with the Clouds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstrainRequest) GetCloudsOk() ([]string, bool) {
	if o == nil || IsNil(o.Clouds) {
		return nil, false
	}
	return o.Clouds, true
}

// HasClouds returns a boolean if a field has been set.
func (o *ConstrainRequest) HasClouds() bool {
	if o != nil && !IsNil(o.Clouds) {
		return true
	}

	return false
}

// SetClouds gets a reference to the given []string and assigns it to the Clouds field.
func (o *ConstrainRequest) SetClouds(v []string) {
	o.Clouds = v
}

func (o ConstrainRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstrainRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Accelerators != nil {
		toSerialize["Accelerators"] = o.Accelerators
	}
	if !IsNil(o.Storage) {
		toSerialize["Storage"] = o.Storage
	}
	if !IsNil(o.Memory) {
		toSerialize["Memory"] = o.Memory
	}
	if !IsNil(o.VCPUs) {
		toSerialize["VCPUs"] = o.VCPUs
	}
	if !IsNil(o.Regions) {
		toSerialize["Regions"] = o.Regions
	}
	if !IsNil(o.Clouds) {
		toSerialize["Clouds"] = o.Clouds
	}
	return toSerialize, nil
}

type NullableConstrainRequest struct {
	value *ConstrainRequest
	isSet bool
}

func (v NullableConstrainRequest) Get() *ConstrainRequest {
	return v.value
}

func (v *NullableConstrainRequest) Set(val *ConstrainRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConstrainRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConstrainRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstrainRequest(val *ConstrainRequest) *NullableConstrainRequest {
	return &NullableConstrainRequest{value: val, isSet: true}
}

func (v NullableConstrainRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstrainRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


