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

// Outputs Outputs configuration serializer.
type Outputs struct {
	Buckets []Bucket `json:"buckets,omitempty"`
}

// NewOutputs instantiates a new Outputs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputs() *Outputs {
	this := Outputs{}
	return &this
}

// NewOutputsWithDefaults instantiates a new Outputs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputsWithDefaults() *Outputs {
	this := Outputs{}
	return &this
}

// GetBuckets returns the Buckets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Outputs) GetBuckets() []Bucket {
	if o == nil  {
		var ret []Bucket
		return ret
	}
	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Outputs) GetBucketsOk() (*[]Bucket, bool) {
	if o == nil || o.Buckets == nil {
		return nil, false
	}
	return &o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *Outputs) HasBuckets() bool {
	if o != nil && o.Buckets != nil {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given []Bucket and assigns it to the Buckets field.
func (o *Outputs) SetBuckets(v []Bucket) {
	o.Buckets = v
}

func (o Outputs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Buckets != nil {
		toSerialize["buckets"] = o.Buckets
	}
	return json.Marshal(toSerialize)
}

type NullableOutputs struct {
	value *Outputs
	isSet bool
}

func (v NullableOutputs) Get() *Outputs {
	return v.value
}

func (v *NullableOutputs) Set(val *Outputs) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputs) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputs(val *Outputs) *NullableOutputs {
	return &NullableOutputs{value: val, isSet: true}
}

func (v NullableOutputs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


