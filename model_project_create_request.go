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

// ProjectCreateRequest Project create request serializer.
type ProjectCreateRequest struct {
	Name string `json:"name"`
}

// NewProjectCreateRequest instantiates a new ProjectCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCreateRequest(name string) *ProjectCreateRequest {
	this := ProjectCreateRequest{}
	this.Name = name
	return &this
}

// NewProjectCreateRequestWithDefaults instantiates a new ProjectCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCreateRequestWithDefaults() *ProjectCreateRequest {
	this := ProjectCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ProjectCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectCreateRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectCreateRequest) SetName(v string) {
	o.Name = v
}

func (o ProjectCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableProjectCreateRequest struct {
	value *ProjectCreateRequest
	isSet bool
}

func (v NullableProjectCreateRequest) Get() *ProjectCreateRequest {
	return v.value
}

func (v *NullableProjectCreateRequest) Set(val *ProjectCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCreateRequest(val *ProjectCreateRequest) *NullableProjectCreateRequest {
	return &NullableProjectCreateRequest{value: val, isSet: true}
}

func (v NullableProjectCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


