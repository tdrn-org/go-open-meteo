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

// UnknownWMOIcon is returned by [WMOCode.Icon] for codes without a mapping.
const UnknownWMOIcon = "❔"

// UnknownWMOCode is returned by [WMOCode.Code] for codes without a mapping.
const UnknownWMOCode = "UNKNOWN"

func (code WMOCode) Icon() string {
	info, ok := wmoCodes[code]
	if !ok {
		return UnknownWMOIcon
	}
	return info.Icon
}

func (code WMOCode) Code() string {
	info, ok := wmoCodes[code]
	if !ok {
		return UnknownWMOCode
	}
	return info.Code
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
