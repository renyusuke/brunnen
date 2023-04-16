package tbl

type Client struct {
	Model
	Name         int64 `gorm:"type:varchar(20);comment:'名稱'"`
	ClientType   int64 `gorm:"type:int(12);comment:'錢包類別'"`
	ClientStatus int64 `gorm:"type:int(12);comment:'錢包狀態'"`
	UpdatedAt    int64 `gorm:"autoUpdateTime:milli;comment:'更新時間'"`
	CreatedAt    int64 `gorm:"type:int(14);comment:'創建時間'"`
}
