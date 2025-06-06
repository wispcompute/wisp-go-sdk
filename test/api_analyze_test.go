/*
Wisp API

Testing AnalyzeAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package wisp

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/wispcompute/wisp-go-sdk"
)

func Test_wisp_AnalyzeAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AnalyzeAPIService AnalyzeCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AnalyzeAPI.AnalyzeCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
