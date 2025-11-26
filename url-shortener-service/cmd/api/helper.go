package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func (a *Config) CreateSalt() (string, error) {
	bytes := make([]byte, 10)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func (a *Config) GenerateUniqueStringFromLongUrlPath(urlPath string, length int) string {
	saltStr, err := a.CreateSalt()
	if err != nil {
		panic(err)
	}
	hasher := sha256.New()
	hasher.Write([]byte(urlPath + saltStr))
	hashBytes := hasher.Sum(nil)

	hashHex := hex.EncodeToString(hashBytes)

	if length >= len(hashHex) {
		return hashHex
	}
	return hashHex[:length]

}
