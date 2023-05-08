package tbl

type Client struct {
	Model
	Name         string `gorm:"type:varchar(20);comment:'名稱'"`
	ClientType   int64  `gorm:"type:int(12);comment:'商戶類別';comment:'1:一般商戶 2: vip 商戶 3: 垃圾商戶'"`
	ClientStatus int64  `gorm:"type:int(12);comment:'商戶狀態';comment:'1:啟用 2:停用'"`
	CreatorName  string `gorm:"type:varchar(40)"`
	UpdatedAt    int64  `gorm:"autoUpdateTime:milli;comment:'更新時間'"`
	CreatedAt    int64  `gorm:"type:int(20);comment:'創建時間'"`
}
