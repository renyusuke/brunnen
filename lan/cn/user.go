package cn

import "github.com/renyusuke/brunnen/enums"

// User 相關
var (
	UseTypeAdmin       = "管理員"
	UserTypeManager    = "經理"
	UserTypeOperator   = "操作員"
	UserTypeAccountant = "會計"

	UserTypeMap = map[int64]string{
		enums.UserTypeAdmin:      UseTypeAdmin,
		enums.UserTypeManager:    UserTypeManager,
		enums.UserTypeOperator:   UserTypeOperator,
		enums.UserTypeAccountant: UserTypeAccountant,
	}
	UserStatusFreeze = "凍結"
	UserStatusNormal = "正常使用"
	UserStatusDelete = "已刪除的使用者"
	UserStatusLock   = "被封鎖的使用者"
	UserStatusMap    = map[int64]string{
		enums.UserStatusFreeze: UserStatusFreeze,
		enums.UserStatusNormal: UserStatusNormal,
		enums.UserStatusDelete: UserStatusDelete,
		enums.UserStatusLock:   UserStatusLock,
	}
)
