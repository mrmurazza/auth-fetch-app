package auth

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
)

var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GeneratePassword(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}

	return string(b)
}

func EncryptPassword(password string) string {
	md5pass := md5.Sum([]byte(password))
	sha256pass := sha256.Sum256(md5pass[:])

	str := base64.StdEncoding.EncodeToString(sha256pass[:])
	return str
}
