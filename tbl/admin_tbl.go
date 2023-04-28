package tbl

// Admin 商戶操作人
type Admin struct {
	Model
	Name         string `gorm:"type:varchar(20)"`
	Type         int64  `gorm:"type:int(12)"`
	Status       int64  `gorm:"type:int(12)"`
	Account      string `gorm:"type:varchar(20)"`
	Pass         string `gorm:"type:varchar(20)"`
	GToken       string `gorm:"type:varchar(50)"`
	SecondSecure int64  `gorm:"type:int(12)"`
	UpdatedAt    int64  `gorm:"autoUpdateTime:milli;comment:'更新時間'"`
	CreatedAt    int64  `gorm:"type:int(14);comment:'創建時間'"`
}
