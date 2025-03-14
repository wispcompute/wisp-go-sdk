# NotificationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**EmailOnSuccess** | Pointer to **bool** |  | [optional] 
**EmailOnFailure** | Pointer to **bool** |  | [optional] 
**EmailRecipient** | Pointer to **string** |  | [optional] 

## Methods

### NewNotificationConfig

`func NewNotificationConfig(id int32, ) *NotificationConfig`

NewNotificationConfig instantiates a new NotificationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationConfigWithDefaults

`func NewNotificationConfigWithDefaults() *NotificationConfig`

NewNotificationConfigWithDefaults instantiates a new NotificationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationConfig) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationConfig) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationConfig) SetId(v int32)`

SetId sets Id field to given value.


### GetEmailOnSuccess

`func (o *NotificationConfig) GetEmailOnSuccess() bool`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *NotificationConfig) GetEmailOnSuccessOk() (*bool, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *NotificationConfig) SetEmailOnSuccess(v bool)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *NotificationConfig) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *NotificationConfig) GetEmailOnFailure() bool`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *NotificationConfig) GetEmailOnFailureOk() (*bool, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *NotificationConfig) SetEmailOnFailure(v bool)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *NotificationConfig) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetEmailRecipient

`func (o *NotificationConfig) GetEmailRecipient() string`

GetEmailRecipient returns the EmailRecipient field if non-nil, zero value otherwise.

### GetEmailRecipientOk

`func (o *NotificationConfig) GetEmailRecipientOk() (*string, bool)`

GetEmailRecipientOk returns a tuple with the EmailRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRecipient

`func (o *NotificationConfig) SetEmailRecipient(v string)`

SetEmailRecipient sets EmailRecipient field to given value.

### HasEmailRecipient

`func (o *NotificationConfig) HasEmailRecipient() bool`

HasEmailRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


