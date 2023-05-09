package tbl

type GroupAdmin struct {
	Id        int64 `gorm:"primaryKey;unique;not null;type:int(250)"`
	AdminId   int64 `gorm:"type:int(20)"`
	GroupId   int64 `gorm:"type:int(20)"`
	CreatedAt int64 `gorm:"type:bigint(14)"`
	UpdatedAt int64 `gorm:"type:bigint(14)"`
}
