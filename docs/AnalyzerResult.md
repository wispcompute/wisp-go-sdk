# AnalyzerResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offer** | [**ClusterOffer**](ClusterOffer.md) |  | 
**Score** | [**Score**](Score.md) |  | 

## Methods

### NewAnalyzerResult

`func NewAnalyzerResult(offer ClusterOffer, score Score, ) *AnalyzerResult`

NewAnalyzerResult instantiates a new AnalyzerResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyzerResultWithDefaults

`func NewAnalyzerResultWithDefaults() *AnalyzerResult`

NewAnalyzerResultWithDefaults instantiates a new AnalyzerResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffer

`func (o *AnalyzerResult) GetOffer() ClusterOffer`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *AnalyzerResult) GetOfferOk() (*ClusterOffer, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *AnalyzerResult) SetOffer(v ClusterOffer)`

SetOffer sets Offer field to given value.


### GetScore

`func (o *AnalyzerResult) GetScore() Score`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *AnalyzerResult) GetScoreOk() (*Score, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *AnalyzerResult) SetScore(v Score)`

SetScore sets Score field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


