package helper

import (
	"github.com/renyusuke/brunnen/lan/cn"
	"github.com/renyusuke/brunnen/vars"
)

var (
	AdminStatusMap = map[int64]map[int64]string{
		vars.CN: cn.AdminStatusMap,
	}
	//AdminStatusMap map[int64]map[int64]string
	AdminTypeMap map[int64]map[int64]string
)

func init() {
	AdminTypeMap[vars.CN] = cn.AdminTypeMap
}

// GetAdminStatusMap 拿使用者狀態
func GetAdminStatusMap(status int64, lanCode int64) (desc string) {
	return AdminStatusMap[lanCode][status]
}

// GetAdminTypeMap 拿使用者職位
func GetAdminTypeMap(userType int64, lanCode int64) (desc string) {
	return AdminTypeMap[lanCode][userType]
}
