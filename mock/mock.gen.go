// Generated from api.StrictServerInterface
package mock

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"sync"

	"github.com/tdrn-org/go-open-meteo/api"
)

// ErrNoMock indicates a called handler is not mocked and simply fails with this error.
var ErrNoMock = errors.New("no mock")

// Server represents a mock instance.
type Server struct {
	httpListener net.Listener
	apiURL       *url.URL
	mockHandler  any
	logger       *slog.Logger
	stoppedWG    sync.WaitGroup
	httpServer   *http.Server
}

func startMockServer(mockHandler any, baseURL string, baseRouter api.ServeMux, middlewares ...api.MiddlewareFunc) *Server {
	httpListener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}
	address := httpListener.Addr().String()
	apiURL, err := url.Parse(fmt.Sprintf("http://%s", address))
	logger := slog.Default().With(slog.String("server", address))
	if err != nil {
		panic(err)
	}
	server := &Server{
		httpListener: httpListener,
		apiURL:       apiURL,
		mockHandler:  mockHandler,
		logger:       logger,
	}
	server.setupHttpServer(baseURL, baseRouter, middlewares...)
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

// See [api.GetV1Forecast]
func (s *Server) GetV1Forecast(ctx context.Context, request api.GetV1ForecastRequestObject) (api.GetV1ForecastResponseObject, error) {
	mockMethod := reflect.ValueOf(s.mockHandler).MethodByName("GetV1Forecast")
	if mockMethod.IsValid() {
		results := mockMethod.Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(request)})
		return safeCast[api.GetV1ForecastResponseObject](results[0], nil), safeCast[error](results[1], nil)
	}
	return nil, ErrNoMock
}

func safeCast[T any](v reflect.Value, vNil T) T {
	vInterface := v.Interface()
	if vInterface == nil {
		return vNil
	}
	return vInterface.(T)
}

func (s *Server) setupHttpServer(baseURL string, baseRouter api.ServeMux, middlewares ...api.MiddlewareFunc) {
	strictHandlerMiddlewares := []api.StrictMiddlewareFunc{
		s.logOperationMiddleware,
	}
	strictHandlerOptions := api.StrictHTTPServerOptions{
		RequestErrorHandlerFunc:  s.errorHandler,
		ResponseErrorHandlerFunc: s.errorHandler,
	}
	strictHandler := api.NewStrictHandlerWithOptions(s, strictHandlerMiddlewares, strictHandlerOptions)
	handlerOptions := api.StdHTTPServerOptions{
		BaseURL:          baseURL,
		BaseRouter:       baseRouter,
		Middlewares:      middlewares,
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
