package helper

import (
	"github.com/renyusuke/brunnen/enums"
	"github.com/renyusuke/brunnen/lan/cn"
	"github.com/renyusuke/brunnen/vars"
)

// GetAdminStatusDesc 拿使用者狀態
func GetAdminStatusDesc(status int64, lanCode int64) (desc string) {
	if lanCode == 0 {
		lanCode = defaultLan
	}
	switch status | lanCode {
	case enums.UserStatusNormal | vars.CN:
		desc = cn.AdminStatusNormal
		break
	case enums.UserStatusFreeze | vars.CN:
		desc = cn.AdminStatusFreeze
		break
	case enums.UserStatusDelete | vars.CN:
		desc = cn.AdminStatusDelete
		break
	case enums.UserStatusLock | vars.CN:
		desc = cn.AdminStatusLock
		break
	default:
		desc = defaultStatus
		break
	}
	return desc
}

// GetUserTypeDesc 拿使用者職位
func GetUserTypeDesc(userType int64, lanCode int64) (desc string) {
	if userType == 0 {
		lanCode = defaultLan
	}
	switch userType | lanCode {
	case enums.UserTypeAdmin | vars.CN:
		desc = cn.AdminTypeAdmin
		break
	case enums.UserTypeManager | vars.CN:
		desc = cn.AdminTypeManager
		break
	case enums.UserTypeOperator | vars.CN:
		desc = cn.AdminTypeOperator
		break
	case enums.UserTypeAccountant | vars.CN:
		desc = cn.AdminTypeAccountant
		break
	default:
		desc = defaultAdminType
		break
	}
	return desc
}
