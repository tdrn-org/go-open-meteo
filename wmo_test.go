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
	"golang.org/x/text/language"
)

func TestWMOcodes(t *testing.T) {
	code := openmeteo.WMOCode(77)
	codeIcon := code.Icon()
	require.Equal(t, "❄️", codeIcon)
	codeDescriptionEN := code.Description(language.English)
	require.Equal(t, "Snow grains", codeDescriptionEN)
	codeDescriptionDE := code.Description(language.German)
	require.Equal(t, "Schneegriesel", codeDescriptionDE)
	codeDescriptionGR := code.Description(language.Greek)
	require.Equal(t, "Snow grains", codeDescriptionGR)
}
