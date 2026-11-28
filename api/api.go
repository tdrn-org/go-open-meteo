//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

// Package api contains the generated OpenAPI spec objects and functions.
package api

//go:generate go run github.com/tdrn-org/go-oapi-tools/cmd/generate generate.conf

import (
	_ "embed"
)

// The OpenAPI spec yaml file used to generate the API code.
//
//go:embed forecast.yml
var OpenAPISpecYaml []byte
