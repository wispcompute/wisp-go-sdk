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

// checks if the Bucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Bucket{}

// Bucket Bucket configuration serializer.
type Bucket struct {
	Name NullableString `json:"name"`
	Mode NullableString `json:"mode"`
	MountLocation NullableString `json:"mount_location"`
}

type _Bucket Bucket

// NewBucket instantiates a new Bucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucket(name NullableString, mode NullableString, mountLocation NullableString) *Bucket {
	this := Bucket{}
	this.Name = name
	this.Mode = mode
	this.MountLocation = mountLocation
	return &this
}

// NewBucketWithDefaults instantiates a new Bucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketWithDefaults() *Bucket {
	this := Bucket{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Bucket) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bucket) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *Bucket) SetName(v string) {
	o.Name.Set(&v)
}

// GetMode returns the Mode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Bucket) GetMode() string {
	if o == nil || o.Mode.Get() == nil {
		var ret string
		return ret
	}

	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bucket) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// SetMode sets field value
func (o *Bucket) SetMode(v string) {
	o.Mode.Set(&v)
}

// GetMountLocation returns the MountLocation field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Bucket) GetMountLocation() string {
	if o == nil || o.MountLocation.Get() == nil {
		var ret string
		return ret
	}

	return *o.MountLocation.Get()
}

// GetMountLocationOk returns a tuple with the MountLocation field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bucket) GetMountLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountLocation.Get(), o.MountLocation.IsSet()
}

// SetMountLocation sets field value
func (o *Bucket) SetMountLocation(v string) {
	o.MountLocation.Set(&v)
}

func (o Bucket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Bucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name.Get()
	toSerialize["mode"] = o.Mode.Get()
	toSerialize["mount_location"] = o.MountLocation.Get()
	return toSerialize, nil
}

func (o *Bucket) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"mode",
		"mount_location",
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

	varBucket := _Bucket{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBucket)

	if err != nil {
		return err
	}

	*o = Bucket(varBucket)

	return err
}

type NullableBucket struct {
	value *Bucket
	isSet bool
}

func (v NullableBucket) Get() *Bucket {
	return v.value
}

func (v *NullableBucket) Set(val *Bucket) {
	v.value = val
	v.isSet = true
}

func (v NullableBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucket(val *Bucket) *NullableBucket {
	return &NullableBucket{value: val, isSet: true}
}

func (v NullableBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


