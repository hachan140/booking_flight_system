package helper

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256Hashing(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	hashedPass := h.Sum(nil)
	convertString := hex.EncodeToString(hashedPass)
	return convertString
}
