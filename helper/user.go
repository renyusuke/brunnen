package helper

import (
	"github.com/renyusuke/brunnen/lan/cn"
	"github.com/renyusuke/brunnen/vars"
)

var (
	UserStatusMap map[int64]map[int64]string
	UserTypeMap   map[int64]map[int64]string
)

func init() {
	UserStatusMap[vars.CN] = cn.UserStatusMap
	UserTypeMap[vars.CN] = cn.UserTypeMap
}

// GetUserStatusDesc 拿使用者狀態
func GetUserStatusDesc(status int64, lanCode int64) (desc string) {
	return UserStatusMap[lanCode][status]
}

// GetUserTypeDesc 拿使用者職位
func GetUserTypeDesc(userType int64, lanCode int64) (desc string) {
	return UserTypeMap[lanCode][userType]
}
