package enums

// 組 分類 職位是到 UserType 去區分
const (
	TypeGroupAdmin      = 1 //admin 最大權限後台組
	TypeGroupOperator   = 2 //admin 操作組
	TypeGroupAccountant = 3 //admin 結算組
	TypeGroupSecurity   = 4 //admin 風空組
	TypeGroupCustomer   = 5 //客服組
)

// 組狀態
const (
	StatusGroupNormal = 1 //一般
	StatusGroupFreeze = 2 //凍結
	StatusGroupLock   = 3 //封鎖
)
