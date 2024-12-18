# JobPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Choice** | [**ClusterOffer**](ClusterOffer.md) |  | 
**Config** | [**ConstrainRequest**](ConstrainRequest.md) |  | 
**Project** | [**Project**](Project.md) |  | [readonly] 

## Methods

### NewJobPostRequest

`func NewJobPostRequest(choice ClusterOffer, config ConstrainRequest, project Project, ) *JobPostRequest`

NewJobPostRequest instantiates a new JobPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobPostRequestWithDefaults

`func NewJobPostRequestWithDefaults() *JobPostRequest`

NewJobPostRequestWithDefaults instantiates a new JobPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoice

`func (o *JobPostRequest) GetChoice() ClusterOffer`

GetChoice returns the Choice field if non-nil, zero value otherwise.

### GetChoiceOk

`func (o *JobPostRequest) GetChoiceOk() (*ClusterOffer, bool)`

GetChoiceOk returns a tuple with the Choice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoice

`func (o *JobPostRequest) SetChoice(v ClusterOffer)`

SetChoice sets Choice field to given value.


### GetConfig

`func (o *JobPostRequest) GetConfig() ConstrainRequest`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *JobPostRequest) GetConfigOk() (*ConstrainRequest, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *JobPostRequest) SetConfig(v ConstrainRequest)`

SetConfig sets Config field to given value.


### GetProject

`func (o *JobPostRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *JobPostRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *JobPostRequest) SetProject(v Project)`

SetProject sets Project field to given value.


### SetProjectNil

`func (o *JobPostRequest) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *JobPostRequest) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


