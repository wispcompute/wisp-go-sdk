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

// checks if the ConstrainRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstrainRequest{}

// ConstrainRequest Configuration serializer for cluster constraints.  Validates the full configuration schema including project details, scripts, resource requirements, and IO specifications.
type ConstrainRequest struct {
	Project Project `json:"project"`
	Setup NullableScript `json:"setup,omitempty"`
	Run NullableScript `json:"run,omitempty"`
	Teardown NullableScript `json:"teardown,omitempty"`
	Resources Resources `json:"resources"`
	JobConfig NullableJobConfig `json:"job_config"`
	Io NullableIO `json:"io"`
	Flags map[string]interface{} `json:"flags,omitempty"`
}

type _ConstrainRequest ConstrainRequest

// NewConstrainRequest instantiates a new ConstrainRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstrainRequest(project Project, resources Resources, jobConfig NullableJobConfig, io NullableIO) *ConstrainRequest {
	this := ConstrainRequest{}
	this.Project = project
	this.Resources = resources
	this.JobConfig = jobConfig
	this.Io = io
	return &this
}

// NewConstrainRequestWithDefaults instantiates a new ConstrainRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstrainRequestWithDefaults() *ConstrainRequest {
	this := ConstrainRequest{}
	return &this
}

// GetProject returns the Project field value
func (o *ConstrainRequest) GetProject() Project {
	if o == nil {
		var ret Project
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *ConstrainRequest) GetProjectOk() (*Project, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *ConstrainRequest) SetProject(v Project) {
	o.Project = v
}

// GetSetup returns the Setup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConstrainRequest) GetSetup() Script {
	if o == nil || IsNil(o.Setup.Get()) {
		var ret Script
		return ret
	}
	return *o.Setup.Get()
}

// GetSetupOk returns a tuple with the Setup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConstrainRequest) GetSetupOk() (*Script, bool) {
	if o == nil {
		return nil, false
	}
	return o.Setup.Get(), o.Setup.IsSet()
}

// HasSetup returns a boolean if a field has been set.
func (o *ConstrainRequest) HasSetup() bool {
	if o != nil && o.Setup.IsSet() {
		return true
	}

	return false
}

// SetSetup gets a reference to the given NullableScript and assigns it to the Setup field.
func (o *ConstrainRequest) SetSetup(v Script) {
	o.Setup.Set(&v)
}
// SetSetupNil sets the value for Setup to be an explicit nil
func (o *ConstrainRequest) SetSetupNil() {
	o.Setup.Set(nil)
}

// UnsetSetup ensures that no value is present for Setup, not even an explicit nil
func (o *ConstrainRequest) UnsetSetup() {
	o.Setup.Unset()
}

// GetRun returns the Run field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConstrainRequest) GetRun() Script {
	if o == nil || IsNil(o.Run.Get()) {
		var ret Script
		return ret
	}
	return *o.Run.Get()
}

// GetRunOk returns a tuple with the Run field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConstrainRequest) GetRunOk() (*Script, bool) {
	if o == nil {
		return nil, false
	}
	return o.Run.Get(), o.Run.IsSet()
}

// HasRun returns a boolean if a field has been set.
func (o *ConstrainRequest) HasRun() bool {
	if o != nil && o.Run.IsSet() {
		return true
	}

	return false
}

// SetRun gets a reference to the given NullableScript and assigns it to the Run field.
func (o *ConstrainRequest) SetRun(v Script) {
	o.Run.Set(&v)
}
// SetRunNil sets the value for Run to be an explicit nil
func (o *ConstrainRequest) SetRunNil() {
	o.Run.Set(nil)
}

// UnsetRun ensures that no value is present for Run, not even an explicit nil
func (o *ConstrainRequest) UnsetRun() {
	o.Run.Unset()
}

// GetTeardown returns the Teardown field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConstrainRequest) GetTeardown() Script {
	if o == nil || IsNil(o.Teardown.Get()) {
		var ret Script
		return ret
	}
	return *o.Teardown.Get()
}

// GetTeardownOk returns a tuple with the Teardown field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConstrainRequest) GetTeardownOk() (*Script, bool) {
	if o == nil {
		return nil, false
	}
	return o.Teardown.Get(), o.Teardown.IsSet()
}

// HasTeardown returns a boolean if a field has been set.
func (o *ConstrainRequest) HasTeardown() bool {
	if o != nil && o.Teardown.IsSet() {
		return true
	}

	return false
}

// SetTeardown gets a reference to the given NullableScript and assigns it to the Teardown field.
func (o *ConstrainRequest) SetTeardown(v Script) {
	o.Teardown.Set(&v)
}
// SetTeardownNil sets the value for Teardown to be an explicit nil
func (o *ConstrainRequest) SetTeardownNil() {
	o.Teardown.Set(nil)
}

// UnsetTeardown ensures that no value is present for Teardown, not even an explicit nil
func (o *ConstrainRequest) UnsetTeardown() {
	o.Teardown.Unset()
}

// GetResources returns the Resources field value
func (o *ConstrainRequest) GetResources() Resources {
	if o == nil {
		var ret Resources
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *ConstrainRequest) GetResourcesOk() (*Resources, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *ConstrainRequest) SetResources(v Resources) {
	o.Resources = v
}

// GetJobConfig returns the JobConfig field value
// If the value is explicit nil, the zero value for JobConfig will be returned
func (o *ConstrainRequest) GetJobConfig() JobConfig {
	if o == nil || o.JobConfig.Get() == nil {
		var ret JobConfig
		return ret
	}

	return *o.JobConfig.Get()
}

// GetJobConfigOk returns a tuple with the JobConfig field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConstrainRequest) GetJobConfigOk() (*JobConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobConfig.Get(), o.JobConfig.IsSet()
}

// SetJobConfig sets field value
func (o *ConstrainRequest) SetJobConfig(v JobConfig) {
	o.JobConfig.Set(&v)
}

// GetIo returns the Io field value
// If the value is explicit nil, the zero value for IO will be returned
func (o *ConstrainRequest) GetIo() IO {
	if o == nil || o.Io.Get() == nil {
		var ret IO
		return ret
	}

	return *o.Io.Get()
}

// GetIoOk returns a tuple with the Io field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConstrainRequest) GetIoOk() (*IO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Io.Get(), o.Io.IsSet()
}

// SetIo sets field value
func (o *ConstrainRequest) SetIo(v IO) {
	o.Io.Set(&v)
}

// GetFlags returns the Flags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConstrainRequest) GetFlags() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConstrainRequest) GetFlagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Flags) {
		return map[string]interface{}{}, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *ConstrainRequest) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given map[string]interface{} and assigns it to the Flags field.
func (o *ConstrainRequest) SetFlags(v map[string]interface{}) {
	o.Flags = v
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
	toSerialize["project"] = o.Project
	if o.Setup.IsSet() {
		toSerialize["setup"] = o.Setup.Get()
	}
	if o.Run.IsSet() {
		toSerialize["run"] = o.Run.Get()
	}
	if o.Teardown.IsSet() {
		toSerialize["teardown"] = o.Teardown.Get()
	}
	toSerialize["resources"] = o.Resources
	toSerialize["job_config"] = o.JobConfig.Get()
	toSerialize["io"] = o.Io.Get()
	if o.Flags != nil {
		toSerialize["flags"] = o.Flags
	}
	return toSerialize, nil
}

func (o *ConstrainRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"project",
		"resources",
		"job_config",
		"io",
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

	varConstrainRequest := _ConstrainRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConstrainRequest)

	if err != nil {
		return err
	}

	*o = ConstrainRequest(varConstrainRequest)

	return err
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


