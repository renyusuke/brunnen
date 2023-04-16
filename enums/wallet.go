package enums

// 組 分類 職位是到 UserType 去區分
const (
	WalletTypeCNY  = 1 //人民幣錢包
	WalletTypeUSDT = 2 //USDT錢包
)

// 組狀態
const (
	WalletStatusNormal = 1 //一般
	WalletStatusFreeze = 2 //凍結
	WalletStatusLock   = 3 //封鎖
)

const (
	WalletCurrencyCNY  = 1
	WalletCurrencyUSDT = 2
)
