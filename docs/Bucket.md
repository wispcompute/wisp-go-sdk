# Bucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**Mode** | **NullableString** |  | 
**MountLocation** | **NullableString** |  | 

## Methods

### NewBucket

`func NewBucket(name NullableString, mode NullableString, mountLocation NullableString, ) *Bucket`

NewBucket instantiates a new Bucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketWithDefaults

`func NewBucketWithDefaults() *Bucket`

NewBucketWithDefaults instantiates a new Bucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Bucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bucket) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Bucket) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Bucket) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMode

`func (o *Bucket) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Bucket) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Bucket) SetMode(v string)`

SetMode sets Mode field to given value.


### SetModeNil

`func (o *Bucket) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *Bucket) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetMountLocation

`func (o *Bucket) GetMountLocation() string`

GetMountLocation returns the MountLocation field if non-nil, zero value otherwise.

### GetMountLocationOk

`func (o *Bucket) GetMountLocationOk() (*string, bool)`

GetMountLocationOk returns a tuple with the MountLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountLocation

`func (o *Bucket) SetMountLocation(v string)`

SetMountLocation sets MountLocation field to given value.


### SetMountLocationNil

`func (o *Bucket) SetMountLocationNil(b bool)`

 SetMountLocationNil sets the value for MountLocation to be an explicit nil

### UnsetMountLocation
`func (o *Bucket) UnsetMountLocation()`

UnsetMountLocation ensures that no value is present for MountLocation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


