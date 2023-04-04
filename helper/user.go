package helper

import (
	"brunnen/enums"
	"brunnen/lan/cn"
	"brunnen/vars"
)

var (
	defaultLan    = vars.CN
	defaultStatus = cn.Operator
)

func GetAdminStatusDesc(status int64, lanCode int64) (desc string) {
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
