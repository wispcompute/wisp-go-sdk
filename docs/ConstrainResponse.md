# ConstrainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Choice** | [**[]ClusterOffer**](ClusterOffer.md) |  | 

## Methods

### NewConstrainResponse

`func NewConstrainResponse(choice []ClusterOffer, ) *ConstrainResponse`

NewConstrainResponse instantiates a new ConstrainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstrainResponseWithDefaults

`func NewConstrainResponseWithDefaults() *ConstrainResponse`

NewConstrainResponseWithDefaults instantiates a new ConstrainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoice

`func (o *ConstrainResponse) GetChoice() []ClusterOffer`

GetChoice returns the Choice field if non-nil, zero value otherwise.

### GetChoiceOk

`func (o *ConstrainResponse) GetChoiceOk() (*[]ClusterOffer, bool)`

GetChoiceOk returns a tuple with the Choice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoice

`func (o *ConstrainResponse) SetChoice(v []ClusterOffer)`

SetChoice sets Choice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


