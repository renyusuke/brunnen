package tbl

type Client struct {
	Id           int64  `gorm:"primaryKey;unique;not null;type:int(250)"`
	Name         string `gorm:"type:varchar(40);comment:'名稱'"`
	ClientType   int64  `gorm:"type:int(12);comment:'商戶類別1:一般商戶 2: vip 商戶 3: 垃圾商戶'"`
	ClientStatus int64  `gorm:"type:int(12);comment:'商戶狀態1:啟用 2:停用'"`
	MerchantNo   string `gorm:"type:varchar(40);index:merchant_no,unique;comment:'商戶識別碼'"`
	MerchantKey  string `gorm:"type:varchar(40);index:merchant_key,unique;comment:'商戶私鑰'"`
	Operator     string `gorm:"type:varchar(40)"`
	UpdatedAt    int64  `gorm:"autoUpdateTime:milli;comment:'更新時間'"`
	CreatedAt    int64  `gorm:"type:bigint(30);comment:'創建時間'"`
}
