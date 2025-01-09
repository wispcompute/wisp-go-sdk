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

// checks if the ConstrainResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstrainResponse{}

// ConstrainResponse Serializer for cluster constraint response.  Fields:     choice (list): List of cluster offers matching the constraints
type ConstrainResponse struct {
	Choice []ClusterOffer `json:"choice"`
}

type _ConstrainResponse ConstrainResponse

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
func (o *ConstrainResponse) GetChoiceOk() ([]ClusterOffer, bool) {
	if o == nil {
		return nil, false
	}
	return o.Choice, true
}

// SetChoice sets field value
func (o *ConstrainResponse) SetChoice(v []ClusterOffer) {
	o.Choice = v
}

func (o ConstrainResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstrainResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["choice"] = o.Choice
	return toSerialize, nil
}

func (o *ConstrainResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"choice",
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

	varConstrainResponse := _ConstrainResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConstrainResponse)

	if err != nil {
		return err
	}

	*o = ConstrainResponse(varConstrainResponse)

	return err
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


