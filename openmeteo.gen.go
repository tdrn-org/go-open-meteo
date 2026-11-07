// generated from api.ClientWithResponsesInterface
package openmeteo

import (
	"context"

	"github.com/tdrn-org/go-open-meteo/api"
)

func (client *Client) GetV1Forecast(ctx context.Context, params *api.GetV1ForecastParams) (*api.GetV1ForecastResponse, error) {
	response, err := client.apiClient.GetV1ForecastWithResponse(ctx, params)
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}
