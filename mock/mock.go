//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

// Package mock provides a mock implementation of the Open-Meteo Weather API for testing.
package mock

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"sync"

	"github.com/tdrn-org/go-open-meteo/api"
)

// Server represents a mock instance.
type Server struct {
	httpListener net.Listener
	apiURL       *url.URL
	logger       *slog.Logger
	stoppedWG    sync.WaitGroup
	httpServer   *http.Server
}

// Start starts and returns a new mock instance.
//
// Start panics in case of an error. The returned server
// is listening on localhost using a dynamic port.
func Start() *Server {
	httpListener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatal(err)
	}
	address := httpListener.Addr().String()
	apiURL, err := url.Parse(fmt.Sprintf("http://%s", address))
	logger := slog.Default().With(slog.String("server", address))
	if err != nil {
		log.Fatal(err)
	}
	server := &Server{
		httpListener: httpListener,
		apiURL:       apiURL,
		logger:       logger,
	}
	server.setupHttpServer()
	server.stoppedWG.Go(server.listenAndServe)
	return server
}

// APIURL gets the API URL for this mock instance.
func (s *Server) APIURL() *url.URL {
	return s.apiURL
}

// Stop stops this mock instance.
func (s *Server) Stop(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
	s.stoppedWG.Wait()
}

func (s *Server) setupHttpServer() {
	strictHandlerMiddlewares := []api.StrictMiddlewareFunc{
		s.logOperationMiddleware,
	}
	strictHandlerOptions := api.StrictHTTPServerOptions{
		RequestErrorHandlerFunc:  s.errorHandler,
		ResponseErrorHandlerFunc: s.errorHandler,
	}
	strictHandler := api.NewStrictHandlerWithOptions(s, strictHandlerMiddlewares, strictHandlerOptions)
	handlerMiddlewares := []api.MiddlewareFunc{}
	handlerOptions := api.StdHTTPServerOptions{
		Middlewares:      handlerMiddlewares,
		ErrorHandlerFunc: s.errorHandler,
	}
	handler := api.HandlerWithOptions(strictHandler, handlerOptions)
	s.httpServer = &http.Server{
		Handler: handler,
	}
}

func (s *Server) logOperationMiddleware(f api.StrictHandlerFunc, operationID string) api.StrictHandlerFunc {
	s.logger.Info("mock call", slog.String("operation", operationID))
	return f
}

func (s *Server) errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	s.logger.Error("mock handler failure", slog.String("url", r.URL.String()), slog.Any("err", err))
	w.WriteHeader(http.StatusInternalServerError)
}

func (s *Server) listenAndServe() {
	s.logger.Info("http server starting...")
	err := s.httpServer.Serve(s.httpListener)
	if !errors.Is(err, http.ErrServerClosed) {
		s.logger.Error("http server failure", slog.Any("err", err))
		return
	}
	s.logger.Info("http server stopped")
}
