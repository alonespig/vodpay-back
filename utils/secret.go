package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
)

func GenerateSecret() (string, error) {
	b := make([]byte, 16) // 16 bytes
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
