//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package openmeteo

import (
	_ "embed"
	"encoding/json"

	"golang.org/x/text/language"
)

//go:embed wmo_codes.json
var wmoCodesData []byte

type wmoCodesJSON map[int]struct {
	Icon        string            `json:"icon"`
	Description map[string]string `json:"description"`
}

var wmoCodes wmoCodesJSON

type WMOCode int

func (code WMOCode) Icon() string {
	return wmoCodes[int(code)].Icon
}

func (code WMOCode) Description(locale language.Tag) string {
	matchingLocale := MatchLocale(locale)
	return wmoCodes[int(code)].Description[matchingLocale.String()]
}

func init() {
	err := json.Unmarshal(wmoCodesData, &wmoCodes)
	if err != nil {
		panic(err)
	}
}
