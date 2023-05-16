package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
)

var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func SHA256Hashing(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	hashedPass := h.Sum(nil)
	convertString := hex.EncodeToString(hashedPass)
	return convertString
}

func GetRandomString(n1 int, n2 int) string {
	s := make([]byte, 6)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return strconv.Itoa(n1) + strconv.Itoa(n2) + string(s)
}
