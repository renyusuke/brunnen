package tbl

// Admin 商戶操作人
type Admin struct {
	Model
	Name        int64 `gorm:"type:varchar(20);comment:'名稱'"`
	AdminType   int64 `gorm:"type:int(12);comment:'後台使用者類別'"`
	AdminStatus int64 `gorm:"type:int(12);comment:'後台使用者狀態'"`
	UpdatedAt   int64 `gorm:"autoUpdateTime:milli;comment:'更新時間'"`
	CreatedAt   int64 `gorm:"type:int(14);comment:'創建時間'"`
}
