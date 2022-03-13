package rand

import (
	"math/rand"
	"time"
)

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

const (
	CharsetNumeric      = "0123456789"
	CharsetAlphabet     = "abcdefghijklmnopqrstuvwxyz"
	CharsetAlphaNumeric = "0123456789abcdefghijklmnopqrstuvwxyz"
	CharsetSpecial      = "~!@#$%^&*()_+-=[]{};':,./<>?"
	CharsetAll          = "0123456789abcdefghijklmnopqrstuvwxyz~!@#$%^&*()_+-=[]{};':,./<>?"
)

func RandomNumeric(length int) string {
	return RandomString(length, CharsetNumeric)
}

func RandomAlphabet(length int) string {
	return RandomString(length, CharsetAlphabet)
}

func RandomAlphaNumeric(length int) string {
	return RandomString(length, CharsetAlphaNumeric)
}

func RandomString(length int, charset string) string {
	ret := make([]byte, length)
	for i := range ret {
		ret[i] = charset[seed.Intn(len(charset))]
	}
	return string(ret)
}
