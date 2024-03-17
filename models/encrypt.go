package models

import (
	"crypto/sha256"
	"encoding/hex"
)

const (
	salt = "密码"
)

// 加密密码
func Encrypt(Password string) string {
	h := sha256.New()
	h.Write([]byte(Password))
	res := hex.EncodeToString(h.Sum([]byte(salt)))
	return res
}
