package cn

import "brunnen/enums"

var (
	TypeGroupAdmin      = "最高權限組"
	TypeGroupOperator   = "操作組"
	TypeGroupAccountant = "結算組"
	TypeGroupSecurity   = "風空組"
	TypeGroupCustomer   = "客服組"
	GroupTypeMap        = map[int64]string{
		1: TypeGroupAdmin,
		2: TypeGroupOperator,
		3: TypeGroupAccountant,
		4: TypeGroupSecurity,
		5: TypeGroupCustomer,
	}

	TypeStatusNormal = "一般"
	TypeStatusFreeze = "凍結"
	TypeStatusLock   = "封鎖"
	GroupStatusMap   = map[int64]string{
		enums.StatusGroupNormal: TypeStatusNormal,
		2:                       TypeStatusFreeze,
		3:                       TypeStatusLock,
	}
)
