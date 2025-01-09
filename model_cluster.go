/*
Wisp API

Wisp API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package wisp

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Cluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cluster{}

// Cluster Cluster serializer.
type Cluster struct {
	User int32 `json:"user"`
	Name string `json:"name"`
	LaunchedAt NullableInt64 `json:"launched_at,omitempty"`
	LastUse NullableString `json:"last_use,omitempty"`
	Status *ClusterStatusEnum `json:"status,omitempty"`
	Autostop *int64 `json:"autostop,omitempty"`
	Metadata *string `json:"metadata,omitempty"`
	ToDown *bool `json:"to_down,omitempty"`
	ClusterHash NullableString `json:"cluster_hash,omitempty"`
	Handle NullablePickledHandleField `json:"handle"`
}

type _Cluster Cluster

// NewCluster instantiates a new Cluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCluster(user int32, name string, handle NullablePickledHandleField) *Cluster {
	this := Cluster{}
	this.User = user
	this.Name = name
	this.Handle = handle
	return &this
}

// NewClusterWithDefaults instantiates a new Cluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterWithDefaults() *Cluster {
	this := Cluster{}
	return &this
}

// GetUser returns the User field value
func (o *Cluster) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *Cluster) SetUser(v int32) {
	o.User = v
}

// GetName returns the Name field value
func (o *Cluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Cluster) SetName(v string) {
	o.Name = v
}

// GetLaunchedAt returns the LaunchedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetLaunchedAt() int64 {
	if o == nil || IsNil(o.LaunchedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.LaunchedAt.Get()
}

// GetLaunchedAtOk returns a tuple with the LaunchedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetLaunchedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LaunchedAt.Get(), o.LaunchedAt.IsSet()
}

// HasLaunchedAt returns a boolean if a field has been set.
func (o *Cluster) HasLaunchedAt() bool {
	if o != nil && o.LaunchedAt.IsSet() {
		return true
	}

	return false
}

// SetLaunchedAt gets a reference to the given NullableInt64 and assigns it to the LaunchedAt field.
func (o *Cluster) SetLaunchedAt(v int64) {
	o.LaunchedAt.Set(&v)
}
// SetLaunchedAtNil sets the value for LaunchedAt to be an explicit nil
func (o *Cluster) SetLaunchedAtNil() {
	o.LaunchedAt.Set(nil)
}

// UnsetLaunchedAt ensures that no value is present for LaunchedAt, not even an explicit nil
func (o *Cluster) UnsetLaunchedAt() {
	o.LaunchedAt.Unset()
}

// GetLastUse returns the LastUse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetLastUse() string {
	if o == nil || IsNil(o.LastUse.Get()) {
		var ret string
		return ret
	}
	return *o.LastUse.Get()
}

// GetLastUseOk returns a tuple with the LastUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetLastUseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUse.Get(), o.LastUse.IsSet()
}

// HasLastUse returns a boolean if a field has been set.
func (o *Cluster) HasLastUse() bool {
	if o != nil && o.LastUse.IsSet() {
		return true
	}

	return false
}

// SetLastUse gets a reference to the given NullableString and assigns it to the LastUse field.
func (o *Cluster) SetLastUse(v string) {
	o.LastUse.Set(&v)
}
// SetLastUseNil sets the value for LastUse to be an explicit nil
func (o *Cluster) SetLastUseNil() {
	o.LastUse.Set(nil)
}

// UnsetLastUse ensures that no value is present for LastUse, not even an explicit nil
func (o *Cluster) UnsetLastUse() {
	o.LastUse.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Cluster) GetStatus() ClusterStatusEnum {
	if o == nil || IsNil(o.Status) {
		var ret ClusterStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetStatusOk() (*ClusterStatusEnum, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Cluster) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ClusterStatusEnum and assigns it to the Status field.
func (o *Cluster) SetStatus(v ClusterStatusEnum) {
	o.Status = &v
}

// GetAutostop returns the Autostop field value if set, zero value otherwise.
func (o *Cluster) GetAutostop() int64 {
	if o == nil || IsNil(o.Autostop) {
		var ret int64
		return ret
	}
	return *o.Autostop
}

// GetAutostopOk returns a tuple with the Autostop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetAutostopOk() (*int64, bool) {
	if o == nil || IsNil(o.Autostop) {
		return nil, false
	}
	return o.Autostop, true
}

// HasAutostop returns a boolean if a field has been set.
func (o *Cluster) HasAutostop() bool {
	if o != nil && !IsNil(o.Autostop) {
		return true
	}

	return false
}

// SetAutostop gets a reference to the given int64 and assigns it to the Autostop field.
func (o *Cluster) SetAutostop(v int64) {
	o.Autostop = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Cluster) GetMetadata() string {
	if o == nil || IsNil(o.Metadata) {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Cluster) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *Cluster) SetMetadata(v string) {
	o.Metadata = &v
}

// GetToDown returns the ToDown field value if set, zero value otherwise.
func (o *Cluster) GetToDown() bool {
	if o == nil || IsNil(o.ToDown) {
		var ret bool
		return ret
	}
	return *o.ToDown
}

// GetToDownOk returns a tuple with the ToDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetToDownOk() (*bool, bool) {
	if o == nil || IsNil(o.ToDown) {
		return nil, false
	}
	return o.ToDown, true
}

// HasToDown returns a boolean if a field has been set.
func (o *Cluster) HasToDown() bool {
	if o != nil && !IsNil(o.ToDown) {
		return true
	}

	return false
}

// SetToDown gets a reference to the given bool and assigns it to the ToDown field.
func (o *Cluster) SetToDown(v bool) {
	o.ToDown = &v
}

// GetClusterHash returns the ClusterHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetClusterHash() string {
	if o == nil || IsNil(o.ClusterHash.Get()) {
		var ret string
		return ret
	}
	return *o.ClusterHash.Get()
}

// GetClusterHashOk returns a tuple with the ClusterHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetClusterHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterHash.Get(), o.ClusterHash.IsSet()
}

// HasClusterHash returns a boolean if a field has been set.
func (o *Cluster) HasClusterHash() bool {
	if o != nil && o.ClusterHash.IsSet() {
		return true
	}

	return false
}

// SetClusterHash gets a reference to the given NullableString and assigns it to the ClusterHash field.
func (o *Cluster) SetClusterHash(v string) {
	o.ClusterHash.Set(&v)
}
// SetClusterHashNil sets the value for ClusterHash to be an explicit nil
func (o *Cluster) SetClusterHashNil() {
	o.ClusterHash.Set(nil)
}

// UnsetClusterHash ensures that no value is present for ClusterHash, not even an explicit nil
func (o *Cluster) UnsetClusterHash() {
	o.ClusterHash.Unset()
}

// GetHandle returns the Handle field value
// If the value is explicit nil, the zero value for PickledHandleField will be returned
func (o *Cluster) GetHandle() PickledHandleField {
	if o == nil || o.Handle.Get() == nil {
		var ret PickledHandleField
		return ret
	}

	return *o.Handle.Get()
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetHandleOk() (*PickledHandleField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Handle.Get(), o.Handle.IsSet()
}

// SetHandle sets field value
func (o *Cluster) SetHandle(v PickledHandleField) {
	o.Handle.Set(&v)
}

func (o Cluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	toSerialize["name"] = o.Name
	if o.LaunchedAt.IsSet() {
		toSerialize["launched_at"] = o.LaunchedAt.Get()
	}
	if o.LastUse.IsSet() {
		toSerialize["last_use"] = o.LastUse.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Autostop) {
		toSerialize["autostop"] = o.Autostop
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.ToDown) {
		toSerialize["to_down"] = o.ToDown
	}
	if o.ClusterHash.IsSet() {
		toSerialize["cluster_hash"] = o.ClusterHash.Get()
	}
	toSerialize["handle"] = o.Handle.Get()
	return toSerialize, nil
}

func (o *Cluster) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user",
		"name",
		"handle",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCluster := _Cluster{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCluster)

	if err != nil {
		return err
	}

	*o = Cluster(varCluster)

	return err
}

type NullableCluster struct {
	value *Cluster
	isSet bool
}

func (v NullableCluster) Get() *Cluster {
	return v.value
}

func (v *NullableCluster) Set(val *Cluster) {
	v.value = val
	v.isSet = true
}

func (v NullableCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCluster(val *Cluster) *NullableCluster {
	return &NullableCluster{value: val, isSet: true}
}

func (v NullableCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


