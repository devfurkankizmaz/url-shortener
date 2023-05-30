package utils

import (
	"math/rand"
	"time"
)

const (
	shortURLIDLength = 8
	charset          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateShortURLID() string {
	b := make([]byte, shortURLIDLength)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
