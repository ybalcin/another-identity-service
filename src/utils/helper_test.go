package utils

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestRemoveDiacritics(t *testing.T) {
	str := "éşçİüğ"
	str_removed_diacritics := "escIug"

	result := RemoveDiacritics(str)

	if result != str_removed_diacritics {
		t.Errorf("Could not removed diacritics! %v", result)
	}
}

func TestHashPassword(t *testing.T) {
	pwd := "123456"

	result := HashPassword(pwd)
	if result == "" || result == pwd {
		t.Errorf("hash error")
	}

	err := bcrypt.CompareHashAndPassword([]byte(result), []byte(pwd))
	if err != nil {
		t.Errorf("\n hash and password compare result false!")
	}
}
