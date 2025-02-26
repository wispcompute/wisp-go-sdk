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

// ProvisionLogStatusEnum * `PENDING` - PROVISION_PENDING * `IN_PROGRESS` - PROVISION_IN_PROGRESS * `TRANSFER_SSH` - PROVISION_TRANSFER_SSH * `SETUP_BUCKETS` - PROVISION_SETUP_BUCKETS * `SYNC_FILE_MOUNTS` - PROVISION_SYNC_FILE_MOUNTS * `COMPLETED` - PROVISION_COMPLETED * `FAILED` - PROVISION_FAILED
type ProvisionLogStatusEnum string

// List of ProvisionLogStatusEnum
const (
	PENDING ProvisionLogStatusEnum = "PENDING"
	IN_PROGRESS ProvisionLogStatusEnum = "IN_PROGRESS"
	TRANSFER_SSH ProvisionLogStatusEnum = "TRANSFER_SSH"
	SETUP_BUCKETS ProvisionLogStatusEnum = "SETUP_BUCKETS"
	SYNC_FILE_MOUNTS ProvisionLogStatusEnum = "SYNC_FILE_MOUNTS"
	COMPLETED ProvisionLogStatusEnum = "COMPLETED"
	FAILED ProvisionLogStatusEnum = "FAILED"
)

// All allowed values of ProvisionLogStatusEnum enum
var AllowedProvisionLogStatusEnumEnumValues = []ProvisionLogStatusEnum{
	"PENDING",
	"IN_PROGRESS",
	"TRANSFER_SSH",
	"SETUP_BUCKETS",
	"SYNC_FILE_MOUNTS",
	"COMPLETED",
	"FAILED",
}

func (v *ProvisionLogStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProvisionLogStatusEnum(value)
	for _, existing := range AllowedProvisionLogStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProvisionLogStatusEnum", value)
}

// NewProvisionLogStatusEnumFromValue returns a pointer to a valid ProvisionLogStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProvisionLogStatusEnumFromValue(v string) (*ProvisionLogStatusEnum, error) {
	ev := ProvisionLogStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProvisionLogStatusEnum: valid values are %v", v, AllowedProvisionLogStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProvisionLogStatusEnum) IsValid() bool {
	for _, existing := range AllowedProvisionLogStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProvisionLogStatusEnum value
func (v ProvisionLogStatusEnum) Ptr() *ProvisionLogStatusEnum {
	return &v
}

type NullableProvisionLogStatusEnum struct {
	value *ProvisionLogStatusEnum
	isSet bool
}

func (v NullableProvisionLogStatusEnum) Get() *ProvisionLogStatusEnum {
	return v.value
}

func (v *NullableProvisionLogStatusEnum) Set(val *ProvisionLogStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionLogStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionLogStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionLogStatusEnum(val *ProvisionLogStatusEnum) *NullableProvisionLogStatusEnum {
	return &NullableProvisionLogStatusEnum{value: val, isSet: true}
}

func (v NullableProvisionLogStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionLogStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

