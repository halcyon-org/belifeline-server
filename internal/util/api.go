package util

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

const length = 64

func GenAPIKey() string {
	randBytes := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, randBytes); err != nil {
		panic(err)
	}

	return base64.RawURLEncoding.WithPadding(base64.NoPadding).EncodeToString(randBytes)
}
