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

// checks if the JobConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobConfig{}

// JobConfig Config for the job. Not to be confused with the job itself.
type JobConfig struct {
	Id int32 `json:"id"`
	Autostop AutoStop `json:"autostop"`
	Notifications NotificationConfig `json:"notifications"`
}

type _JobConfig JobConfig

// NewJobConfig instantiates a new JobConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobConfig(id int32, autostop AutoStop, notifications NotificationConfig) *JobConfig {
	this := JobConfig{}
	this.Id = id
	this.Autostop = autostop
	this.Notifications = notifications
	return &this
}

// NewJobConfigWithDefaults instantiates a new JobConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobConfigWithDefaults() *JobConfig {
	this := JobConfig{}
	return &this
}

// GetId returns the Id field value
func (o *JobConfig) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JobConfig) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JobConfig) SetId(v int32) {
	o.Id = v
}

// GetAutostop returns the Autostop field value
func (o *JobConfig) GetAutostop() AutoStop {
	if o == nil {
		var ret AutoStop
		return ret
	}

	return o.Autostop
}

// GetAutostopOk returns a tuple with the Autostop field value
// and a boolean to check if the value has been set.
func (o *JobConfig) GetAutostopOk() (*AutoStop, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Autostop, true
}

// SetAutostop sets field value
func (o *JobConfig) SetAutostop(v AutoStop) {
	o.Autostop = v
}

// GetNotifications returns the Notifications field value
func (o *JobConfig) GetNotifications() NotificationConfig {
	if o == nil {
		var ret NotificationConfig
		return ret
	}

	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value
// and a boolean to check if the value has been set.
func (o *JobConfig) GetNotificationsOk() (*NotificationConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notifications, true
}

// SetNotifications sets field value
func (o *JobConfig) SetNotifications(v NotificationConfig) {
	o.Notifications = v
}

func (o JobConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["autostop"] = o.Autostop
	toSerialize["notifications"] = o.Notifications
	return toSerialize, nil
}

func (o *JobConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"autostop",
		"notifications",
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

	varJobConfig := _JobConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJobConfig)

	if err != nil {
		return err
	}

	*o = JobConfig(varJobConfig)

	return err
}

type NullableJobConfig struct {
	value *JobConfig
	isSet bool
}

func (v NullableJobConfig) Get() *JobConfig {
	return v.value
}

func (v *NullableJobConfig) Set(val *JobConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableJobConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableJobConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobConfig(val *JobConfig) *NullableJobConfig {
	return &NullableJobConfig{value: val, isSet: true}
}

func (v NullableJobConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


