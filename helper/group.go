package helper

import (
	"brunnen/enums"
	"brunnen/lan/cn"
	"brunnen/vars"
)

// GetGroupStatusDesc 拿使用者職位
func GetGroupStatusDesc(status int64, lanCode int64) (desc string) {
	if lanCode == 0 {
		lanCode = defaultLan
	}
	switch status | lanCode {
	case enums.StatusAdmin | vars.CN:
		desc = cn.Admin
		break
	case enums.StatusManager | vars.CN:
		desc = cn.Manager
		break
	case enums.StatusAccountant | vars.CN:
		desc = cn.Accountant
		break
	case enums.StatusOperator | vars.CN:
		desc = cn.Operator
		break
	default:
		desc = defaultStatus
		break
	}
	return desc
}

// GetGroupTypeDesc 使用者狀態
func GetGroupTypeDesc(userType int64, lanCode int64) (desc string) {
	if userType == 0 {
		lanCode = defaultLan
	}
	switch userType | lanCode {
	case enums.TypeAdminNormal | vars.CN:
		desc = cn.AdminNormal
		break
	case enums.TypeAdminFreeze | vars.CN:
		desc = cn.AdminFreeze
		break
	case enums.TypeAdminDelete | vars.CN:
		desc = cn.AdminDelete
		break
	case enums.TypeAdminLock | vars.CN:
		desc = cn.AdminLock
		break
	default:
		desc = defaultAdminType
		break
	}
	return desc
}
