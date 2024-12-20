# Go API client for wisp-sdk

Wisp API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate your own API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./wisp"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ConstraintsApi* | [**ConstraintsCreate**](docs/ConstraintsApi.md#constraintscreate) | **Post** /api/constraints/ | 
*ConstraintsApi* | [**ConstraintsPartialUpdate**](docs/ConstraintsApi.md#constraintspartialupdate) | **Patch** /api/constraints/ | 
*DownloadApi* | [**DownloadRetrieve**](docs/DownloadApi.md#downloadretrieve) | **Get** /api/download/ | 
*JobsApi* | [**CreateJob**](docs/JobsApi.md#createjob) | **Post** /api/jobs/ | 
*JobsApi* | [**JobsDestroy**](docs/JobsApi.md#jobsdestroy) | **Delete** /api/jobs/{job_id}/ | 
*JobsApi* | [**JobsRetrieve**](docs/JobsApi.md#jobsretrieve) | **Get** /api/jobs/{job_id}/ | 
*ProjectsApi* | [**CreateProject**](docs/ProjectsApi.md#createproject) | **Post** /api/projects/ | 
*ProjectsApi* | [**DeleteProject**](docs/ProjectsApi.md#deleteproject) | **Delete** /api/projects/{project_id}/ | 
*ProjectsApi* | [**ListProjects**](docs/ProjectsApi.md#listprojects) | **Get** /api/projects/ | 
*ProjectsApi* | [**ProjectsJobsRetrieve**](docs/ProjectsApi.md#projectsjobsretrieve) | **Get** /api/projects/{project_id}/jobs/ | 
*ProjectsApi* | [**RetrieveProject**](docs/ProjectsApi.md#retrieveproject) | **Get** /api/projects/{project_id}/ | 
*UsersApi* | [**UsersMePublicKeyCreate**](docs/UsersApi.md#usersmepublickeycreate) | **Post** /api/users/me/public-key/ | 
*UsersApi* | [**UsersMePublicKeyRetrieve**](docs/UsersApi.md#usersmepublickeyretrieve) | **Get** /api/users/me/public-key/ | 
*UsersApi* | [**UsersMeRetrieve**](docs/UsersApi.md#usersmeretrieve) | **Get** /api/users/me/ | 


## Documentation For Models

 - [Bucket](docs/Bucket.md)
 - [Cluster](docs/Cluster.md)
 - [ClusterOffer](docs/ClusterOffer.md)
 - [ClusterStatusEnum](docs/ClusterStatusEnum.md)
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
 - [TypeEnum](docs/TypeEnum.md)
 - [UserPublicKeyRequest](docs/UserPublicKeyRequest.md)
 - [UserPublicKeyResponse](docs/UserPublicKeyResponse.md)
 - [UserResponse](docs/UserResponse.md)


## Documentation For Authorization



### oauth2


- **Type**: OAuth
- **Flow**: password
- **Authorization URL**: 
- **Scopes**: 
 - **read**: Read access
 - **write**: Write access

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
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

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
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

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


### tokenAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


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



