package internal

import (
	"math/rand"
)

func GenerateRandomString() string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	l := 6
	s := ""

	for i := 0; i < l; i++ {
		p := rand.Int() % len(letters)
		s = s + string(letters[p])
	}

	return s
}
