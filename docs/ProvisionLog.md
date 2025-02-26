# ProvisionLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**Cluster**](Cluster.md) |  | 
**ProvisionId** | **string** |  | 
**StepName** | **string** |  | 
**Status** | Pointer to [**ProvisionLogStatusEnum**](ProvisionLogStatusEnum.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**StartedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 

## Methods

### NewProvisionLog

`func NewProvisionLog(cluster Cluster, provisionId string, stepName string, startedAt time.Time, updatedAt time.Time, ) *ProvisionLog`

NewProvisionLog instantiates a new ProvisionLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionLogWithDefaults

`func NewProvisionLogWithDefaults() *ProvisionLog`

NewProvisionLogWithDefaults instantiates a new ProvisionLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *ProvisionLog) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ProvisionLog) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ProvisionLog) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.


### GetProvisionId

`func (o *ProvisionLog) GetProvisionId() string`

GetProvisionId returns the ProvisionId field if non-nil, zero value otherwise.

### GetProvisionIdOk

`func (o *ProvisionLog) GetProvisionIdOk() (*string, bool)`

GetProvisionIdOk returns a tuple with the ProvisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionId

`func (o *ProvisionLog) SetProvisionId(v string)`

SetProvisionId sets ProvisionId field to given value.


### GetStepName

`func (o *ProvisionLog) GetStepName() string`

GetStepName returns the StepName field if non-nil, zero value otherwise.

### GetStepNameOk

`func (o *ProvisionLog) GetStepNameOk() (*string, bool)`

GetStepNameOk returns a tuple with the StepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepName

`func (o *ProvisionLog) SetStepName(v string)`

SetStepName sets StepName field to given value.


### GetStatus

`func (o *ProvisionLog) GetStatus() ProvisionLogStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProvisionLog) GetStatusOk() (*ProvisionLogStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProvisionLog) SetStatus(v ProvisionLogStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProvisionLog) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *ProvisionLog) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProvisionLog) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProvisionLog) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProvisionLog) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ProvisionLog) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ProvisionLog) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetStartedAt

`func (o *ProvisionLog) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ProvisionLog) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ProvisionLog) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetUpdatedAt

`func (o *ProvisionLog) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProvisionLog) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProvisionLog) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


