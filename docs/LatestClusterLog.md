# LatestClusterLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** |  | [readonly] 
**Status** | [**LatestClusterLogStatusEnum**](LatestClusterLogStatusEnum.md) |  | 
**Message** | **string** |  | 

## Methods

### NewLatestClusterLog

`func NewLatestClusterLog(timestamp time.Time, status LatestClusterLogStatusEnum, message string, ) *LatestClusterLog`

NewLatestClusterLog instantiates a new LatestClusterLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLatestClusterLogWithDefaults

`func NewLatestClusterLogWithDefaults() *LatestClusterLog`

NewLatestClusterLogWithDefaults instantiates a new LatestClusterLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *LatestClusterLog) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LatestClusterLog) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LatestClusterLog) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetStatus

`func (o *LatestClusterLog) GetStatus() LatestClusterLogStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LatestClusterLog) GetStatusOk() (*LatestClusterLogStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LatestClusterLog) SetStatus(v LatestClusterLogStatusEnum)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *LatestClusterLog) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LatestClusterLog) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LatestClusterLog) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


