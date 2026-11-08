//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package openmeteo

import (
	_ "embed"
	"encoding/json"
)

type WMOCode int

func (code WMOCode) Icon() string {
	return wmoCodes[code].Icon
}

func (code WMOCode) Code() string {
	return wmoCodes[code].Code
}

//go:embed wmo_codes.json
var wmoCodesData []byte

type wmoCodesJSON map[WMOCode]wmoCodeInfo

type wmoCodeInfo struct {
	Icon string `json:"icon"`
	Code string `json:"code"`
}

var wmoCodes wmoCodesJSON

func init() {
	err := json.Unmarshal(wmoCodesData, &wmoCodes)
	if err != nil {
		panic(err)
	}
}
