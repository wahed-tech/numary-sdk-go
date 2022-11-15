/*
Ledger API

Testing ServerApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ledgerclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	client "github.com/numary/numary-sdk-go"
)

func Test_ledgerclient_ServerApiService(t *testing.T) {

	configuration := client.NewConfiguration()
	apiClient := client.NewAPIClient(configuration)

	t.Run("Test ServerApiService GetInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ServerApi.GetInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
