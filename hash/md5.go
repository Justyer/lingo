package hash

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func FileMD5(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr), nil
}

func StrMD5(data string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(data))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
