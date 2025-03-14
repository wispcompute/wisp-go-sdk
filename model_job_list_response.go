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

// checks if the JobListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobListResponse{}

// JobListResponse Job list response serializer.
type JobListResponse struct {
	Jobs []Job `json:"jobs"`
}

type _JobListResponse JobListResponse

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
func (o *JobListResponse) GetJobsOk() ([]Job, bool) {
	if o == nil {
		return nil, false
	}
	return o.Jobs, true
}

// SetJobs sets field value
func (o *JobListResponse) SetJobs(v []Job) {
	o.Jobs = v
}

func (o JobListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobs"] = o.Jobs
	return toSerialize, nil
}

func (o *JobListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"jobs",
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

	varJobListResponse := _JobListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJobListResponse)

	if err != nil {
		return err
	}

	*o = JobListResponse(varJobListResponse)

	return err
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


