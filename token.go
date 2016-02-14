package token

import (
	"io"
	"encoding/base64"
	"crypto/rand"
)

//New returns an unpadded, url-safe base64 encoding of a random byte slice of
//length n, defaulting to 256 bytes.
func New(n int) (string, error) {
	if n == 0 {
		n = 256
	}
	data := make([]byte, int(n))
	if _, err := io.ReadFull(rand.Reader, data); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(data), nil
}
