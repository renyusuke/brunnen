package tbl

type WithdrawOrder struct {
	Model
	OrderNo         string `gorm:"type:varchar(60)"` // 我方訂單編號
	MerchantNo      string `gorm:"type:varchar(60)"` // 平台訂單編號
	Type            int64  `gorm:"type:int(12)"`     // 訂單模式
	Status          int64  `gorm:"type:int(12)"`     // 訂單裝太
	Sign            string `gorm:"type:varchar(60)"` // 加密sign
	Shift           string `gorm:"type:varchar(40)"` //更次
	ClientCreatedAt int64  `gorm:"type:int(12)"`     //平台創建時間
	CreatedAt       int64  `gorm:"type:int(12)"`     //創建時間
	UpdatedAt       int64  `gorm:"type:int(12)"`     //更新時間
}
