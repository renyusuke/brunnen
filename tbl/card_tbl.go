package tbl

type Card struct {
	Id       int64  `gorm:"primaryKey;unique;not null;type:int(250)"`
	BankName string `gorm:"type:varchar(30);comment:'銀行名稱'"`
	BackArea string `gorm:"type:varchar(30);comment:'銀行網點'"`

	Balance   float64 `gorm:"type:decimal(12,4);comment:'金額'"`                  // 金額
	Operator  string  `gorm:"type:varchar(60);DefaultValue('_');comment:'經手人'"` // 經手人
	CreatedAt int64   `gorm:"type:bigint(20);comment:'創建時間'"`                   // 創建時間
	UpdatedAt int64   `gorm:"type:bigint(20);comment:'更新時間'"`                   // 更新時間
}
