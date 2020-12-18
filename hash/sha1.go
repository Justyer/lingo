package hash

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

func StrSHA1(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

func HmacSHA1(s, k string) string {
	key := []byte(k)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(s))
	// return hex.EncodeToString(mac.Sum(nil))
	return string(mac.Sum(nil))
}
