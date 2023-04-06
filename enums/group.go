package enums

// 組職位
const (
	StatusGroupAdmin     = 1 //admin 後台組
	StatusGroupCarAdmin  = 2 //車隊後台組
	StatusGroupCarDriver = 3 //車手操作組
	StatusGroupAgent     = 4 //代理後台組
)

// 組狀態
const (
	TypeGroupNormal = 1 //一般
	TypeGroupFreeze = 2 //凍結
	TypeGroupLock   = 3 //封鎖
)
