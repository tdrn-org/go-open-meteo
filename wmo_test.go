//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package openmeteo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	openmeteo "github.com/tdrn-org/go-open-meteo"
)

func TestWMOcodes(t *testing.T) {
	code := openmeteo.WMOCode(77)
	codeIcon := code.Icon()
	require.Equal(t, "❄️", codeIcon)
	codeName := code.Code()
	require.Equal(t, "SNOW_GRAINS", codeName)
}
