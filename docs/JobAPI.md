# \JobAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJobWithoutCluster**](JobAPI.md#CreateJobWithoutCluster) | **Post** /api/job/ | 



## CreateJobWithoutCluster

> JobGetResponse CreateJobWithoutCluster(ctx).TempJobPostRequest(tempJobPostRequest).Execute()





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
	tempJobPostRequest := *openapiclient.NewTempJobPostRequest(*openapiclient.NewCluster(*openapiclient.NewUser("Email_example"), "Name_example", "TODO"), *openapiclient.NewConstrainRequest(*openapiclient.NewProject("Name_example", "ProjectId_example", "Type_example", time.Now(), time.Now()), *openapiclient.NewResources([]*string{nil}), "TODO"), *openapiclient.NewProject("Name_example", "ProjectId_example", "Type_example", time.Now(), time.Now())) // TempJobPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.CreateJobWithoutCluster(context.Background()).TempJobPostRequest(tempJobPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.CreateJobWithoutCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJobWithoutCluster`: JobGetResponse
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.CreateJobWithoutCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobWithoutClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tempJobPostRequest** | [**TempJobPostRequest**](TempJobPostRequest.md) |  | 

### Return type

[**JobGetResponse**](JobGetResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

