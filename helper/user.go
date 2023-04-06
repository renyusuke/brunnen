package helper

import (
	"brunnen/enums"
	"brunnen/lan/cn"
	"brunnen/vars"
)

// GetAdminStatusDesc 拿使用者狀態
func GetAdminStatusDesc(status int64, lanCode int64) (desc string) {
	if lanCode == 0 {
		lanCode = defaultLan
	}
	switch status | lanCode {
	case enums.AdminStatusNormal | vars.CN:
		desc = cn.AdminStatusNormal
		break
	case enums.AdminStatusFreeze | vars.CN:
		desc = cn.AdminStatusFreeze
		break
	case enums.AdminStatusDelete | vars.CN:
		desc = cn.AdminStatusDelete
		break
	case enums.AdminStatusLock | vars.CN:
		desc = cn.AdminStatusLock
		break
	default:
		desc = defaultStatus
		break
	}
	return desc
}

// GetAdminTypeDesc 拿使用者職位
func GetAdminTypeDesc(userType int64, lanCode int64) (desc string) {
	if userType == 0 {
		lanCode = defaultLan
	}
	switch userType | lanCode {
	case enums.AdminTypeAdmin | vars.CN:
		desc = cn.AdminTypeAdmin
		break
	case enums.AdminTypeManager | vars.CN:
		desc = cn.AdminTypeManager
		break
	case enums.AdminTypeOperator | vars.CN:
		desc = cn.AdminTypeOperator
		break
	case enums.AdminTypeAccountant | vars.CN:
		desc = cn.AdminTypeAccountant
		break
	default:
		desc = defaultAdminType
		break
	}
	return desc
}
