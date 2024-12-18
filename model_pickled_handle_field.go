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

// PickledHandleField Custom field serializer for pickled handle data.
type PickledHandleField struct {
	StableInternalIp NullableString `json:"stable_internal_ip"`
	StableExternalIp NullableString `json:"stable_external_ip"`
	StableSshPorts []int32 `json:"stable_ssh_ports"`
	SshUser NullableString `json:"ssh_user"`
}

// NewPickledHandleField instantiates a new PickledHandleField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPickledHandleField(stableInternalIp NullableString, stableExternalIp NullableString, stableSshPorts []int32, sshUser NullableString) *PickledHandleField {
	this := PickledHandleField{}
	this.StableInternalIp = stableInternalIp
	this.StableExternalIp = stableExternalIp
	this.StableSshPorts = stableSshPorts
	this.SshUser = sshUser
	return &this
}

// NewPickledHandleFieldWithDefaults instantiates a new PickledHandleField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPickledHandleFieldWithDefaults() *PickledHandleField {
	this := PickledHandleField{}
	return &this
}

// GetStableInternalIp returns the StableInternalIp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PickledHandleField) GetStableInternalIp() string {
	if o == nil || o.StableInternalIp.Get() == nil {
		var ret string
		return ret
	}

	return *o.StableInternalIp.Get()
}

// GetStableInternalIpOk returns a tuple with the StableInternalIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PickledHandleField) GetStableInternalIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StableInternalIp.Get(), o.StableInternalIp.IsSet()
}

// SetStableInternalIp sets field value
func (o *PickledHandleField) SetStableInternalIp(v string) {
	o.StableInternalIp.Set(&v)
}

// GetStableExternalIp returns the StableExternalIp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PickledHandleField) GetStableExternalIp() string {
	if o == nil || o.StableExternalIp.Get() == nil {
		var ret string
		return ret
	}

	return *o.StableExternalIp.Get()
}

// GetStableExternalIpOk returns a tuple with the StableExternalIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PickledHandleField) GetStableExternalIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StableExternalIp.Get(), o.StableExternalIp.IsSet()
}

// SetStableExternalIp sets field value
func (o *PickledHandleField) SetStableExternalIp(v string) {
	o.StableExternalIp.Set(&v)
}

// GetStableSshPorts returns the StableSshPorts field value
func (o *PickledHandleField) GetStableSshPorts() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.StableSshPorts
}

// GetStableSshPortsOk returns a tuple with the StableSshPorts field value
// and a boolean to check if the value has been set.
func (o *PickledHandleField) GetStableSshPortsOk() (*[]int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StableSshPorts, true
}

// SetStableSshPorts sets field value
func (o *PickledHandleField) SetStableSshPorts(v []int32) {
	o.StableSshPorts = v
}

// GetSshUser returns the SshUser field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PickledHandleField) GetSshUser() string {
	if o == nil || o.SshUser.Get() == nil {
		var ret string
		return ret
	}

	return *o.SshUser.Get()
}

// GetSshUserOk returns a tuple with the SshUser field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PickledHandleField) GetSshUserOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SshUser.Get(), o.SshUser.IsSet()
}

// SetSshUser sets field value
func (o *PickledHandleField) SetSshUser(v string) {
	o.SshUser.Set(&v)
}

func (o PickledHandleField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["stable_internal_ip"] = o.StableInternalIp.Get()
	}
	if true {
		toSerialize["stable_external_ip"] = o.StableExternalIp.Get()
	}
	if true {
		toSerialize["stable_ssh_ports"] = o.StableSshPorts
	}
	if true {
		toSerialize["ssh_user"] = o.SshUser.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePickledHandleField struct {
	value *PickledHandleField
	isSet bool
}

func (v NullablePickledHandleField) Get() *PickledHandleField {
	return v.value
}

func (v *NullablePickledHandleField) Set(val *PickledHandleField) {
	v.value = val
	v.isSet = true
}

func (v NullablePickledHandleField) IsSet() bool {
	return v.isSet
}

func (v *NullablePickledHandleField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePickledHandleField(val *PickledHandleField) *NullablePickledHandleField {
	return &NullablePickledHandleField{value: val, isSet: true}
}

func (v NullablePickledHandleField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePickledHandleField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


