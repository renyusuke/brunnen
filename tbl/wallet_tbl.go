package tbl

// Wallet 商戶的錢包
type Wallet struct {
	Model
	ClientId            int64 `gorm:"type:int(12);comment:'商戶ID'"`
	CurrencyId          int64 `gorm:"type:int(12);comment:'幣種'"`
	WalletType          int64 `gorm:"type:int(12);comment:'錢包類別'"`
	WalletStatus        int64 `gorm:"type:int(12);comment:'錢包狀態'"`
	Balance             int64 `gorm:"type:bigint(20);comment:'錢包餘額'"`
	TempWithdrawBalance int64 `gorm:"type:bigint(20);comment:'異動中提款餘額'"`
	TempDepositBalance  int64 `gorm:"type:bigint(20);comment:'異動中儲值餘額'"`
	FreezeBalance       int64 `gorm:"type:bigint(20);comment:'凍結餘額'"`
	UpdatedAt           int64 `gorm:"autoUpdateTime:milli"`
	CreatedAt           int64 `gorm:"type:bigint(20);comment:'創建時間'"`
}
