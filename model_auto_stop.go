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

// checks if the AutoStop type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoStop{}

// AutoStop struct for AutoStop
type AutoStop struct {
	Id int32 `json:"id"`
	Enabled *bool `json:"enabled,omitempty"`
	TimeoutMinutes *int32 `json:"timeout_minutes,omitempty"`
}

type _AutoStop AutoStop

// NewAutoStop instantiates a new AutoStop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoStop(id int32) *AutoStop {
	this := AutoStop{}
	this.Id = id
	return &this
}

// NewAutoStopWithDefaults instantiates a new AutoStop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoStopWithDefaults() *AutoStop {
	this := AutoStop{}
	return &this
}

// GetId returns the Id field value
func (o *AutoStop) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AutoStop) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AutoStop) SetId(v int32) {
	o.Id = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AutoStop) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoStop) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AutoStop) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AutoStop) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTimeoutMinutes returns the TimeoutMinutes field value if set, zero value otherwise.
func (o *AutoStop) GetTimeoutMinutes() int32 {
	if o == nil || IsNil(o.TimeoutMinutes) {
		var ret int32
		return ret
	}
	return *o.TimeoutMinutes
}

// GetTimeoutMinutesOk returns a tuple with the TimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoStop) GetTimeoutMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeoutMinutes) {
		return nil, false
	}
	return o.TimeoutMinutes, true
}

// HasTimeoutMinutes returns a boolean if a field has been set.
func (o *AutoStop) HasTimeoutMinutes() bool {
	if o != nil && !IsNil(o.TimeoutMinutes) {
		return true
	}

	return false
}

// SetTimeoutMinutes gets a reference to the given int32 and assigns it to the TimeoutMinutes field.
func (o *AutoStop) SetTimeoutMinutes(v int32) {
	o.TimeoutMinutes = &v
}

func (o AutoStop) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoStop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.TimeoutMinutes) {
		toSerialize["timeout_minutes"] = o.TimeoutMinutes
	}
	return toSerialize, nil
}

func (o *AutoStop) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varAutoStop := _AutoStop{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAutoStop)

	if err != nil {
		return err
	}

	*o = AutoStop(varAutoStop)

	return err
}

type NullableAutoStop struct {
	value *AutoStop
	isSet bool
}

func (v NullableAutoStop) Get() *AutoStop {
	return v.value
}

func (v *NullableAutoStop) Set(val *AutoStop) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoStop) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoStop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoStop(val *AutoStop) *NullableAutoStop {
	return &NullableAutoStop{value: val, isSet: true}
}

func (v NullableAutoStop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoStop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


