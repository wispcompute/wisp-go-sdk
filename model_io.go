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

// IO IO configuration serializer.
type IO struct {
	Inputs Inputs `json:"inputs,omitempty"`
	Outputs Outputs `json:"outputs,omitempty"`
}

// NewIO instantiates a new IO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIO() *IO {
	this := IO{}
	return &this
}

// NewIOWithDefaults instantiates a new IO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIOWithDefaults() *IO {
	this := IO{}
	return &this
}

// GetInputs returns the Inputs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IO) GetInputs() Inputs {
	if o == nil  {
		var ret Inputs
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IO) GetInputsOk() (*Inputs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *IO) HasInputs() bool {
	if o != nil {
		return true
	}

	return false
}

// SetInputs gets a reference to the given Inputs and assigns it to the Inputs field.
func (o *IO) SetInputs(v Inputs) {
	o.Inputs = v
}

// GetOutputs returns the Outputs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IO) GetOutputs() Outputs {
	if o == nil  {
		var ret Outputs
		return ret
	}
	return o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IO) GetOutputsOk() (*Outputs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *IO) HasOutputs() bool {
	if o != nil {
		return true
	}

	return false
}

// SetOutputs gets a reference to the given Outputs and assigns it to the Outputs field.
func (o *IO) SetOutputs(v Outputs) {
	o.Outputs = v
}

func (o IO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	toSerialize["inputs"] = o.Inputs
	toSerialize["outputs"] = o.Outputs

	return json.Marshal(toSerialize)
}

type NullableIO struct {
	value *IO
	isSet bool
}

func (v NullableIO) Get() *IO {
	return v.value
}

func (v *NullableIO) Set(val *IO) {
	v.value = val
	v.isSet = true
}

func (v NullableIO) IsSet() bool {
	return v.isSet
}

func (v *NullableIO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIO(val *IO) *NullableIO {
	return &NullableIO{value: val, isSet: true}
}

func (v NullableIO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


