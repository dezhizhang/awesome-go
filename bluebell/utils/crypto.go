package utils

import (
	"crypto/md5"
	"encoding/hex"
)

const salt = "bluebell"

func CryptoMd5(str string) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s)
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
