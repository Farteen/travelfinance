package util

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5String(str string) string {
	md5Item := md5.New()
	io.WriteString(md5Item, str)
	md5str := fmt.Sprintf("%x", md5Item.Sum(nil))
	//md5Str := string(md5Bytes)
	return md5str
}
