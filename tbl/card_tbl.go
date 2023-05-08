package tbl

type Card struct {
	Model
	BankName string `gorm:"type:varchar(30);comment:'銀行名稱'"`
	BackArea string `gorm:"type:varchar(30);comment:'銀行網點'"`

	OperatorName string `gorm:"type:varchar(60);DefaultValue('_')"` // 經手人
	CreatedAt    int64  `gorm:"type:bigint(20)"`                    // 創建時間
	UpdatedAt    int64  `gorm:"type:bigint(20)"`                    // 更新時間
}
