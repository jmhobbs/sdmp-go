package util

import (
	"crypto/sha512"
	"encoding/base64"
)

func SHA512(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}
