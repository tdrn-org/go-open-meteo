// generated from api.ClientWithResponsesInterface
package openmeteo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	openmeteo "github.com/tdrn-org/go-open-meteo"
	"github.com/tdrn-org/go-open-meteo/api"
	"github.com/tdrn-org/go-open-meteo/mock"
)

func TestAPI(t *testing.T) {
	mockServer := mock.StartMockServer(&mockHandler{})
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

type mockHandler struct{}

func (h *mockHandler) GetV1Forecast(ctx context.Context, request api.GetV1ForecastRequestObject) (api.GetV1ForecastResponseObject, error) {
	return &api.GetV1Forecast200JSONResponse{}, nil
}
