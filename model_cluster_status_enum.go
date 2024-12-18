/*
Wisp API

Wisp API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package wisp-sdk

import (
	"encoding/json"
	"fmt"
)

// ClusterStatusEnum * `INIT` - INIT * `UP` - UP * `STOPPED` - STOPPED * `ERROR` - ERROR
type ClusterStatusEnum string

// List of ClusterStatusEnum
const (
	INIT ClusterStatusEnum = "INIT"
	UP ClusterStatusEnum = "UP"
	STOPPED ClusterStatusEnum = "STOPPED"
	ERROR ClusterStatusEnum = "ERROR"
)

// All allowed values of ClusterStatusEnum enum
var AllowedClusterStatusEnumEnumValues = []ClusterStatusEnum{
	"INIT",
	"UP",
	"STOPPED",
	"ERROR",
}

func (v *ClusterStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClusterStatusEnum(value)
	for _, existing := range AllowedClusterStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClusterStatusEnum", value)
}

// NewClusterStatusEnumFromValue returns a pointer to a valid ClusterStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClusterStatusEnumFromValue(v string) (*ClusterStatusEnum, error) {
	ev := ClusterStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClusterStatusEnum: valid values are %v", v, AllowedClusterStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClusterStatusEnum) IsValid() bool {
	for _, existing := range AllowedClusterStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterStatusEnum value
func (v ClusterStatusEnum) Ptr() *ClusterStatusEnum {
	return &v
}

type NullableClusterStatusEnum struct {
	value *ClusterStatusEnum
	isSet bool
}

func (v NullableClusterStatusEnum) Get() *ClusterStatusEnum {
	return v.value
}

func (v *NullableClusterStatusEnum) Set(val *ClusterStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterStatusEnum(val *ClusterStatusEnum) *NullableClusterStatusEnum {
	return &NullableClusterStatusEnum{value: val, isSet: true}
}

func (v NullableClusterStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

