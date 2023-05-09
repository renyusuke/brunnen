package enums

// 籌碼交易類型 type 適用於 tbl_ledger_chips_record
const (
	CasinoBuy       = 1  // 大廠買碼
	CasinoRolling   = 2  // 大廠轉碼
	CasinoRefund    = 3  // 大廠退碼
	OpenChipsBuy    = 4  // 開工買碼
	OpenChipsAddBuy = 5  // 開工加彩
	ChipsRolling    = 6  // 轉碼
	ChipsRefund     = 7  // 回碼
	PublicTips      = 8  // 公水
	DepositChips    = 9  // 暫存
	WithdrawChips   = 10 // 取暫存
	CloseWork       = 11 // 收工退碼
)

// trade_kind 交易方式 1 现金 2 支票 3 本票 4 Credit 5 电汇 6 转码
const (
	Cash           = 1
	Cheque         = 2
	PromissoryNote = 3
	Credit         = 4
	WireTransfer   = 5
	Rolling        = 6
)
