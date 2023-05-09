package tbl

type Group struct {
	Id        int64  `gorm:"primaryKey;unique;not null;type:int(250)"`
	GroupName string `gorm:"type:varchar(30)"`
	Path      string `gorm:"type:varchar(30)"`
	Type      int64  `gorm:"type:int(12)"`
	Status    int64  `gorm:"type:int(12)"`
	CreatedAt int64  `gorm:"type:bigint(20)"`
	UpdatedAt int64  `gorm:"type:bigint(12)"`
}
