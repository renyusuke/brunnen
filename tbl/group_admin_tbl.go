package tbl

type GroupAdmin struct {
	Model
	AdminId   int64 `gorm:"type:int(20)"`
	GroupId   int64 `gorm:"type:int(20)"`
	CreatedAt int64 `gorm:"type:int(14)"`
	UpdatedAt int64 `gorm:"type:int(14)"`
}