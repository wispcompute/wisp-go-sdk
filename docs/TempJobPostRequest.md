# TempJobPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**Cluster**](Cluster.md) |  | 
**Config** | [**ConstrainRequest**](ConstrainRequest.md) |  | 
**Project** | [**Project**](Project.md) |  | 

## Methods

### NewTempJobPostRequest

`func NewTempJobPostRequest(cluster Cluster, config ConstrainRequest, project Project, ) *TempJobPostRequest`

NewTempJobPostRequest instantiates a new TempJobPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTempJobPostRequestWithDefaults

`func NewTempJobPostRequestWithDefaults() *TempJobPostRequest`

NewTempJobPostRequestWithDefaults instantiates a new TempJobPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *TempJobPostRequest) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *TempJobPostRequest) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *TempJobPostRequest) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.


### GetConfig

`func (o *TempJobPostRequest) GetConfig() ConstrainRequest`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *TempJobPostRequest) GetConfigOk() (*ConstrainRequest, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *TempJobPostRequest) SetConfig(v ConstrainRequest)`

SetConfig sets Config field to given value.


### GetProject

`func (o *TempJobPostRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *TempJobPostRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *TempJobPostRequest) SetProject(v Project)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


