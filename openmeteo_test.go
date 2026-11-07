// generated from api.ClientWithResponsesInterface
package openmeteo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-open-meteo"
	"github.com/tdrn-org/go-open-meteo/api"
	"github.com/tdrn-org/go-open-meteo/mock"
)

func TestAPI(t *testing.T) {
	mockServer := mock.Start()
	defer mockServer.Stop(t.Context())

	client, err := openmeteo.NewClient(mockServer.APIURL())
	require.NoError(t, err)

	{
		ctx := t.Context()
		params := &api.GetV1ForecastParams{}
		_, err := client.GetV1Forecast(ctx, params)
		require.NoError(t, err)
	}
}
