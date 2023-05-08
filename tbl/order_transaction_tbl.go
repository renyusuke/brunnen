package tbl

// OrderTransaction 代付
type OrderTransaction struct {
	Model
	OrderNo         string  `gorm:"type:varchar(60);DefaultValue('_')"`                                               // 我方訂單編號
	MerchantNo      string  `gorm:"type:varchar(60)"`                                                                 // 平台號
	ClientId        int64   `gorm:"type:int(12)"`                                                                     // 申請商戶
	Type            int64   `gorm:"type:int(12);DefaultValue(0);comment:'訂單模式 (1. 等待派發 2. 自有卡 3. 三方 4. 車隊)'"`         // 訂單模式 (1. 等待派發 2. 自有卡 3. 三方 4. 車隊)
	Status          int64   `gorm:"type:int(12);DefaultValue(0);comment:'訂單狀態 (1. 入單成功 2. 處理中（拉單後） 3. 已完成 4. 退回  )'"` // 訂單狀態 (1. 入單成功 2. 處理中（拉單後） 3. 已完成 4. 退回  )
	ForType         int64   `gorm:"type:int(12);DefaultValue(0);comment:'1. 代收 2. 代付'"`                               //1. 代收 2. 代付
	FromCurrencyId  int64   `gorm:"type:int(12);DefaultValue(0);comment:'申請幣種'"`                                      // 申請幣種
	ToCurrencyId    int64   `gorm:"type:int(12);DefaultValue(0);comment:'兌換幣種'"`                                      // 兌換幣種
	Rate            int64   `gorm:"type:int(12);DefaultValue(0)"`                                                     // 兌換匯率
	AssociationID   int64   `gorm:"type:int(12);DefaultValue(0)"`                                                     // 依照Type 判別 如果是自有卡 就是卡id 如果是 三方 就是 三方ID 以此類推
	Sign            string  `gorm:"type:varchar(60)"`                                                                 // 加密sign
	Amount          float64 `gorm:"type:decimal(12,4);comment:'金額';DefaultValue(0)"`                                  // 金額
	BeforeBalance   float64 `gorm:"type:decimal(12,4);comment:'帳變前餘額';DefaultValue(0)"`                               // 帳變前餘額
	AfterBalance    float64 `gorm:"type:decimal(12,4);comment:'帳變後餘額';DefaultValue(0)"`                               // 帳變後餘額
	BankName        string  `gorm:"type:varchar(30)"`                                                                 // 玩家銀行名稱
	BankAccount     string  `gorm:"type:varchar(50)"`                                                                 // 玩家銀行帳號
	BankArea        string  `gorm:"type:varchar(50)"`                                                                 // 支行
	UserSign        string  `gorm:"type:varchar(50)"`                                                                 // 玩家唯一識別
	OperatorName    string  `gorm:"type:varchar(60);DefaultValue(0)"`                                                 // 經手人
	ClientCreatedAt int64   `gorm:"type:int(12)"`                                                                     // 平台創建時間
	CreatedAt       int64   `gorm:"type:bigint(20)"`                                                                  // 創建時間
	UpdatedAt       int64   `gorm:"type:bigint(20)"`                                                                  // 更新時間
}
