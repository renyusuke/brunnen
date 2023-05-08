package tbl

// ClientUser 商戶使用者
type ClientUser struct {
	Model
	ClientUserSign string `gorm:"type:varchar(50);comment:'商戶使用者唯一標記'"`
	Name           int64  `gorm:"type:varchar(20);primaryKey;comment:'名稱'"`
	Account        string `gorm:"type:varchar(30);comment:'帳號'"`
	Pass           string `gorm:"type:varchar(30);comment:'密碼'"`
	Status         string `gorm:"type:varchar(30);comment:'狀態'"`
	Type           string `gorm:"type:varchar(30);comment:'類別'"`
	UpdatedAt      int64  `gorm:"autoUpdateTime:milli;comment:'更新時間'"`
	CreatedAt      int64  `gorm:"type:int(20);comment:'創建時間'"`
}
