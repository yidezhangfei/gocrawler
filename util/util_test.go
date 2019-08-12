package util

import "testing"

func TestMd5String(t *testing.T) {
	str := "hello world"
	md5str := Md5String(str)
	if md5str == "" {
		t.Fatal("error")
	}
}
