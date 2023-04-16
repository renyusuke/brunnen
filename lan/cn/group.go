package cn

import "brunnen/enums"

var (
	TypeGroupAdmin      = "最高權限組"
	TypeGroupOperator   = "操作組"
	TypeGroupAccountant = "結算組"
	TypeGroupSecurity   = "風空組"
	TypeGroupCustomer   = "客服組"
	GroupTypeMap        = map[int64]string{
		enums.TypeGroupAdmin:      TypeGroupAdmin,
		enums.TypeGroupOperator:   TypeGroupOperator,
		enums.TypeGroupAccountant: TypeGroupAccountant,
		enums.TypeGroupSecurity:   TypeGroupSecurity,
		enums.TypeGroupCustomer:   TypeGroupCustomer,
	}

	StatusGroupNormal = "一般"
	StatusGroupFreeze = "凍結"
	StatusGroupLock   = "封鎖"
	GroupStatusMap    = map[int64]string{
		enums.StatusGroupNormal: StatusGroupNormal,
		enums.StatusGroupFreeze: StatusGroupFreeze,
		enums.StatusGroupLock:   StatusGroupLock,
	}
)
