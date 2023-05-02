package tbl

type RouteBindUser struct {
	Model
	RouteId   int64 `gorm:"type:int(12)"`
	UserId    int64 `gorm:"type:int(12)"`
	CreatedAt int64 `gorm:"type:int(20)"`
	UpdatedAt int64 `gorm:"type:int(20)"`
}
