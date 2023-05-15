package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(value string) string {
	hash := sha256.New()
	hash.Write([]byte(value))
	hashedValue := hash.Sum(nil)
	hashString := hex.EncodeToString(hashedValue)
	return hashString
}
