# \ApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiConstrainCreate**](ApiAPI.md#ApiConstrainCreate) | **Post** /api/constrain | 
[**ApiConstrainPartialUpdate**](ApiAPI.md#ApiConstrainPartialUpdate) | **Patch** /api/constrain | 
[**ApiDownloadRetrieve**](ApiAPI.md#ApiDownloadRetrieve) | **Get** /api/download | 
[**ApiJobCreate**](ApiAPI.md#ApiJobCreate) | **Post** /api/job | 
[**ApiJobDestroy**](ApiAPI.md#ApiJobDestroy) | **Delete** /api/job | 
[**ApiJobRetrieve**](ApiAPI.md#ApiJobRetrieve) | **Get** /api/job | 
[**ApiProjectCreate**](ApiAPI.md#ApiProjectCreate) | **Post** /api/project | 
[**ApiProjectDestroy**](ApiAPI.md#ApiProjectDestroy) | **Delete** /api/project | 
[**ApiProjectJobsRetrieve**](ApiAPI.md#ApiProjectJobsRetrieve) | **Get** /api/project/{project_id}/jobs | 
[**ApiProjectRetrieve**](ApiAPI.md#ApiProjectRetrieve) | **Get** /api/project | 
[**ApiUserPublicKeyCreate**](ApiAPI.md#ApiUserPublicKeyCreate) | **Post** /api/user/public_key | 
[**ApiUserPublicKeyRetrieve**](ApiAPI.md#ApiUserPublicKeyRetrieve) | **Get** /api/user/public_key | 
[**ApiUserRetrieve**](ApiAPI.md#ApiUserRetrieve) | **Get** /api/user | 



## ApiConstrainCreate

> ConstrainResponse ApiConstrainCreate(ctx).ConstrainRequest(constrainRequest).Execute()





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
	resp, r, err := apiClient.ApiAPI.ApiConstrainCreate(context.Background()).ConstrainRequest(constrainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiConstrainCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiConstrainCreate`: ConstrainResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiConstrainCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiConstrainCreateRequest struct via the builder pattern


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


## ApiConstrainPartialUpdate

> ConstrainResponse ApiConstrainPartialUpdate(ctx).PatchedConstrainPatchRequest(patchedConstrainPatchRequest).Execute()





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
	resp, r, err := apiClient.ApiAPI.ApiConstrainPartialUpdate(context.Background()).PatchedConstrainPatchRequest(patchedConstrainPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiConstrainPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiConstrainPartialUpdate`: ConstrainResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiConstrainPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiConstrainPartialUpdateRequest struct via the builder pattern


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


## ApiDownloadRetrieve

> DownloadResponse ApiDownloadRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.ApiAPI.ApiDownloadRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiDownloadRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiDownloadRetrieve`: DownloadResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiDownloadRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiDownloadRetrieveRequest struct via the builder pattern


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


## ApiJobCreate

> JobGetRequest ApiJobCreate(ctx).JobGetRequest(jobGetRequest).Execute()





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
	resp, r, err := apiClient.ApiAPI.ApiJobCreate(context.Background()).JobGetRequest(jobGetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiJobCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiJobCreate`: JobGetRequest
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiJobCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiJobCreateRequest struct via the builder pattern


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


## ApiJobDestroy

> ApiJobDestroy(ctx).Execute()





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
	r, err := apiClient.ApiAPI.ApiJobDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiJobDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiJobDestroyRequest struct via the builder pattern


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


## ApiJobRetrieve

> JobGetResponse ApiJobRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.ApiAPI.ApiJobRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiJobRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiJobRetrieve`: JobGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiJobRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiJobRetrieveRequest struct via the builder pattern


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


## ApiProjectCreate

> Project ApiProjectCreate(ctx).ProjectCreateRequest(projectCreateRequest).Execute()





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
	resp, r, err := apiClient.ApiAPI.ApiProjectCreate(context.Background()).ProjectCreateRequest(projectCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiProjectCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiProjectCreate`: Project
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiProjectCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectCreateRequest struct via the builder pattern


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


## ApiProjectDestroy

> ApiProjectDestroy(ctx).Execute()





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
	r, err := apiClient.ApiAPI.ApiProjectDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiProjectDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectDestroyRequest struct via the builder pattern


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


## ApiProjectJobsRetrieve

> Job ApiProjectJobsRetrieve(ctx, projectId).Execute()





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
	resp, r, err := apiClient.ApiAPI.ApiProjectJobsRetrieve(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiProjectJobsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiProjectJobsRetrieve`: Job
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiProjectJobsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectJobsRetrieveRequest struct via the builder pattern


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


## ApiProjectRetrieve

> ProjectsGetResponse ApiProjectRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.ApiAPI.ApiProjectRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiProjectRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiProjectRetrieve`: ProjectsGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiProjectRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectRetrieveRequest struct via the builder pattern


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


## ApiUserPublicKeyCreate

> ApiUserPublicKeyCreate(ctx).UserPublicKeyRequest(userPublicKeyRequest).Execute()





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
	r, err := apiClient.ApiAPI.ApiUserPublicKeyCreate(context.Background()).UserPublicKeyRequest(userPublicKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiUserPublicKeyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiUserPublicKeyCreateRequest struct via the builder pattern


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


## ApiUserPublicKeyRetrieve

> UserPublicKeyResponse ApiUserPublicKeyRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.ApiAPI.ApiUserPublicKeyRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiUserPublicKeyRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiUserPublicKeyRetrieve`: UserPublicKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiUserPublicKeyRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserPublicKeyRetrieveRequest struct via the builder pattern


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


## ApiUserRetrieve

> UserResponse ApiUserRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.ApiAPI.ApiUserRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiUserRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiUserRetrieve`: UserResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiUserRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserRetrieveRequest struct via the builder pattern


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

