# AutoStop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Enabled** | Pointer to **bool** |  | [optional] 
**TimeoutMinutes** | Pointer to **int32** |  | [optional] 

## Methods

### NewAutoStop

`func NewAutoStop(id int32, ) *AutoStop`

NewAutoStop instantiates a new AutoStop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoStopWithDefaults

`func NewAutoStopWithDefaults() *AutoStop`

NewAutoStopWithDefaults instantiates a new AutoStop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoStop) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoStop) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoStop) SetId(v int32)`

SetId sets Id field to given value.


### GetEnabled

`func (o *AutoStop) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoStop) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoStop) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AutoStop) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTimeoutMinutes

`func (o *AutoStop) GetTimeoutMinutes() int32`

GetTimeoutMinutes returns the TimeoutMinutes field if non-nil, zero value otherwise.

### GetTimeoutMinutesOk

`func (o *AutoStop) GetTimeoutMinutesOk() (*int32, bool)`

GetTimeoutMinutesOk returns a tuple with the TimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMinutes

`func (o *AutoStop) SetTimeoutMinutes(v int32)`

SetTimeoutMinutes sets TimeoutMinutes field to given value.

### HasTimeoutMinutes

`func (o *AutoStop) HasTimeoutMinutes() bool`

HasTimeoutMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


