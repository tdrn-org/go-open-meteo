//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

// Package openmeteo provides the necessary functions to access the Open-Meteo Weather API.
package openmeteo

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"runtime"
	"strings"

	"github.com/tdrn-org/go-open-meteo/api"
)

// ErrClientFailure indicates a system error while invoking the REST API.
var ErrClientFailure = errors.New("client call failure")

// ErrAPIFailure indicates an API error while invoking the REST API.
var ErrAPIFailure = errors.New("API failure")

// Client instances are used to access the REST API.
type Client struct {
	apiURL     *url.URL
	httpClient *http.Client
	apiClient  api.ClientWithResponsesInterface
	logger     *slog.Logger
}

// ClientOption interface is used to set Client options during Client creation.
type ClientOption interface {
	// Apply sets one or more options in the given Client instance.
	Apply(client *Client)
}

// ClientOptionFunc type is used to wrap functions into a ClientOption instance.
type ClientOptionFunc func(client *Client)

// Apply sets one or more options in the given Client instance.
func (f ClientOptionFunc) Apply(client *Client) {
	f(client)
}

// WithHttpClient sets a Client instance's [http.Client] which is used
// to access the REST API.
func WithHttpClient(httpClient *http.Client) ClientOptionFunc {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}

// WithLogger sets a Client instance's [slog.Logger] which is used
// for logging.
func WithLogger(logger *slog.Logger) ClientOptionFunc {
	return func(client *Client) {
		client.logger = logger
	}
}

// NewClient creates a new Client instance using the given
// connect URL as well as Client options.
func NewClient(apiURL *url.URL, options ...ClientOption) (*Client, error) {
	client := &Client{
		apiURL: apiURL,
	}
	for _, option := range options {
		option.Apply(client)
	}
	if client.httpClient == nil {
		client.httpClient = &http.Client{}
	}
	if client.logger == nil {
		client.logger = slog.With(slog.String("client", apiURL.String()))
	}
	httpClientOption := func(apiClient *api.Client) error {
		apiClient.Client = client.httpClient
		return nil
	}
	apiClient, err := api.NewClientWithResponses(apiURL.String(), api.WithRequestEditorFn(client.prepareRequest()), httpClientOption)
	if err != nil {
		return nil, fmt.Errorf("failed to create client (cause: %w)", err)
	}
	client.apiClient = apiClient
	return client, nil
}

func (client *Client) prepareRequest() api.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Accept", "application/json")
		return nil
	}
}

func (client *Client) operationID() string {
	pc, _, _, _ := runtime.Caller(3)
	caller := runtime.FuncForPC(pc)
	callerName := caller.Name()
	return callerName[strings.LastIndex(callerName, ".")+1:]
}

func (client *Client) wrapSystemError(err error) error {
	return fmt.Errorf("%w %s (cause: %w)", ErrClientFailure, client.operationID(), err)
}

func (client *Client) checkAPIResponse(httpResponse *http.Response) error {
	if httpResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("%w %s (status: %s)", ErrAPIFailure, client.operationID(), httpResponse.Status)
	}
	return nil
}
