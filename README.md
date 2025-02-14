# Go API client for wisp

Wisp API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Generator version: 7.10.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import wisp "github.com/wispcompute/wisp-go-sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `wisp.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), wisp.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `wisp.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), wisp.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `wisp.ContextOperationServerIndices` and `wisp.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), wisp.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), wisp.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ClustersAPI* | [**ClustersCreate**](docs/ClustersAPI.md#clusterscreate) | **Post** /api/clusters/{cluster_name}/{action}/ | 
*ClustersAPI* | [**ClustersRetrieve**](docs/ClustersAPI.md#clustersretrieve) | **Get** /api/clusters/ | 
*ClustersAPI* | [**ClustersStatusRetrieve**](docs/ClustersAPI.md#clustersstatusretrieve) | **Get** /api/clusters/{cluster_name}/status/ | 
*ConstraintsAPI* | [**ConstraintsCreate**](docs/ConstraintsAPI.md#constraintscreate) | **Post** /api/constraints/ | 
*ConstraintsAPI* | [**ConstraintsPartialUpdate**](docs/ConstraintsAPI.md#constraintspartialupdate) | **Patch** /api/constraints/ | 
*DownloadAPI* | [**DownloadRetrieve**](docs/DownloadAPI.md#downloadretrieve) | **Get** /api/download/ | 
*JobsAPI* | [**CreateJob**](docs/JobsAPI.md#createjob) | **Post** /api/jobs/ | 
*JobsAPI* | [**JobsDestroy**](docs/JobsAPI.md#jobsdestroy) | **Delete** /api/jobs/{job_id}/ | 
*JobsAPI* | [**JobsRetrieve**](docs/JobsAPI.md#jobsretrieve) | **Get** /api/jobs/{job_id}/ | 
*ProjectsAPI* | [**CreateProject**](docs/ProjectsAPI.md#createproject) | **Post** /api/projects/ | 
*ProjectsAPI* | [**DeleteProject**](docs/ProjectsAPI.md#deleteproject) | **Delete** /api/projects/{project_id}/ | 
*ProjectsAPI* | [**ListProjects**](docs/ProjectsAPI.md#listprojects) | **Get** /api/projects/ | 
*ProjectsAPI* | [**ProjectsJobsRetrieve**](docs/ProjectsAPI.md#projectsjobsretrieve) | **Get** /api/projects/{project_id}/jobs/ | 
*ProjectsAPI* | [**RetrieveProject**](docs/ProjectsAPI.md#retrieveproject) | **Get** /api/projects/{project_id}/ | 
*UsersAPI* | [**UsersMePublicKeyCreate**](docs/UsersAPI.md#usersmepublickeycreate) | **Post** /api/users/me/public-key/ | 
*UsersAPI* | [**UsersMePublicKeyRetrieve**](docs/UsersAPI.md#usersmepublickeyretrieve) | **Get** /api/users/me/public-key/ | 
*UsersAPI* | [**UsersMeRetrieve**](docs/UsersAPI.md#usersmeretrieve) | **Get** /api/users/me/ | 


## Documentation For Models

 - [Bucket](docs/Bucket.md)
 - [Cluster](docs/Cluster.md)
 - [ClusterActionResponse](docs/ClusterActionResponse.md)
 - [ClusterListResponse](docs/ClusterListResponse.md)
 - [ClusterOffer](docs/ClusterOffer.md)
 - [ClusterStatusEnum](docs/ClusterStatusEnum.md)
 - [ClusterStatusResponse](docs/ClusterStatusResponse.md)
 - [ConstrainRequest](docs/ConstrainRequest.md)
 - [ConstrainResponse](docs/ConstrainResponse.md)
 - [DownloadResponse](docs/DownloadResponse.md)
 - [IO](docs/IO.md)
 - [Inputs](docs/Inputs.md)
 - [Job](docs/Job.md)
 - [JobGetResponse](docs/JobGetResponse.md)
 - [JobListResponse](docs/JobListResponse.md)
 - [JobPostRequest](docs/JobPostRequest.md)
 - [LatestClusterLog](docs/LatestClusterLog.md)
 - [LatestClusterLogStatusEnum](docs/LatestClusterLogStatusEnum.md)
 - [Outputs](docs/Outputs.md)
 - [PatchedConstrainPatchRequest](docs/PatchedConstrainPatchRequest.md)
 - [PickledHandleField](docs/PickledHandleField.md)
 - [Project](docs/Project.md)
 - [ProjectCreateRequest](docs/ProjectCreateRequest.md)
 - [ProjectsGetResponse](docs/ProjectsGetResponse.md)
 - [Resources](docs/Resources.md)
 - [Script](docs/Script.md)
 - [User](docs/User.md)
 - [UserPublicKeyRequest](docs/UserPublicKeyRequest.md)
 - [UserPublicKeyResponse](docs/UserPublicKeyResponse.md)
 - [UserResponse](docs/UserResponse.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### oauth2


- **Type**: OAuth
- **Flow**: password
- **Authorization URL**: 
- **Scopes**: 
 - **read**: Read access
 - **write**: Write access

Example

```go
auth := context.WithValue(context.Background(), wisp.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, wisp.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

### oauth2


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **read**: Read access
 - **write**: Write access

Example

```go
auth := context.WithValue(context.Background(), wisp.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, wisp.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: /o/authorize/
- **Scopes**: 
 - **read**: Read access
 - **write**: Write access

Example

```go
auth := context.WithValue(context.Background(), wisp.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, wisp.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

### tokenAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: tokenAuth and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		wisp.ContextAPIKeys,
		map[string]wisp.APIKey{
			"tokenAuth": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



