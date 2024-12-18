/*
Wisp API

Wisp API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package wisp-sdk

import (
	"encoding/json"
)

// JobListResponse Job list response serializer.
type JobListResponse struct {
	Jobs []Job `json:"jobs"`
}

// NewJobListResponse instantiates a new JobListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobListResponse(jobs []Job) *JobListResponse {
	this := JobListResponse{}
	this.Jobs = jobs
	return &this
}

// NewJobListResponseWithDefaults instantiates a new JobListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobListResponseWithDefaults() *JobListResponse {
	this := JobListResponse{}
	return &this
}

// GetJobs returns the Jobs field value
func (o *JobListResponse) GetJobs() []Job {
	if o == nil {
		var ret []Job
		return ret
	}

	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value
// and a boolean to check if the value has been set.
func (o *JobListResponse) GetJobsOk() (*[]Job, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Jobs, true
}

// SetJobs sets field value
func (o *JobListResponse) SetJobs(v []Job) {
	o.Jobs = v
}

func (o JobListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["jobs"] = o.Jobs
	}
	return json.Marshal(toSerialize)
}

type NullableJobListResponse struct {
	value *JobListResponse
	isSet bool
}

func (v NullableJobListResponse) Get() *JobListResponse {
	return v.value
}

func (v *NullableJobListResponse) Set(val *JobListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJobListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJobListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobListResponse(val *JobListResponse) *NullableJobListResponse {
	return &NullableJobListResponse{value: val, isSet: true}
}

func (v NullableJobListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


