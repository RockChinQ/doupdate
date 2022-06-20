package util

import (
	"crypto/md5"
	"fmt"

	"github.com/google/uuid"
)

func GenerateFingerPrint() string {
	return uuid.New().String()
}

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
