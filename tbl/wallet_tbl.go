package tbl

// Wallet 商戶的錢包
type Wallet struct {
	Id                  int64   `gorm:"primaryKey;unique;not null;type:int(250)"`
	ClientId            int64   `gorm:"type:int(12);comment:'商戶ID'"`
	CurrencyId          int64   `gorm:"type:int(12);comment:'幣種'"`
	WalletType          int64   `gorm:"type:int(12);comment:'錢包類別'"`
	WalletStatus        int64   `gorm:"type:int(12);comment:'錢包狀態'"`
	Balance             float64 `gorm:"type:decimal(14,4);comment:'錢包餘額'"`
	TempWithdrawBalance float64 `gorm:"type:decimal(14,4);comment:'異動中提款餘額'"`
	TempDepositBalance  float64 `gorm:"type:decimal(14,4);comment:'異動中儲值餘額'"`
	FreezeBalance       float64 `gorm:"type:decimal(14,4);comment:'凍結餘額'"`
	Operator            string  `gorm:"type:varchar(20);comment:'經手人'"`
	SingleDepositCost   float64 `gorm:"type:decimal(14,4);comment:'代收單筆扣費用'"`
	DepositRate         float64 `gorm:"type:decimal(14,4);comment:'代收費率'"`
	SingleWithdrawCost  float64 `gorm:"type:decimal(14,4);comment:'代付單筆扣費用'"`
	WithdrawRate        float64 `gorm:"type:decimal(14,4);comment:'代付費率'"`
	UpdatedAt           int64   `gorm:"autoUpdateTime:milli;comment:'更新時間'"`
	CreatedAt           int64   `gorm:"type:bigint(30);comment:'創建時間'"`
}
