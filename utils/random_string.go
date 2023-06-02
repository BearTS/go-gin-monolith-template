package utils

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/BearTS/go-gin-monolith/constants"
)

func GenerateRandomStringFromSet(length int, chars string) string {
	if length <= 0 {
		return ""
	}
	if chars == "" {
		chars = constants.Charset.NUMERIC
	}

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	randomChar := chars[seededRand.Intn(len(chars))]
	return string(randomChar) + GenerateRandomStringFromSet(length-1, chars)
}

func GenerateShareCode() int {
	res := GenerateRandomStringFromSet(4, "")
	intVar, _ := strconv.Atoi(res)
	return intVar
}
