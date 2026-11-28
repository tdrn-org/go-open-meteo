//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

// Package mock provides a plugable mock implementation for the Open-Meteo Weather API service.
package mock

// StartMockServer starts and returns a new mock instance.
//
// StartMockServer panics in case of an error. The returned server
// is listening on localhost using a dynamic port.
// If the necessary handler function is available in the
// given mock handler, the mock instance invokes them.
// Otherwiese the mock call simply fails with [ErrNoMock].
func StartMockServer(mockHandler any) *Server {
	return startMockServer(mockHandler, "", nil)
}
