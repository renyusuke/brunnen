package tbl

type Client struct {
	Model
	Name         int64  `gorm:"type:varchar(20);comment:'名稱'"`
	ClientType   int64  `gorm:"type:int(12);comment:'商戶類別'"`
	ClientStatus int64  `gorm:"type:int(12);comment:'商戶狀態'"`
	CreatorName  string `gorm:"type:varchar(40)"`
	UpdatedAt    int64  `gorm:"autoUpdateTime:milli;comment:'更新時間'"`
	CreatedAt    int64  `gorm:"type:int(14);comment:'創建時間'"`
}
