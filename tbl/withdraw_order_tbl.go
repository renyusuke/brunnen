package tbl

type WithdrawOrder struct {
	Model
	OrderNo         string `gorm:"type:varchar(60)"` // 我方訂單編號
	MerchantNo      string `gorm:"type:varchar(60)"` // 平台號
	Type            int64  `gorm:"type:int(12)"`     // 訂單模式
	Status          int64  `gorm:"type:int(12)"`     // 訂單狀態
	FromCurrencyId  int64  `gorm:"type:int(12)"`     // 異動幣種
	ToCurrencyId    int64  `gorm:"type:int(12)"`     // 兌換幣種
	Sign            string `gorm:"type:varchar(60)"` // 加密sign
	Shift           string `gorm:"type:varchar(40)"` //更次
	ClientCreatedAt int64  `gorm:"type:int(12)"`     //平台創建時間
	CreatedAt       int64  `gorm:"type:bigint(20)"`  //創建時間
	UpdatedAt       int64  `gorm:"type:bigint(20)"`  //更新時間
}
