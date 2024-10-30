package util

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func GenAPIKey() string {
	randBytes := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, randBytes)
	if err != nil {
		panic(err)
	}

	return base64.RawURLEncoding.WithPadding(base64.NoPadding).EncodeToString(randBytes)
}
