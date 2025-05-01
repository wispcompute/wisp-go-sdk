# \AnalyzeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyzeCreate**](AnalyzeAPI.md#AnalyzeCreate) | **Post** /api/analyze | 



## AnalyzeCreate

> AnalyzeGetResponse AnalyzeCreate(ctx).AnalyzeGetRequest(analyzeGetRequest).Execute()





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
	analyzeGetRequest := *openapiclient.NewAnalyzeGetRequest(*openapiclient.NewWorkload(openapiclient.WorkloadTypeEnum("batch"), "WorkloadPrompt_example", float64(123)), *openapiclient.NewClusterOffer(), *openapiclient.NewResources([]*string{nil})) // AnalyzeGetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyzeAPI.AnalyzeCreate(context.Background()).AnalyzeGetRequest(analyzeGetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyzeAPI.AnalyzeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyzeCreate`: AnalyzeGetResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyzeAPI.AnalyzeCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyzeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyzeGetRequest** | [**AnalyzeGetRequest**](AnalyzeGetRequest.md) |  | 

### Return type

[**AnalyzeGetResponse**](AnalyzeGetResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

