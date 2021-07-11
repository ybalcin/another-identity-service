package utils

import (
	"testing"
)

func TestRemoveDiacritics(t *testing.T) {
	str := "éşçİüğ"
	str_removed_diacritics := "escIug"

	result := RemoveDiacritics(&str)

	if *result != str_removed_diacritics {
		t.Errorf("Could not removed diacritics! %v", *result)
	}
}
