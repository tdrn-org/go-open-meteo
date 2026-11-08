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
	tests := []struct {
		name         string
		code         openmeteo.WMOCode
		expectedIcon string
		expectedCode string
	}{
		{
			name:         "snow grains",
			code:         openmeteo.WMOCode(77),
			expectedIcon: "❄️",
			expectedCode: "SNOW_GRAINS",
		},
		{
			name:         "clear sky",
			code:         openmeteo.WMOCode(0),
			expectedIcon: "☀️",
			expectedCode: "CLEAR_SKY",
		},
		{
			name:         "unmapped code",
			code:         openmeteo.WMOCode(42),
			expectedIcon: openmeteo.UnknownWMOIcon,
			expectedCode: openmeteo.UnknownWMOCode,
		},
		{
			name:         "code outside of range",
			code:         openmeteo.WMOCode(999),
			expectedIcon: openmeteo.UnknownWMOIcon,
			expectedCode: openmeteo.UnknownWMOCode,
		},
		{
			name:         "negative code",
			code:         openmeteo.WMOCode(-1),
			expectedIcon: openmeteo.UnknownWMOIcon,
			expectedCode: openmeteo.UnknownWMOCode,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expectedIcon, test.code.Icon())
			require.Equal(t, test.expectedCode, test.code.Code())
		})
	}
}
