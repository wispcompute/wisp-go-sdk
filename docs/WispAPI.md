# \WispAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WispConstrainCreate**](WispAPI.md#WispConstrainCreate) | **Post** /wisp/constrain | 
[**WispConstrainPartialUpdate**](WispAPI.md#WispConstrainPartialUpdate) | **Patch** /wisp/constrain | 
[**WispDownloadRetrieve**](WispAPI.md#WispDownloadRetrieve) | **Get** /wisp/download | 
[**WispJobCreate**](WispAPI.md#WispJobCreate) | **Post** /wisp/job | 
[**WispJobDestroy**](WispAPI.md#WispJobDestroy) | **Delete** /wisp/job | 
[**WispJobRetrieve**](WispAPI.md#WispJobRetrieve) | **Get** /wisp/job | 
[**WispProjectCreate**](WispAPI.md#WispProjectCreate) | **Post** /wisp/project | 
[**WispProjectDestroy**](WispAPI.md#WispProjectDestroy) | **Delete** /wisp/project | 
[**WispProjectJobsRetrieve**](WispAPI.md#WispProjectJobsRetrieve) | **Get** /wisp/project/{project_id}/jobs | 
[**WispProjectRetrieve**](WispAPI.md#WispProjectRetrieve) | **Get** /wisp/project | 
[**WispUserPublicKeyCreate**](WispAPI.md#WispUserPublicKeyCreate) | **Post** /wisp/user/public_key | 
[**WispUserPublicKeyRetrieve**](WispAPI.md#WispUserPublicKeyRetrieve) | **Get** /wisp/user/public_key | 
[**WispUserRetrieve**](WispAPI.md#WispUserRetrieve) | **Get** /wisp/user | 



## WispConstrainCreate

> ConstrainResponse WispConstrainCreate(ctx).ConstrainRequest(constrainRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wispcompute/wisp_go_sdk"
)

func main() {
	constrainRequest := *openapiclient.NewConstrainRequest("TODO", "TODO", "TODO", "TODO") // ConstrainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WispAPI.WispConstrainCreate(context.Background()).ConstrainRequest(constrainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WispAPI.WispConstrainCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WispConstrainCreate`: ConstrainResponse
	fmt.Fprintf(os.Stdout, "Response from `WispAPI.WispConstrainCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWispConstrainCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **constrainRequest** | [**ConstrainRequest**](ConstrainRequest.md) |  | 

### Return type

[**ConstrainResponse**](ConstrainResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WispConstrainPartialUpdate

> ConstrainResponse WispConstrainPartialUpdate(ctx).PatchedConstrainPatchRequest(patchedConstrainPatchRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wispcompute/wisp_go_sdk"
)

func main() {
	patchedConstrainPatchRequest := *openapiclient.NewPatchedConstrainPatchRequest() // PatchedConstrainPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WispAPI.WispConstrainPartialUpdate(context.Background()).PatchedConstrainPatchRequest(patchedConstrainPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WispAPI.WispConstrainPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WispConstrainPartialUpdate`: ConstrainResponse
	fmt.Fprintf(os.Stdout, "Response from `WispAPI.WispConstrainPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWispConstrainPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedConstrainPatchRequest** | [**PatchedConstrainPatchRequest**](PatchedConstrainPatchRequest.md) |  | 

### Return type

[**ConstrainResponse**](ConstrainResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WispDownloadRetrieve

> DownloadResponse WispDownloadRetrieve(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wispcompute/wisp_go_sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WispAPI.WispDownloadRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WispAPI.WispDownloadRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WispDownloadRetrieve`: DownloadResponse
	fmt.Fprintf(os.Stdout, "Response from `WispAPI.WispDownloadRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWispDownloadRetrieveRequest struct via the builder pattern


### Return type

[**DownloadResponse**](DownloadResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WispJobCreate

> JobGetRequest WispJobCreate(ctx).JobGetRequest(jobGetRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wispcompute/wisp_go_sdk"
)

func main() {
	jobGetRequest := *openapiclient.NewJobGetRequest("Id_example") // JobGetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WispAPI.WispJobCreate(context.Background()).JobGetRequest(jobGetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WispAPI.WispJobCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WispJobCreate`: JobGetRequest
	fmt.Fprintf(os.Stdout, "Response from `WispAPI.WispJobCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWispJobCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobGetRequest** | [**JobGetRequest**](JobGetRequest.md) |  | 

### Return type

[**JobGetRequest**](JobGetRequest.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WispJobDestroy

> WispJobDestroy(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wispcompute/wisp_go_sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WispAPI.WispJobDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WispAPI.WispJobDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWispJobDestroyRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WispJobRetrieve

> JobGetResponse WispJobRetrieve(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wispcompute/wisp_go_sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WispAPI.WispJobRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WispAPI.WispJobRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WispJobRetrieve`: JobGetResponse
	fmt.Fprintf(os.Stdout, "Response from `WispAPI.WispJobRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWispJobRetrieveRequest struct via the builder pattern


### Return type

[**JobGetResponse**](JobGetResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WispProjectCreate

> Project WispProjectCreate(ctx).ProjectCreateRequest(projectCreateRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wispcompute/wisp_go_sdk"
)

func main() {
	projectCreateRequest := *openapiclient.NewProjectCreateRequest("Name_example") // ProjectCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WispAPI.WispProjectCreate(context.Background()).ProjectCreateRequest(projectCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WispAPI.WispProjectCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WispProjectCreate`: Project
	fmt.Fprintf(os.Stdout, "Response from `WispAPI.WispProjectCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWispProjectCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectCreateRequest** | [**ProjectCreateRequest**](ProjectCreateRequest.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WispProjectDestroy

> WispProjectDestroy(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wispcompute/wisp_go_sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WispAPI.WispProjectDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WispAPI.WispProjectDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWispProjectDestroyRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WispProjectJobsRetrieve

> Job WispProjectJobsRetrieve(ctx, projectId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wispcompute/wisp_go_sdk"
)

func main() {
	projectId := "projectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WispAPI.WispProjectJobsRetrieve(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WispAPI.WispProjectJobsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WispProjectJobsRetrieve`: Job
	fmt.Fprintf(os.Stdout, "Response from `WispAPI.WispProjectJobsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWispProjectJobsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Job**](Job.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WispProjectRetrieve

> ProjectsGetResponse WispProjectRetrieve(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wispcompute/wisp_go_sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WispAPI.WispProjectRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WispAPI.WispProjectRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WispProjectRetrieve`: ProjectsGetResponse
	fmt.Fprintf(os.Stdout, "Response from `WispAPI.WispProjectRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWispProjectRetrieveRequest struct via the builder pattern


### Return type

[**ProjectsGetResponse**](ProjectsGetResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WispUserPublicKeyCreate

> WispUserPublicKeyCreate(ctx).UserPublicKeyRequest(userPublicKeyRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wispcompute/wisp_go_sdk"
)

func main() {
	userPublicKeyRequest := *openapiclient.NewUserPublicKeyRequest("PublicKey_example") // UserPublicKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WispAPI.WispUserPublicKeyCreate(context.Background()).UserPublicKeyRequest(userPublicKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WispAPI.WispUserPublicKeyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWispUserPublicKeyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userPublicKeyRequest** | [**UserPublicKeyRequest**](UserPublicKeyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WispUserPublicKeyRetrieve

> UserPublicKeyResponse WispUserPublicKeyRetrieve(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wispcompute/wisp_go_sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WispAPI.WispUserPublicKeyRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WispAPI.WispUserPublicKeyRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WispUserPublicKeyRetrieve`: UserPublicKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `WispAPI.WispUserPublicKeyRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWispUserPublicKeyRetrieveRequest struct via the builder pattern


### Return type

[**UserPublicKeyResponse**](UserPublicKeyResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WispUserRetrieve

> UserResponse WispUserRetrieve(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wispcompute/wisp_go_sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WispAPI.WispUserRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WispAPI.WispUserRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WispUserRetrieve`: UserResponse
	fmt.Fprintf(os.Stdout, "Response from `WispAPI.WispUserRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWispUserRetrieveRequest struct via the builder pattern


### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

