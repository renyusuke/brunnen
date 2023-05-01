package helper

import (
	"github.com/renyusuke/brunnen/lan/cn"
	"github.com/renyusuke/brunnen/vars"
)

var (
	GroupTypeMap   map[int64]map[int64]string
	GroupStatusMap map[int64]map[int64]string
)

func init() {

	GroupTypeMap[vars.CN] = cn.GroupTypeMap
	GroupStatusMap[vars.CN] = cn.GroupStatusMap
}

// GetGroupTypeDesc 組 職位
func GetGroupTypeDesc(status int64, lanCode int64) (desc string) {
	return GroupTypeMap[lanCode][status]
}

// GetGroupStatusDesc   拿使用者職位
func GetGroupStatusDesc(userType int64, lanCode int64) (desc string) {
	return GroupStatusMap[lanCode][userType]
}
