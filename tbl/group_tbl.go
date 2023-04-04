package tbl

type Group struct {
	Model
	GroupName string `gorm:"type:varchar(20)"`
	Type      int64  `gorm:"type:int(12)"`
	Status    int64  `gorm:"type:int(12)"`
	CreatedAt int64  `gorm:"type:int(12)"`
	UpdatedAt int64  `gorm:"type:int(12)"`
}
