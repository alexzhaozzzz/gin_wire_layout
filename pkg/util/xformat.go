package util

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

// MD5WithSalt ...
func MD5WithSalt(toSign, salt string) string {
	h := md5.New()
	h.Write([]byte(toSign))
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum(nil))
}

// MD5 ...
func MD5(toSign string) string {
	return MD5WithSalt(toSign, "")
}

// Sha1WithSalt ...
func Sha1WithSalt(data, salt string) []byte {
	h := sha1.New()
	h.Write([]byte(data))
	h.Write([]byte(salt))
	return h.Sum(nil)
}

// Sha1 ...
func Sha1(data string) []byte {
	return Sha1WithSalt(data, "")
}

// Sha256WithSalt ...
func Sha256WithSalt(data, salt string) []byte {
	h := sha256.New()
	h.Write([]byte(data))
	h.Write([]byte(salt))
	return h.Sum(nil)
}

// Sha256 ...
func Sha256(data string) []byte {
	return Sha256WithSalt(data, "")
}

// RandomString 生成一个指定长度的随机字符串
func RandomString(length int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	return string(b)
}
