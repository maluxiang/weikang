package common

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// MD5 密码MD5加密
func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return hex.EncodeToString(has[:])
}

// GenerateRandomString 生成16位随机数字符串
func GenerateRandomString() string {
	rand.Seed(time.Now().UnixNano())
	
	chars := "0123456789"
	randomString := make([]byte, 16)
	
	for i := range randomString {
		randomString[i] = chars[rand.Intn(len(chars))]
	}
	
	return string(randomString)
}
