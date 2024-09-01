package main

import (
	"math/rand"
)

func GenerateRandomString() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	l := 6
	str := make([]rune, l)

	for i := 0; i < l; i++ {
		p := rand.Int() % len(letters)
		str = append(str, letters[p])
	}

	return string(str)
}
