package tbl

// RateCurrency 匯率
type RateCurrency struct {
	Model
	ClientId       int64   `gorm:"type:int(12);comment:'商戶ID'"`
	FromCurrencyId int64   `gorm:"type:int(12);comment:'來源幣種'"`
	ToCurrencyId   int64   `gorm:"type:int(12);comment:'目標幣種'"`
	Rate           float64 `gorm:"type:decimal(14,4);comment:'匯率'"`
	UpdatedAt      int64   `gorm:"autoUpdateTime:milli"`
	CreatedAt      int64   `gorm:"type:bigint(20);comment:'創建時間'"`
}
