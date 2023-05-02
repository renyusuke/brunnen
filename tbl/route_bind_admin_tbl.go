package tbl

type RouteBindAdmin struct {
	Model
	RouteId   int64 `gorm:"type:int(12)"`
	AdminId   int64 `gorm:"type:int(12)"`
	CreatedAt int64 `gorm:"type:int(20)"`
	UpdatedAt int64 `gorm:"type:int(20)"`
}
