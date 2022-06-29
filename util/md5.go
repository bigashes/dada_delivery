package util

import (
	"crypto/md5"
	"encoding/hex"
)

//返回一个32位md5
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
