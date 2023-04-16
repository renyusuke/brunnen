package cn

import "brunnen/enums"

var (
	WalletTypeCNY  = "人民幣錢包"
	WalletTypeUSDT = "USDT錢包"
	WalletTypeMap  = map[int64]string{
		enums.WalletTypeCNY:  WalletTypeCNY,
		enums.WalletTypeUSDT: WalletTypeUSDT,
	}

	WalletStatusNormal = "一般"
	WalletStatusFreeze = "凍結"
	WalletStatusLock   = "封鎖"
	WalletStatusMap    = map[int64]string{
		enums.WalletStatusNormal: WalletStatusNormal,
		enums.WalletStatusFreeze: WalletStatusFreeze,
		enums.WalletStatusLock:   WalletStatusLock,
	}
)
