package utils

import (
	"log"
	"reflect"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

const (
	ERR_HASH = "[log_utils_helper_hashpassword_generatefrompassword]: %v"
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
func Normalize(str string) string {
	if str == "" {
		return str
	}

	return strings.ToUpper(RemoveDiacritics(str))
}

//	HashPassword hashes password
func HashPassword(pwd string) string {
	if pwd == "" {
		return pwd
	}

	pwd_slice := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(pwd_slice, bcrypt.MinCost)
	if err != nil {
		log.Fatalf(ERR_HASH, err)
	}

	return string(hash)
}

func ExistInSlice(slice interface{}, value interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid data-type!")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == value {
			return true
		}
	}

	return false
}
