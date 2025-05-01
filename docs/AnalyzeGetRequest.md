# AnalyzeGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workload** | [**Workload**](Workload.md) |  | 
**CurrentResources** | [**ClusterOffer**](ClusterOffer.md) |  | 
**Constraints** | [**Resources**](Resources.md) |  | 

## Methods

### NewAnalyzeGetRequest

`func NewAnalyzeGetRequest(workload Workload, currentResources ClusterOffer, constraints Resources, ) *AnalyzeGetRequest`

NewAnalyzeGetRequest instantiates a new AnalyzeGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyzeGetRequestWithDefaults

`func NewAnalyzeGetRequestWithDefaults() *AnalyzeGetRequest`

NewAnalyzeGetRequestWithDefaults instantiates a new AnalyzeGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkload

`func (o *AnalyzeGetRequest) GetWorkload() Workload`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *AnalyzeGetRequest) GetWorkloadOk() (*Workload, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *AnalyzeGetRequest) SetWorkload(v Workload)`

SetWorkload sets Workload field to given value.


### GetCurrentResources

`func (o *AnalyzeGetRequest) GetCurrentResources() ClusterOffer`

GetCurrentResources returns the CurrentResources field if non-nil, zero value otherwise.

### GetCurrentResourcesOk

`func (o *AnalyzeGetRequest) GetCurrentResourcesOk() (*ClusterOffer, bool)`

GetCurrentResourcesOk returns a tuple with the CurrentResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentResources

`func (o *AnalyzeGetRequest) SetCurrentResources(v ClusterOffer)`

SetCurrentResources sets CurrentResources field to given value.


### GetConstraints

`func (o *AnalyzeGetRequest) GetConstraints() Resources`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *AnalyzeGetRequest) GetConstraintsOk() (*Resources, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *AnalyzeGetRequest) SetConstraints(v Resources)`

SetConstraints sets Constraints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


