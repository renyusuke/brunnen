package tbl

// Users 商戶使用者
type Users struct {
	Model
	UserName     string `gorm:"type:varchar(20)"`
	Type         int64  `gorm:"type:int(12)"`
	Status       int64  `gorm:"type:int(12)"`
	Account      string `gorm:"type:varchar(20)"`
	Pass         string `gorm:"type:varchar(20)"`
	GToken       string `gorm:"type:varchar(50)"`
	SecondSecure int64  `gorm:"type:int(12)"`
	CreatedAt    int64  `gorm:"type:int(12)"`
	UpdatedAt    int64  `gorm:"type:int(12)"`
}
