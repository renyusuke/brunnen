package tbl

// OrderWithdraw 代付
type OrderWithdraw struct {
	Model
	OrderNo         string  `gorm:"type:varchar(60);DefaultValue('_')"` // 我方訂單編號
	MerchantNo      string  `gorm:"type:varchar(60)"`                   // 平台號
	ClientId        int64   `gorm:"type:int(12)"`                       // 申請商戶
	Type            int64   `gorm:"type:int(12);DefaultValue(0)"`       // 訂單模式 (1. 等待派發 2. 自有卡 3. 三方 4. 車隊)
	Status          int64   `gorm:"type:int(12);DefaultValue(0)"`       // 訂單狀態 (1. 入單成功 2. 處理中（拉單後） 3. 已完成 4. 退回  )
	FromCurrencyId  int64   `gorm:"type:int(12);DefaultValue(0)"`       // 申請幣種
	ToCurrencyId    int64   `gorm:"type:int(12);DefaultValue(0)"`       // 兌換幣種
	Rate            int64   `gorm:"type:int(12);DefaultValue(0)"`       // 兌換匯率
	AssociationID   int64   `gorm:"type:int(12);DefaultValue(0)"`       // 依照Type 判別 如果是自有卡 就是卡id 如果是 三方 就是 三方ID 以此類推
	Sign            string  `gorm:"type:varchar(60)"`                   // 加密sign
	Amount          float64 `gorm:"type:decimal(12,4);"`                // 金額
	BankName        string  `gorm:"type:varchar(30)"`                   // 玩家銀行名稱
	BankAccount     string  `gorm:"type:varchar(50)"`                   // 玩家銀行帳號
	BankArea        string  `gorm:"type:varchar(50)"`                   // 支行
	UserSign        string  `gorm:"type:varchar(50)"`                   // 玩家唯一識別
	OperatorName    string  `gorm:"type:varchar(60);DefaultValue(0)"`   // 經手人
	ClientCreatedAt int64   `gorm:"type:int(12)"`                       // 平台創建時間
	CreatedAt       int64   `gorm:"type:bigint(20)"`                    // 創建時間
	UpdatedAt       int64   `gorm:"type:bigint(20)"`                    // 更新時間
}
