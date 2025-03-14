# JobConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Autostop** | [**AutoStop**](AutoStop.md) |  | 
**Notifications** | [**NotificationConfig**](NotificationConfig.md) |  | 

## Methods

### NewJobConfig

`func NewJobConfig(id int32, autostop AutoStop, notifications NotificationConfig, ) *JobConfig`

NewJobConfig instantiates a new JobConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobConfigWithDefaults

`func NewJobConfigWithDefaults() *JobConfig`

NewJobConfigWithDefaults instantiates a new JobConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobConfig) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobConfig) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobConfig) SetId(v int32)`

SetId sets Id field to given value.


### GetAutostop

`func (o *JobConfig) GetAutostop() AutoStop`

GetAutostop returns the Autostop field if non-nil, zero value otherwise.

### GetAutostopOk

`func (o *JobConfig) GetAutostopOk() (*AutoStop, bool)`

GetAutostopOk returns a tuple with the Autostop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostop

`func (o *JobConfig) SetAutostop(v AutoStop)`

SetAutostop sets Autostop field to given value.


### GetNotifications

`func (o *JobConfig) GetNotifications() NotificationConfig`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *JobConfig) GetNotificationsOk() (*NotificationConfig, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *JobConfig) SetNotifications(v NotificationConfig)`

SetNotifications sets Notifications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


