package tbl

// Users 商戶使用者
type Users struct {
	Id           int64  `gorm:"primaryKey;unique;not null;type:int(250)"`
	UserName     string `gorm:"type:varchar(20)"`
	Type         int64  `gorm:"type:int(12)"`
	Status       int64  `gorm:"type:int(12)"`
	Account      string `gorm:"type:varchar(20)"`
	Pass         string `gorm:"type:varchar(40)"`
	GToken       string `gorm:"type:varchar(50)"`
	SecondSecure int64  `gorm:"type:int(12)"`
	CreatedAt    int64  `gorm:"type:bigint(20)"`
	UpdatedAt    int64  `gorm:"type:bigint(12)"`
}
