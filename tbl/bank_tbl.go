package tbl

// Bank 沒使用
type Bank struct {
	Id       int64 `gorm:"primaryKey;unique;not null;type:int(250)"`
	BankName string
}
