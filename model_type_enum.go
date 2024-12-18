/*
Wisp API

Wisp API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package wisp

import (
	"encoding/json"
	"fmt"
)

// TypeEnum * `local` - local * `git` - git * `docker` - docker
type TypeEnum string

// List of TypeEnum
const (
	LOCAL TypeEnum = "local"
	GIT TypeEnum = "git"
	DOCKER TypeEnum = "docker"
)

// All allowed values of TypeEnum enum
var AllowedTypeEnumEnumValues = []TypeEnum{
	"local",
	"git",
	"docker",
}

func (v *TypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeEnum(value)
	for _, existing := range AllowedTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeEnum", value)
}

// NewTypeEnumFromValue returns a pointer to a valid TypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeEnumFromValue(v string) (*TypeEnum, error) {
	ev := TypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypeEnum: valid values are %v", v, AllowedTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeEnum) IsValid() bool {
	for _, existing := range AllowedTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TypeEnum value
func (v TypeEnum) Ptr() *TypeEnum {
	return &v
}

type NullableTypeEnum struct {
	value *TypeEnum
	isSet bool
}

func (v NullableTypeEnum) Get() *TypeEnum {
	return v.value
}

func (v *NullableTypeEnum) Set(val *TypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypeEnum(val *TypeEnum) *NullableTypeEnum {
	return &NullableTypeEnum{value: val, isSet: true}
}

func (v NullableTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

