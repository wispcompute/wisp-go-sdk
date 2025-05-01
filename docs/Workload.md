# Workload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkloadType** | [**WorkloadTypeEnum**](WorkloadTypeEnum.md) |  | 
**WorkloadPrompt** | **string** | A prompt describing the workload type you&#39;re running. Can be a LLM model name with precision, web hosting and more. | 
**CostPerfWeight** | **float64** | Weight factor balancing cost vs performance. -1 prioritizes cost, 1 prioritizes performance. For a balanced result, select 0. | 

## Methods

### NewWorkload

`func NewWorkload(workloadType WorkloadTypeEnum, workloadPrompt string, costPerfWeight float64, ) *Workload`

NewWorkload instantiates a new Workload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadWithDefaults

`func NewWorkloadWithDefaults() *Workload`

NewWorkloadWithDefaults instantiates a new Workload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkloadType

`func (o *Workload) GetWorkloadType() WorkloadTypeEnum`

GetWorkloadType returns the WorkloadType field if non-nil, zero value otherwise.

### GetWorkloadTypeOk

`func (o *Workload) GetWorkloadTypeOk() (*WorkloadTypeEnum, bool)`

GetWorkloadTypeOk returns a tuple with the WorkloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadType

`func (o *Workload) SetWorkloadType(v WorkloadTypeEnum)`

SetWorkloadType sets WorkloadType field to given value.


### GetWorkloadPrompt

`func (o *Workload) GetWorkloadPrompt() string`

GetWorkloadPrompt returns the WorkloadPrompt field if non-nil, zero value otherwise.

### GetWorkloadPromptOk

`func (o *Workload) GetWorkloadPromptOk() (*string, bool)`

GetWorkloadPromptOk returns a tuple with the WorkloadPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadPrompt

`func (o *Workload) SetWorkloadPrompt(v string)`

SetWorkloadPrompt sets WorkloadPrompt field to given value.


### GetCostPerfWeight

`func (o *Workload) GetCostPerfWeight() float64`

GetCostPerfWeight returns the CostPerfWeight field if non-nil, zero value otherwise.

### GetCostPerfWeightOk

`func (o *Workload) GetCostPerfWeightOk() (*float64, bool)`

GetCostPerfWeightOk returns a tuple with the CostPerfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerfWeight

`func (o *Workload) SetCostPerfWeight(v float64)`

SetCostPerfWeight sets CostPerfWeight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


