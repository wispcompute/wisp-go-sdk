/*
Wisp API

Wisp API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package wisp

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ProvisionLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisionLog{}

// ProvisionLog struct for ProvisionLog
type ProvisionLog struct {
	Cluster Cluster `json:"cluster"`
	ProvisionId string `json:"provision_id"`
	StepName string `json:"step_name"`
	Status *ProvisionLogStatusEnum `json:"status,omitempty"`
	Message NullableString `json:"message,omitempty"`
	StartedAt time.Time `json:"started_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type _ProvisionLog ProvisionLog

// NewProvisionLog instantiates a new ProvisionLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionLog(cluster Cluster, provisionId string, stepName string, startedAt time.Time, updatedAt time.Time) *ProvisionLog {
	this := ProvisionLog{}
	this.Cluster = cluster
	this.ProvisionId = provisionId
	this.StepName = stepName
	this.StartedAt = startedAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewProvisionLogWithDefaults instantiates a new ProvisionLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionLogWithDefaults() *ProvisionLog {
	this := ProvisionLog{}
	return &this
}

// GetCluster returns the Cluster field value
func (o *ProvisionLog) GetCluster() Cluster {
	if o == nil {
		var ret Cluster
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *ProvisionLog) GetClusterOk() (*Cluster, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *ProvisionLog) SetCluster(v Cluster) {
	o.Cluster = v
}

// GetProvisionId returns the ProvisionId field value
func (o *ProvisionLog) GetProvisionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProvisionId
}

// GetProvisionIdOk returns a tuple with the ProvisionId field value
// and a boolean to check if the value has been set.
func (o *ProvisionLog) GetProvisionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisionId, true
}

// SetProvisionId sets field value
func (o *ProvisionLog) SetProvisionId(v string) {
	o.ProvisionId = v
}

// GetStepName returns the StepName field value
func (o *ProvisionLog) GetStepName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StepName
}

// GetStepNameOk returns a tuple with the StepName field value
// and a boolean to check if the value has been set.
func (o *ProvisionLog) GetStepNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StepName, true
}

// SetStepName sets field value
func (o *ProvisionLog) SetStepName(v string) {
	o.StepName = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProvisionLog) GetStatus() ProvisionLogStatusEnum {
	if o == nil || IsNil(o.Status) {
		var ret ProvisionLogStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionLog) GetStatusOk() (*ProvisionLogStatusEnum, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProvisionLog) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ProvisionLogStatusEnum and assigns it to the Status field.
func (o *ProvisionLog) SetStatus(v ProvisionLogStatusEnum) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisionLog) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisionLog) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *ProvisionLog) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *ProvisionLog) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *ProvisionLog) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *ProvisionLog) UnsetMessage() {
	o.Message.Unset()
}

// GetStartedAt returns the StartedAt field value
func (o *ProvisionLog) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *ProvisionLog) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *ProvisionLog) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ProvisionLog) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ProvisionLog) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ProvisionLog) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o ProvisionLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisionLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster"] = o.Cluster
	toSerialize["provision_id"] = o.ProvisionId
	toSerialize["step_name"] = o.StepName
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	toSerialize["started_at"] = o.StartedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *ProvisionLog) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cluster",
		"provision_id",
		"step_name",
		"started_at",
		"updated_at",
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

	varProvisionLog := _ProvisionLog{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProvisionLog)

	if err != nil {
		return err
	}

	*o = ProvisionLog(varProvisionLog)

	return err
}

type NullableProvisionLog struct {
	value *ProvisionLog
	isSet bool
}

func (v NullableProvisionLog) Get() *ProvisionLog {
	return v.value
}

func (v *NullableProvisionLog) Set(val *ProvisionLog) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionLog) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionLog(val *ProvisionLog) *NullableProvisionLog {
	return &NullableProvisionLog{value: val, isSet: true}
}

func (v NullableProvisionLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


