package tbl

type RouteBindAdmin struct {
	Id        int64 `gorm:"primaryKey;unique;not null;type:int(250)"`
	RouteId   int64 `gorm:"type:int(12)"`
	AdminId   int64 `gorm:"type:int(12)"`
	CreatedAt int64 `gorm:"type:bigint(20)"`
	UpdatedAt int64 `gorm:"type:bigint(20)"`
}
