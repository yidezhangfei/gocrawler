package util

import (
	"crypto/md5"
	"fmt"
)

func Md5String(str string) string {
	data := []byte(str)
	sum := md5.Sum(data)
	md5str := fmt.Sprintf("%x", sum)
	return md5str
}
