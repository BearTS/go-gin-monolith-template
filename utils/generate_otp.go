package utils

import (
	"math/rand"
	"time"
)

const (
	CHARSET_NUMERIC = "0123456789"
)

func GenerateOtp(length int) string {
	if length == 0 {
		return ""
	}
	charset := CHARSET_NUMERIC

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	randomChar := charset[seededRand.Intn(len(charset))]
	return string(randomChar) + GenerateOtp(length-1)
}
