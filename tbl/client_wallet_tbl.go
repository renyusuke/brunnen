package tbl

// ClientWallet 商戶綁定錢包
type ClientWallet struct {
	Model
	ClientId  int64 `gorm:"type:varchar(20);comment:'客戶ID'"`
	WalletId  int64 `gorm:"type:int(12);comment:'錢包ID'"`
	CreatedAt int64 `gorm:"type:int(14);comment:'創建時間'"`
}
