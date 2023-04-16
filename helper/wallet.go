package helper

import (
	"brunnen/enums"
	"brunnen/lan/cn"
	"brunnen/vars"
)

var (
	WalletTypeMap   map[int64]map[int64]string
	WalletStatusMap map[int64]map[int64]string
)

func init() {
	WalletTypeMap[vars.CN] = cn.WalletTypeMap
	WalletStatusMap[vars.CN] = cn.WalletStatusMap
}

// GetWalletStatusDesc 拿使用者狀態
func GetWalletStatusDesc(status int64, lanCode int64) (desc string) {
	if lanCode == 0 {
		lanCode = defaultLan
	}
	switch status | lanCode {
	case enums.WalletStatusNormal | vars.CN:
		desc = cn.WalletStatusNormal
		break
	case enums.WalletStatusFreeze | vars.CN:
		desc = cn.WalletStatusFreeze
		break
	case enums.WalletStatusLock | vars.CN:
		desc = cn.WalletStatusLock
		break
	default:
		desc = defaultWalletStatus
		break
	}
	return desc
}
