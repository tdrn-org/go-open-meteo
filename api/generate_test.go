//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package api_test

import (
	"bytes"
	"go/format"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-open-meteo/api"
)

func TestGenerateClientWrapper(t *testing.T) {
	buffer := &bytes.Buffer{}
	err := api.GenerateClientWrapper(buffer)
	require.NoError(t, err)
	codeBytes := buffer.Bytes()
	_, err = format.Source(codeBytes)
	require.NoError(t, err)
}

func TestGenerateMock(t *testing.T) {
	buffer := &bytes.Buffer{}
	err := api.GenerateMock(buffer)
	require.NoError(t, err)
	codeBytes := buffer.Bytes()
	_, err = format.Source(codeBytes)
	require.NoError(t, err)
}

func TestGenerateAPITest(t *testing.T) {
	buffer := &bytes.Buffer{}
	err := api.GenerateAPITest(buffer)
	require.NoError(t, err)
	codeBytes := buffer.Bytes()
	_, err = format.Source(codeBytes)
	require.NoError(t, err)
}
