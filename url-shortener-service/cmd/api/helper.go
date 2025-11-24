package main

import (
	"crypto/sha256"
	"encoding/hex"
)

// GenerateUniqueStringFromLongUrlPath: generate url shortener
func (a *Config) GenerateUniqueStringFromLongUrlPath(urlPath string, length int) string {
	hasher := sha256.New()
	hasher.Write([]byte(urlPath))
	hashBytes := hasher.Sum(nil)

	hashHex := hex.EncodeToString(hashBytes)

	if length >= len(hashHex) {
		return hashHex
	}
	return hashHex[:length]

}
