//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

// Package api contains the generated objects and functions of the Open-Meteo Weather API.
package api

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=oapi-codegen-open-meteo.yaml forecast.yml
//go:generate go run ../cmd/build/build.go genwrapper ../openmeteo.gen.go
//go:generate go run ../cmd/build/build.go genmock ../mock/mock.gen.go
//go:generate go run ../cmd/build/build.go genapitest ../openmeteo_test.go

import (
	_ "embed"
)

// The Open-Meteo Weather API spec yaml file used to generate the API code.
//
//go:embed forecast.yml
var OpenMeteoForecastRestApiSpecYaml []byte
