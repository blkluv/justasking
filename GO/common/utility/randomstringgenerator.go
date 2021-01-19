package utility

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var gen = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandSeq returns a random string of length n
func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[gen.Intn(len(letters))]
	}
	return string(b)
}
