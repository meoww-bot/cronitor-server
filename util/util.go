package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os/user"
	"time"
)

func Md5encode(s string) string {

	data := []byte(s)
	h := fmt.Sprintf("%x", md5.Sum(data))
	return h
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func ConvStampToDateStr(stamp float64) string {
	tm := time.Unix(int64(stamp), 0)
	return tm.Format("2006-01-02 15:04:05 UTC+8")
}

func GenCode() string {

	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, 6)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)

}

func CheckEnv() string {
	_, err := user.Lookup("bighead")
	if err != nil {
		return "dev"
	}
	return "production"
}
