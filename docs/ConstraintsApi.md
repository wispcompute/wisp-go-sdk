# \ConstraintsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConstraintsCreate**](ConstraintsAPI.md#ConstraintsCreate) | **Post** /api/constraints/ | 
[**ConstraintsPartialUpdate**](ConstraintsAPI.md#ConstraintsPartialUpdate) | **Patch** /api/constraints/ | 



## ConstraintsCreate

> ConstrainResponse ConstraintsCreate(ctx).ConstrainRequest(constrainRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/wispcompute/wisp-go-sdk"
)

func main() {
	constrainRequest := *openapiclient.NewConstrainRequest(*openapiclient.NewProject("Name_example", "ProjectId_example", "Type_example", time.Now(), time.Now()), *openapiclient.NewResources([]*string{nil}), "TODO", "TODO") // ConstrainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstraintsAPI.ConstraintsCreate(context.Background()).ConstrainRequest(constrainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstraintsAPI.ConstraintsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConstraintsCreate`: ConstrainResponse
	fmt.Fprintf(os.Stdout, "Response from `ConstraintsAPI.ConstraintsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConstraintsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **constrainRequest** | [**ConstrainRequest**](ConstrainRequest.md) |  | 

### Return type

[**ConstrainResponse**](ConstrainResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConstraintsPartialUpdate

> ConstrainRequest ConstraintsPartialUpdate(ctx).PatchedConstrainPatchRequest(patchedConstrainPatchRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wispcompute/wisp-go-sdk"
)

func main() {
	patchedConstrainPatchRequest := *openapiclient.NewPatchedConstrainPatchRequest() // PatchedConstrainPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstraintsAPI.ConstraintsPartialUpdate(context.Background()).PatchedConstrainPatchRequest(patchedConstrainPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstraintsAPI.ConstraintsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConstraintsPartialUpdate`: ConstrainRequest
	fmt.Fprintf(os.Stdout, "Response from `ConstraintsAPI.ConstraintsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConstraintsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedConstrainPatchRequest** | [**PatchedConstrainPatchRequest**](PatchedConstrainPatchRequest.md) |  | 

### Return type

[**ConstrainRequest**](ConstrainRequest.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

