package helper

import (
	"github.com/renyusuke/brunnen/lan/cn"
	"github.com/renyusuke/brunnen/vars"
)

var (
	WalletTypeMap = map[int64]map[int64]string{
		vars.CN: cn.WalletTypeMap,
	}
	WalletStatusMap = map[int64]map[int64]string{
		vars.CN: cn.WalletStatusMap,
	}
)

func init() {
}

// GetWalletTypeDesc 拿錢包
func GetWalletTypeDesc(WalletType int64, lanCode int64) (desc string) {
	return WalletTypeMap[lanCode][WalletType]
}

// GetWalletStatusDesc 拿錢包狀態
func GetWalletStatusDesc(status int64, lanCode int64) (desc string) {
	return WalletStatusMap[lanCode][status]
}
