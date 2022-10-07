package rs

import (
	"math/rand"
)

func GenerateRandomString(length int) string {

	str := make([]byte, length)
	charset := "abcdefghijklmnopqrstuvwxyz1234567890#!$%&"

	for i := range str {
		str[i] = charset[rand.Intn(len(charset))]
	}

	return string(str)
}
