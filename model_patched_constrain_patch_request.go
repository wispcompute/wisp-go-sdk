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

// PatchedConstrainPatchRequest Serializer for cluster constraint patch request.
type PatchedConstrainPatchRequest struct {
	JobId *string `json:"job_id,omitempty"`
	Config interface{} `json:"config,omitempty"`
}

// NewPatchedConstrainPatchRequest instantiates a new PatchedConstrainPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedConstrainPatchRequest() *PatchedConstrainPatchRequest {
	this := PatchedConstrainPatchRequest{}
	return &this
}

// NewPatchedConstrainPatchRequestWithDefaults instantiates a new PatchedConstrainPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedConstrainPatchRequestWithDefaults() *PatchedConstrainPatchRequest {
	this := PatchedConstrainPatchRequest{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *PatchedConstrainPatchRequest) GetJobId() string {
	if o == nil || o.JobId == nil {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedConstrainPatchRequest) GetJobIdOk() (*string, bool) {
	if o == nil || o.JobId == nil {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *PatchedConstrainPatchRequest) HasJobId() bool {
	if o != nil && o.JobId != nil {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *PatchedConstrainPatchRequest) SetJobId(v string) {
	o.JobId = &v
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedConstrainPatchRequest) GetConfig() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedConstrainPatchRequest) GetConfigOk() (*interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return &o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *PatchedConstrainPatchRequest) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given interface{} and assigns it to the Config field.
func (o *PatchedConstrainPatchRequest) SetConfig(v interface{}) {
	o.Config = v
}

func (o PatchedConstrainPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JobId != nil {
		toSerialize["job_id"] = o.JobId
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedConstrainPatchRequest struct {
	value *PatchedConstrainPatchRequest
	isSet bool
}

func (v NullablePatchedConstrainPatchRequest) Get() *PatchedConstrainPatchRequest {
	return v.value
}

func (v *NullablePatchedConstrainPatchRequest) Set(val *PatchedConstrainPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedConstrainPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedConstrainPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedConstrainPatchRequest(val *PatchedConstrainPatchRequest) *NullablePatchedConstrainPatchRequest {
	return &NullablePatchedConstrainPatchRequest{value: val, isSet: true}
}

func (v NullablePatchedConstrainPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedConstrainPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


