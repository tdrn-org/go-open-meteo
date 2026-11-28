//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package openmeteo

import (
	"context"
	"net/http"
	"net/url"
)

// PublicEndpointURL is the Endpoint URL for the public Open-Meteo service.
var PublicEndpointURL *url.URL = func() *url.URL {
	url, err := url.Parse("https://api.open-meteo.com")
	if err != nil {
		panic(err)
	}
	return url
}()

// NewClient creates a new Client instance using the given
// endpoint URL as well as Client options.
func NewClient(apiURL *url.URL, options ...ClientOption) (*Client, error) {
	prepareReq := func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Accept", "application/json")
		return nil
	}
	return newClient(apiURL, append(options, WithRequestEditors(prepareReq))...)
}
