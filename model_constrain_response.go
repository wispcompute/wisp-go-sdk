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

// ConstrainResponse Serializer for cluster constraint response.  Fields:     choice (list): List of cluster offers matching the constraints
type ConstrainResponse struct {
	Choice []ClusterOffer `json:"choice"`
}

// NewConstrainResponse instantiates a new ConstrainResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstrainResponse(choice []ClusterOffer) *ConstrainResponse {
	this := ConstrainResponse{}
	this.Choice = choice
	return &this
}

// NewConstrainResponseWithDefaults instantiates a new ConstrainResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstrainResponseWithDefaults() *ConstrainResponse {
	this := ConstrainResponse{}
	return &this
}

// GetChoice returns the Choice field value
func (o *ConstrainResponse) GetChoice() []ClusterOffer {
	if o == nil {
		var ret []ClusterOffer
		return ret
	}

	return o.Choice
}

// GetChoiceOk returns a tuple with the Choice field value
// and a boolean to check if the value has been set.
func (o *ConstrainResponse) GetChoiceOk() (*[]ClusterOffer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Choice, true
}

// SetChoice sets field value
func (o *ConstrainResponse) SetChoice(v []ClusterOffer) {
	o.Choice = v
}

func (o ConstrainResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["choice"] = o.Choice
	}
	return json.Marshal(toSerialize)
}

type NullableConstrainResponse struct {
	value *ConstrainResponse
	isSet bool
}

func (v NullableConstrainResponse) Get() *ConstrainResponse {
	return v.value
}

func (v *NullableConstrainResponse) Set(val *ConstrainResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConstrainResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConstrainResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstrainResponse(val *ConstrainResponse) *NullableConstrainResponse {
	return &NullableConstrainResponse{value: val, isSet: true}
}

func (v NullableConstrainResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstrainResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


