# \ConstraintsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConstraintsCreate**](ConstraintsApi.md#ConstraintsCreate) | **Post** /api/constraints/ | 
[**ConstraintsPartialUpdate**](ConstraintsApi.md#ConstraintsPartialUpdate) | **Patch** /api/constraints/ | 



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
    openapiclient "./openapi"
)

func main() {
    constrainRequest := *openapiclient.NewConstrainRequest(*openapiclient.NewProject("Name_example", time.Now(), time.Now()), *openapiclient.NewResources([]*string{nil}), *openapiclient.NewIO()) // ConstrainRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConstraintsApi.ConstraintsCreate(context.Background()).ConstrainRequest(constrainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConstraintsApi.ConstraintsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConstraintsCreate`: ConstrainResponse
    fmt.Fprintf(os.Stdout, "Response from `ConstraintsApi.ConstraintsCreate`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConstraintsPartialUpdate

> ConstrainResponse ConstraintsPartialUpdate(ctx).PatchedConstrainPatchRequest(patchedConstrainPatchRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    patchedConstrainPatchRequest := *openapiclient.NewPatchedConstrainPatchRequest() // PatchedConstrainPatchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConstraintsApi.ConstraintsPartialUpdate(context.Background()).PatchedConstrainPatchRequest(patchedConstrainPatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConstraintsApi.ConstraintsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConstraintsPartialUpdate`: ConstrainResponse
    fmt.Fprintf(os.Stdout, "Response from `ConstraintsApi.ConstraintsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConstraintsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedConstrainPatchRequest** | [**PatchedConstrainPatchRequest**](PatchedConstrainPatchRequest.md) |  | 

### Return type

[**ConstrainResponse**](ConstrainResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

