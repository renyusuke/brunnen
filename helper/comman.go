package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/renyusuke/brunnen/lan/cn"
	"github.com/renyusuke/brunnen/vars"
)

var (
	DefaultLan          = vars.CN
	DefaultStatus       = cn.AdminStatusNormal
	DefaultAdminType    = cn.AdminTypeOperator
	DefaultWalletStatus = cn.WalletStatusNormal
)

func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}
