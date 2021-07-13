package utils

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

//	RemoveDiacritics removes diacritics in text
func RemoveDiacritics(str string) string {
	if str == "" {
		return str
	}

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, str)
	return result
}

//	NormalizedWithUpper removes diacritics from str then upper it
func NormalizeWithUpper(str string) string {
	if str == "" {
		return str
	}

	return strings.ToUpper(RemoveDiacritics(str))
}
