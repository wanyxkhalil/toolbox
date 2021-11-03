package mkpasswd

import (
	"fmt"
	"math/rand"
	"time"
)

var rander = rand.New(rand.NewSource(time.Now().UnixNano()))

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
	specialChars = `!"#$%&'()*+,-./:;<=>?@`
)

func GenerateString(lower, upper, digit, special uint) string {
	s1 := randStr(lower, lowerChars)
	s2 := randStr(upper, upperChars)
	s3 := randStr(digit, digits)
	s4 := randStr(special, specialChars)

	s := s1 + s2 + s3 + s4
	return shuffle(s)
}

func shuffle(s string) string {
	bytes := []byte(s)
	rander.Shuffle(len(s), func(i, j int) {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	})
	return string(bytes)
}

func randStr(l uint, chars string) string {
	n := len(chars)

	b := make([]byte, l)
	for i := range b {
		b[i] = chars[rander.Intn(n)]
	}
	return string(b)
}

// Run pad with lowercase letters if the sum of the options is too low.
func Run(length, lower, upper, digit, special uint) {
	if n := upper + digit + special; n < length {
		lower = length - n
	}

	fmt.Printf("%s\n", GenerateString(lower, upper, digit, special))
}
