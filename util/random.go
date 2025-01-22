package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const alphaNum = "0123456789abcdefghijklmnopqrstuvwxyz"

//func init() {
//	rand.New(rand.NewSource(time.Now().UnixNano()))
//}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomStringN(min, max int64) string {
	c := RandomInt(min, max)
	return RandomString(int(c))
}

func RandomAccount() string {

	return RandomStringN(3, 8)
}

func RandomPassword() string {
	return RandomStringN(6, 12)
}
