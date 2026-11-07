//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package openmeteo

import "golang.org/x/text/language"

var supportedLocales []language.Tag = []language.Tag{language.English, language.German}

func DefaultLocale() language.Tag {
	return supportedLocales[0]
}

var languageMatcher language.Matcher = language.NewMatcher(supportedLocales)

func MatchLocale(l ...language.Tag) language.Tag {
	_, index, _ := languageMatcher.Match(l...)
	return supportedLocales[index]
}
