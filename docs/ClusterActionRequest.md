# ClusterActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **string** |  | 
**Action** | [**ActionEnum**](ActionEnum.md) |  | 
**Wait** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewClusterActionRequest

`func NewClusterActionRequest(clusterName string, action ActionEnum, ) *ClusterActionRequest`

NewClusterActionRequest instantiates a new ClusterActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterActionRequestWithDefaults

`func NewClusterActionRequestWithDefaults() *ClusterActionRequest`

NewClusterActionRequestWithDefaults instantiates a new ClusterActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *ClusterActionRequest) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ClusterActionRequest) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ClusterActionRequest) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetAction

`func (o *ClusterActionRequest) GetAction() ActionEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ClusterActionRequest) GetActionOk() (*ActionEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ClusterActionRequest) SetAction(v ActionEnum)`

SetAction sets Action field to given value.


### GetWait

`func (o *ClusterActionRequest) GetWait() bool`

GetWait returns the Wait field if non-nil, zero value otherwise.

### GetWaitOk

`func (o *ClusterActionRequest) GetWaitOk() (*bool, bool)`

GetWaitOk returns a tuple with the Wait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWait

`func (o *ClusterActionRequest) SetWait(v bool)`

SetWait sets Wait field to given value.

### HasWait

`func (o *ClusterActionRequest) HasWait() bool`

HasWait returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


