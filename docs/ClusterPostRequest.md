# ClusterPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Choice** | [**ClusterOffer**](ClusterOffer.md) |  | 
**Config** | [**ConstrainRequest**](ConstrainRequest.md) |  | 
**Project** | [**Project**](Project.md) |  | 

## Methods

### NewClusterPostRequest

`func NewClusterPostRequest(choice ClusterOffer, config ConstrainRequest, project Project, ) *ClusterPostRequest`

NewClusterPostRequest instantiates a new ClusterPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPostRequestWithDefaults

`func NewClusterPostRequestWithDefaults() *ClusterPostRequest`

NewClusterPostRequestWithDefaults instantiates a new ClusterPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoice

`func (o *ClusterPostRequest) GetChoice() ClusterOffer`

GetChoice returns the Choice field if non-nil, zero value otherwise.

### GetChoiceOk

`func (o *ClusterPostRequest) GetChoiceOk() (*ClusterOffer, bool)`

GetChoiceOk returns a tuple with the Choice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoice

`func (o *ClusterPostRequest) SetChoice(v ClusterOffer)`

SetChoice sets Choice field to given value.


### GetConfig

`func (o *ClusterPostRequest) GetConfig() ConstrainRequest`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ClusterPostRequest) GetConfigOk() (*ConstrainRequest, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ClusterPostRequest) SetConfig(v ConstrainRequest)`

SetConfig sets Config field to given value.


### GetProject

`func (o *ClusterPostRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ClusterPostRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ClusterPostRequest) SetProject(v Project)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


