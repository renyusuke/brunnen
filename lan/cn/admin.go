package cn

import "github.com/renyusuke/brunnen/enums"

// User 相關
var (
	AdminTypeAdmin      = "管理員"
	AdminTypeManager    = "經理"
	AdminTypeOperator   = "操作員"
	AdminTypeAccountant = "會計"

	AdminTypeMap = map[int64]string{
		enums.AdminTypeAdmin:      AdminTypeAdmin,
		enums.AdminTypeManager:    AdminTypeManager,
		enums.AdminTypeOperator:   AdminTypeOperator,
		enums.AdminTypeAccountant: AdminTypeAccountant,
	}
	AdminStatusFreeze = "凍結"
	AdminStatusNormal = "正常使用"
	AdminStatusDelete = "已刪除的使用者"
	AdminStatusLock   = "被封鎖的使用者"

	AdminStatusMap = map[int64]string{
		enums.AdminStatusFreeze: AdminStatusFreeze,
		enums.AdminStatusNormal: AdminStatusNormal,
		enums.AdminStatusDelete: AdminStatusDelete,
		enums.AdminStatusLock:   AdminStatusLock,
	}
)
