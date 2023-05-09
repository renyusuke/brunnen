package enums

const (
	WalletDeposit  = 1 // 2-存款
	WalletWithdraw = 2 // 3-提款
	WalletFreeze   = 3 // 3-冻结
	WalletUnFreeze = 4 // 4-解冻
)
const (
	WalletDepositSubCashIn          = 1  // 账户存款
	WalletDepositSubTransferIn      = 2  // 转账入款
	WalletDepositSubChangeIn        = 3  // 货币兑入
	WalletDepositSubReturnCode      = 4  // 回码存入
	WalletDepositSubConsume         = 5  // 消费退款
	WalletDepositSubGame            = 6  // 游戏下分
	WalletDepositSubCommission      = 7  // 佣金存卡
	WalletDepositSubStopWork        = 8  // 收工存入
	WalletDepositSubOperate         = 9  // 营运存入
	WalletDepositSubIncome          = 10 // 占成收益
	WalletWithdrawSubCashOut        = 11 // 账户取款
	WalletWithdrawSubAddWinnings    = 12 // 取款加彩
	WalletWithdrawSubTransferOut    = 13 // 转账出款
	WalletWithdrawSubChangeOut      = 14 // 货币兑出
	WalletWithdrawSubStartWork      = 15 // 开工取款
	WalletWithdrawSubGame           = 16 // 游戏上分
	WalletWithdrawSubConsume        = 17 // 消费结算
	WalletWithdrawSubIncome         = 18 // 占成找数
	WalletWithdrawSubOperate        = 19 // 营运支出
	WalletWithdrawSubCommission     = 20 // 占成佣金找数
	WalletWithdrawSubCommissionCash = 21 // 佣金现金
	WalletWithdrawSubMarker         = 22 // 归还Marker
	WalletFreezeSubBail             = 23 // 保证金冻结
	WalletFreezeSubManual           = 24 // 手动冻结
	WalletFreezeSubConsume          = 25 // 消费冻结
	WalletUnFreezeSubBail           = 26 // 保证金解冻
	WalletUnFreezeSubManual         = 27 // 手动解冻
	WalletUnFreezeSubConsume        = 28 // 消费解冻
)

// 1 加钱 2 扣钱 3 冻结 4 解冻
const (
	WalletStateOne   = 1
	WalletStateTwo   = 2
	WalletStateThree = 3
	WalletStateFour  = 4
)

// 钱包类型 1=客户钱包  2=公司钱包  3=内部卡
const (
	WalletTypeOne  = 1
	WalletTypTwo   = 2
	WalletTypThree = 3
)

// 商户类型 (先写死 1=账房服务 2=电投服务)
const (
	MerchantTypeOne = 1
	MerchantTypTwo  = 2
)
