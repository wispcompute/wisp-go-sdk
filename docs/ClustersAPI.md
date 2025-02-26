# \ClustersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterActionCreate**](ClustersAPI.md#ClusterActionCreate) | **Post** /api/clusters/{cluster_name}/{action}/ | 
[**ClustersCreate**](ClustersAPI.md#ClustersCreate) | **Post** /api/clusters/ | 
[**ClustersRetrieve**](ClustersAPI.md#ClustersRetrieve) | **Get** /api/clusters/ | 
[**ClustersStatusRetrieve**](ClustersAPI.md#ClustersStatusRetrieve) | **Get** /api/clusters/{cluster_name}/status/ | 



## ClusterActionCreate

> ClusterActionResponse ClusterActionCreate(ctx, action, clusterName).Wait(wait).Execute()





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
	action := "action_example" // string | 
	clusterName := "clusterName_example" // string | 
	wait := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.ClusterActionCreate(context.Background(), action, clusterName).Wait(wait).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.ClusterActionCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClusterActionCreate`: ClusterActionResponse
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.ClusterActionCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**action** | **string** |  | 
**clusterName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClusterActionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wait** | **bool** |  | [default to false]

### Return type

[**ClusterActionResponse**](ClusterActionResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersCreate

> ClusterPostResponse ClustersCreate(ctx).ClusterPostRequest(clusterPostRequest).Execute()





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
	clusterPostRequest := *openapiclient.NewClusterPostRequest(*openapiclient.NewClusterOffer(), *openapiclient.NewConstrainRequest(*openapiclient.NewProject("Name_example", "ProjectId_example", "Type_example", time.Now(), time.Now()), *openapiclient.NewResources([]*string{nil}), "TODO"), *openapiclient.NewProject("Name_example", "ProjectId_example", "Type_example", time.Now(), time.Now())) // ClusterPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.ClustersCreate(context.Background()).ClusterPostRequest(clusterPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.ClustersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClustersCreate`: ClusterPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.ClustersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClustersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterPostRequest** | [**ClusterPostRequest**](ClusterPostRequest.md) |  | 

### Return type

[**ClusterPostResponse**](ClusterPostResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersRetrieve

> ClusterListResponse ClustersRetrieve(ctx).Constraints(constraints).Execute()





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
	constraints := *openapiclient.NewConstrainRequest(*openapiclient.NewProject("Name_example", "ProjectId_example", "Type_example", time.Now(), time.Now()), *openapiclient.NewResources([]*string{nil}), "TODO") // ConstrainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.ClustersRetrieve(context.Background()).Constraints(constraints).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.ClustersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClustersRetrieve`: ClusterListResponse
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.ClustersRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClustersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **constraints** | [**ConstrainRequest**](ConstrainRequest.md) |  | 

### Return type

[**ClusterListResponse**](ClusterListResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersStatusRetrieve

> ClusterStatusResponse ClustersStatusRetrieve(ctx, clusterName).Execute()





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
	clusterName := "clusterName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.ClustersStatusRetrieve(context.Background(), clusterName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.ClustersStatusRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClustersStatusRetrieve`: ClusterStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.ClustersStatusRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersStatusRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterStatusResponse**](ClusterStatusResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

