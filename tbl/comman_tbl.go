package tbl

type Model struct {
	Id int64 `gorm:"primaryKey;unique;not null;type:int(20)"`
}
