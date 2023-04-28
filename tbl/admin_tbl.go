package tbl

// Admin 操作人
type Admin struct {
	Model
	Name         string `gorm:"type:varchar(20);comment:'使用者名稱'"`
	Type         int64  `gorm:"type:int(12);comment:'職位'"`
	Status       int64  `gorm:"type:int(12);comment:'狀態'"`
	Account      string `gorm:"type:varchar(20)"`
	Pass         string `gorm:"type:varchar(40)"`
	GToken       string `gorm:"type:varchar(50)"`
	SecondSecure int64  `gorm:"type:int(12)"`
	UpdatedAt    int64  `gorm:"autoUpdateTime:milli;comment:'更新時間'"`
	CreatedAt    int64  `gorm:"type:int(14);comment:'創建時間'"`
}
