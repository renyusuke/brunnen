package tbl

type RouteUser struct {
	Model
	RouteId   int64 `gorm:"type:int(12)"`
	UserId    int64 `gorm:"type:int(12)"`
	CreatedAt int64 `gorm:"type:int(12)"`
	UpdatedAt int64 `gorm:"type:int(12)"`
}
