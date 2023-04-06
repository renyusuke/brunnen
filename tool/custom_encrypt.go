package tool

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptMd5(str string) (encryptRes string) {
	d := []byte(str)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
