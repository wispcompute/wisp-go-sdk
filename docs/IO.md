# IO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inputs** | Pointer to [**Inputs**](Inputs.md) |  | [optional] 
**Outputs** | Pointer to [**Outputs**](Outputs.md) |  | [optional] 

## Methods

### NewIO

`func NewIO() *IO`

NewIO instantiates a new IO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIOWithDefaults

`func NewIOWithDefaults() *IO`

NewIOWithDefaults instantiates a new IO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputs

`func (o *IO) GetInputs() Inputs`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *IO) GetInputsOk() (*Inputs, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *IO) SetInputs(v Inputs)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *IO) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### SetInputsNil

`func (o *IO) SetInputsNil(b bool)`

 SetInputsNil sets the value for Inputs to be an explicit nil

### UnsetInputs
`func (o *IO) UnsetInputs()`

UnsetInputs ensures that no value is present for Inputs, not even an explicit nil
### GetOutputs

`func (o *IO) GetOutputs() Outputs`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *IO) GetOutputsOk() (*Outputs, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *IO) SetOutputs(v Outputs)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *IO) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *IO) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *IO) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


