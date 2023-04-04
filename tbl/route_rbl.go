package tbl

type Route struct {
	Model
	RouteName string `gorm:"type:varchar(20)"`
	RoutePath string `gorm:"type:varchar(40)"`
	CreatedAt int64  `gorm:"type:int(12)"`
	UpdatedAt int64  `gorm:"type:int(12)"`
}
