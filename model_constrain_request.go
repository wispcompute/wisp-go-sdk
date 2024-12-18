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

// ConstrainRequest Configuration serializer for cluster constraints.  Validates the full configuration schema including project details, scripts, resource requirements, and IO specifications.
type ConstrainRequest struct {
	Project Project `json:"project"`
	Setup Script `json:"setup,omitempty"`
	Run Script `json:"run,omitempty"`
	Teardown Script `json:"teardown,omitempty"`
	Resources Resources `json:"resources"`
	Io IO `json:"io"`
}

// NewConstrainRequest instantiates a new ConstrainRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstrainRequest(project Project, resources Resources, io IO) *ConstrainRequest {
	this := ConstrainRequest{}
	this.Project = project
	this.Resources = resources
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
// If the value is explicit nil, the zero value for Project will be returned
func (o *ConstrainRequest) GetProject() Project {
	if o == nil {
		var ret Project
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
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
	if o == nil  {
		var ret Script
		return ret
	}
	return o.Setup
}

// GetSetupOk returns a tuple with the Setup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConstrainRequest) GetSetupOk() (*Script, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Setup, true
}

// HasSetup returns a boolean if a field has been set.
func (o *ConstrainRequest) HasSetup() bool {
	if o != nil {
		return true
	}

	return false
}

// SetSetup gets a reference to the given Script and assigns it to the Setup field.
func (o *ConstrainRequest) SetSetup(v Script) {
	o.Setup = v
}

// GetRun returns the Run field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConstrainRequest) GetRun() Script {
	if o == nil  {
		var ret Script
		return ret
	}
	return o.Run
}

// GetRunOk returns a tuple with the Run field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConstrainRequest) GetRunOk() (*Script, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Run, true
}

// HasRun returns a boolean if a field has been set.
func (o *ConstrainRequest) HasRun() bool {
	if o != nil {
		return true
	}

	return false
}

// SetRun gets a reference to the given Script and assigns it to the Run field.
func (o *ConstrainRequest) SetRun(v Script) {
	o.Run = v
}

// GetTeardown returns the Teardown field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConstrainRequest) GetTeardown() Script {
	if o == nil  {
		var ret Script
		return ret
	}
	return o.Teardown
}

// GetTeardownOk returns a tuple with the Teardown field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConstrainRequest) GetTeardownOk() (*Script, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Teardown, true
}

// HasTeardown returns a boolean if a field has been set.
func (o *ConstrainRequest) HasTeardown() bool {
	if o != nil {
		return true
	}

	return false
}

// SetTeardown gets a reference to the given Script and assigns it to the Teardown field.
func (o *ConstrainRequest) SetTeardown(v Script) {
	o.Teardown = v
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
	if o == nil  {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *ConstrainRequest) SetResources(v Resources) {
	o.Resources = v
}

// GetIo returns the Io field value
// If the value is explicit nil, the zero value for IO will be returned
func (o *ConstrainRequest) GetIo() IO {
	if o == nil {
		var ret IO
		return ret
	}

	return o.Io
}

// GetIoOk returns a tuple with the Io field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConstrainRequest) GetIoOk() (*IO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Io, true
}

// SetIo sets field value
func (o *ConstrainRequest) SetIo(v IO) {
	o.Io = v
}

func (o ConstrainRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	toSerialize["project"] = o.Project
	toSerialize["setup"] = o.Setup
	toSerialize["run"] = o.Run
	toSerialize["teardown"] = o.Teardown
	toSerialize["resources"] = o.Resources
	toSerialize["io"] = o.Io
	
	return json.Marshal(toSerialize)
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


