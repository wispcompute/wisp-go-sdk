# \ProvisionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProvisionRetrieve**](ProvisionAPI.md#ProvisionRetrieve) | **Get** /api/provision/{provision_id} | 



## ProvisionRetrieve

> ProvisionLog ProvisionRetrieve(ctx, provisionId).Execute()





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
	provisionId := "provisionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisionAPI.ProvisionRetrieve(context.Background(), provisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisionAPI.ProvisionRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvisionRetrieve`: ProvisionLog
	fmt.Fprintf(os.Stdout, "Response from `ProvisionAPI.ProvisionRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisionRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProvisionLog**](ProvisionLog.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

