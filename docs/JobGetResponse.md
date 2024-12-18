# JobGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | [**Job**](Job.md) |  | 
**LatestClusterLog** | [**LatestClusterLog**](LatestClusterLog.md) |  | 

## Methods

### NewJobGetResponse

`func NewJobGetResponse(job Job, latestClusterLog LatestClusterLog, ) *JobGetResponse`

NewJobGetResponse instantiates a new JobGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobGetResponseWithDefaults

`func NewJobGetResponseWithDefaults() *JobGetResponse`

NewJobGetResponseWithDefaults instantiates a new JobGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *JobGetResponse) GetJob() Job`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *JobGetResponse) GetJobOk() (*Job, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *JobGetResponse) SetJob(v Job)`

SetJob sets Job field to given value.


### GetLatestClusterLog

`func (o *JobGetResponse) GetLatestClusterLog() LatestClusterLog`

GetLatestClusterLog returns the LatestClusterLog field if non-nil, zero value otherwise.

### GetLatestClusterLogOk

`func (o *JobGetResponse) GetLatestClusterLogOk() (*LatestClusterLog, bool)`

GetLatestClusterLogOk returns a tuple with the LatestClusterLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestClusterLog

`func (o *JobGetResponse) SetLatestClusterLog(v LatestClusterLog)`

SetLatestClusterLog sets LatestClusterLog field to given value.


### SetLatestClusterLogNil

`func (o *JobGetResponse) SetLatestClusterLogNil(b bool)`

 SetLatestClusterLogNil sets the value for LatestClusterLog to be an explicit nil

### UnsetLatestClusterLog
`func (o *JobGetResponse) UnsetLatestClusterLog()`

UnsetLatestClusterLog ensures that no value is present for LatestClusterLog, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


