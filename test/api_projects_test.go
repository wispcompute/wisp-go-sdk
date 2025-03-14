/*
Wisp API

Testing ProjectsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/wispcompute/wisp-go-sdk"
)

func Test_openapi_ProjectsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectsAPIService CreateProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectsAPI.CreateProject(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService DeleteProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		httpRes, err := apiClient.ProjectsAPI.DeleteProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ListProjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectsAPI.ListProjects(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsJobsRetrieve", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsJobsRetrieve(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService RetrieveProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.RetrieveProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
