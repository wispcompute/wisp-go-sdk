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

// checks if the JobGetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobGetResponse{}

// JobGetResponse Job get response serializer.
type JobGetResponse struct {
	Job Job `json:"job"`
	LatestClusterLog NullableLatestClusterLog `json:"latest_cluster_log"`
}

type _JobGetResponse JobGetResponse

// NewJobGetResponse instantiates a new JobGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobGetResponse(job Job, latestClusterLog NullableLatestClusterLog) *JobGetResponse {
	this := JobGetResponse{}
	this.Job = job
	this.LatestClusterLog = latestClusterLog
	return &this
}

// NewJobGetResponseWithDefaults instantiates a new JobGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobGetResponseWithDefaults() *JobGetResponse {
	this := JobGetResponse{}
	return &this
}

// GetJob returns the Job field value
func (o *JobGetResponse) GetJob() Job {
	if o == nil {
		var ret Job
		return ret
	}

	return o.Job
}

// GetJobOk returns a tuple with the Job field value
// and a boolean to check if the value has been set.
func (o *JobGetResponse) GetJobOk() (*Job, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Job, true
}

// SetJob sets field value
func (o *JobGetResponse) SetJob(v Job) {
	o.Job = v
}

// GetLatestClusterLog returns the LatestClusterLog field value
// If the value is explicit nil, the zero value for LatestClusterLog will be returned
func (o *JobGetResponse) GetLatestClusterLog() LatestClusterLog {
	if o == nil || o.LatestClusterLog.Get() == nil {
		var ret LatestClusterLog
		return ret
	}

	return *o.LatestClusterLog.Get()
}

// GetLatestClusterLogOk returns a tuple with the LatestClusterLog field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobGetResponse) GetLatestClusterLogOk() (*LatestClusterLog, bool) {
	if o == nil {
		return nil, false
	}
	return o.LatestClusterLog.Get(), o.LatestClusterLog.IsSet()
}

// SetLatestClusterLog sets field value
func (o *JobGetResponse) SetLatestClusterLog(v LatestClusterLog) {
	o.LatestClusterLog.Set(&v)
}

func (o JobGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobGetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["job"] = o.Job
	toSerialize["latest_cluster_log"] = o.LatestClusterLog.Get()
	return toSerialize, nil
}

func (o *JobGetResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"job",
		"latest_cluster_log",
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

	varJobGetResponse := _JobGetResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJobGetResponse)

	if err != nil {
		return err
	}

	*o = JobGetResponse(varJobGetResponse)

	return err
}

type NullableJobGetResponse struct {
	value *JobGetResponse
	isSet bool
}

func (v NullableJobGetResponse) Get() *JobGetResponse {
	return v.value
}

func (v *NullableJobGetResponse) Set(val *JobGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJobGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJobGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobGetResponse(val *JobGetResponse) *NullableJobGetResponse {
	return &NullableJobGetResponse{value: val, isSet: true}
}

func (v NullableJobGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


