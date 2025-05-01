# Score

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelPerf** | **float64** | Relative performance of the offer compared to the current configuration. A float value describing estimated proportional performance of the resource with the given workload. A value of 1 is equal performance, 2 is double performance and 0.5 is half performance. | 
**RelPrice** | **float64** | Relative price of the offer compared to the current configuration. A float value describing estimated proportional price of the resource with the given workload. A value of 1 is equal price, 2 is double price and 0.5 is half price. | 
**WeightedScore** | **float64** | Weighted score of the configuration taking into consideration the workload type, cost to performance weighting and more. Higher is better. | 

## Methods

### NewScore

`func NewScore(relPerf float64, relPrice float64, weightedScore float64, ) *Score`

NewScore instantiates a new Score object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScoreWithDefaults

`func NewScoreWithDefaults() *Score`

NewScoreWithDefaults instantiates a new Score object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelPerf

`func (o *Score) GetRelPerf() float64`

GetRelPerf returns the RelPerf field if non-nil, zero value otherwise.

### GetRelPerfOk

`func (o *Score) GetRelPerfOk() (*float64, bool)`

GetRelPerfOk returns a tuple with the RelPerf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelPerf

`func (o *Score) SetRelPerf(v float64)`

SetRelPerf sets RelPerf field to given value.


### GetRelPrice

`func (o *Score) GetRelPrice() float64`

GetRelPrice returns the RelPrice field if non-nil, zero value otherwise.

### GetRelPriceOk

`func (o *Score) GetRelPriceOk() (*float64, bool)`

GetRelPriceOk returns a tuple with the RelPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelPrice

`func (o *Score) SetRelPrice(v float64)`

SetRelPrice sets RelPrice field to given value.


### GetWeightedScore

`func (o *Score) GetWeightedScore() float64`

GetWeightedScore returns the WeightedScore field if non-nil, zero value otherwise.

### GetWeightedScoreOk

`func (o *Score) GetWeightedScoreOk() (*float64, bool)`

GetWeightedScoreOk returns a tuple with the WeightedScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedScore

`func (o *Score) SetWeightedScore(v float64)`

SetWeightedScore sets WeightedScore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


