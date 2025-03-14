# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**User**](User.md) |  | 
**Name** | **string** |  | 
**LaunchedAt** | Pointer to **NullableInt32** |  | [optional] 
**LastUse** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**ClusterStatusEnum**](ClusterStatusEnum.md) |  | [optional] 
**Autostop** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 
**ToDown** | Pointer to **bool** |  | [optional] 
**ClusterHash** | Pointer to **NullableString** |  | [optional] 
**Handle** | [**NullablePickledHandleField**](PickledHandleField.md) |  | 

## Methods

### NewCluster

`func NewCluster(user User, name string, handle NullablePickledHandleField, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *Cluster) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Cluster) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Cluster) SetUser(v User)`

SetUser sets User field to given value.


### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.


### GetLaunchedAt

`func (o *Cluster) GetLaunchedAt() int32`

GetLaunchedAt returns the LaunchedAt field if non-nil, zero value otherwise.

### GetLaunchedAtOk

`func (o *Cluster) GetLaunchedAtOk() (*int32, bool)`

GetLaunchedAtOk returns a tuple with the LaunchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchedAt

`func (o *Cluster) SetLaunchedAt(v int32)`

SetLaunchedAt sets LaunchedAt field to given value.

### HasLaunchedAt

`func (o *Cluster) HasLaunchedAt() bool`

HasLaunchedAt returns a boolean if a field has been set.

### SetLaunchedAtNil

`func (o *Cluster) SetLaunchedAtNil(b bool)`

 SetLaunchedAtNil sets the value for LaunchedAt to be an explicit nil

### UnsetLaunchedAt
`func (o *Cluster) UnsetLaunchedAt()`

UnsetLaunchedAt ensures that no value is present for LaunchedAt, not even an explicit nil
### GetLastUse

`func (o *Cluster) GetLastUse() string`

GetLastUse returns the LastUse field if non-nil, zero value otherwise.

### GetLastUseOk

`func (o *Cluster) GetLastUseOk() (*string, bool)`

GetLastUseOk returns a tuple with the LastUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUse

`func (o *Cluster) SetLastUse(v string)`

SetLastUse sets LastUse field to given value.

### HasLastUse

`func (o *Cluster) HasLastUse() bool`

HasLastUse returns a boolean if a field has been set.

### SetLastUseNil

`func (o *Cluster) SetLastUseNil(b bool)`

 SetLastUseNil sets the value for LastUse to be an explicit nil

### UnsetLastUse
`func (o *Cluster) UnsetLastUse()`

UnsetLastUse ensures that no value is present for LastUse, not even an explicit nil
### GetStatus

`func (o *Cluster) GetStatus() ClusterStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cluster) GetStatusOk() (*ClusterStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cluster) SetStatus(v ClusterStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Cluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAutostop

`func (o *Cluster) GetAutostop() int32`

GetAutostop returns the Autostop field if non-nil, zero value otherwise.

### GetAutostopOk

`func (o *Cluster) GetAutostopOk() (*int32, bool)`

GetAutostopOk returns a tuple with the Autostop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostop

`func (o *Cluster) SetAutostop(v int32)`

SetAutostop sets Autostop field to given value.

### HasAutostop

`func (o *Cluster) HasAutostop() bool`

HasAutostop returns a boolean if a field has been set.

### GetMetadata

`func (o *Cluster) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Cluster) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Cluster) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Cluster) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetToDown

`func (o *Cluster) GetToDown() bool`

GetToDown returns the ToDown field if non-nil, zero value otherwise.

### GetToDownOk

`func (o *Cluster) GetToDownOk() (*bool, bool)`

GetToDownOk returns a tuple with the ToDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDown

`func (o *Cluster) SetToDown(v bool)`

SetToDown sets ToDown field to given value.

### HasToDown

`func (o *Cluster) HasToDown() bool`

HasToDown returns a boolean if a field has been set.

### GetClusterHash

`func (o *Cluster) GetClusterHash() string`

GetClusterHash returns the ClusterHash field if non-nil, zero value otherwise.

### GetClusterHashOk

`func (o *Cluster) GetClusterHashOk() (*string, bool)`

GetClusterHashOk returns a tuple with the ClusterHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterHash

`func (o *Cluster) SetClusterHash(v string)`

SetClusterHash sets ClusterHash field to given value.

### HasClusterHash

`func (o *Cluster) HasClusterHash() bool`

HasClusterHash returns a boolean if a field has been set.

### SetClusterHashNil

`func (o *Cluster) SetClusterHashNil(b bool)`

 SetClusterHashNil sets the value for ClusterHash to be an explicit nil

### UnsetClusterHash
`func (o *Cluster) UnsetClusterHash()`

UnsetClusterHash ensures that no value is present for ClusterHash, not even an explicit nil
### GetHandle

`func (o *Cluster) GetHandle() PickledHandleField`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *Cluster) GetHandleOk() (*PickledHandleField, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *Cluster) SetHandle(v PickledHandleField)`

SetHandle sets Handle field to given value.


### SetHandleNil

`func (o *Cluster) SetHandleNil(b bool)`

 SetHandleNil sets the value for Handle to be an explicit nil

### UnsetHandle
`func (o *Cluster) UnsetHandle()`

UnsetHandle ensures that no value is present for Handle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


