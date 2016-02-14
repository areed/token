package token

import (
	"io"
	"encoding/base64"
)

//New returns an unpadded, url-safe base64 encoding of a random byte slice of
//length n, defaulting to 256 bytes.
func New(n uint) (string, error) {
	if n == 0 {
		n = 256
	}
	data := make([]byte, int(n))
	if _, err := io.ReadFull(rand.Reader, data); err != nil {
		return "", err
	}
	s := base64.RawURLEncoding.EncodeToString(data)
	return s, nil
}
