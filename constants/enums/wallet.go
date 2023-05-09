package enums

const (
	TradeDefault    = iota // 0-无效订单
	TradeDeposit           // 1-充值-加钱
	TradeWithdrawal        // 2-取款-扣钱
	TradeBet               // 3-下注-扣钱
	TradeUnBet             // 4-撤单-加钱
	TradeWin               // 5-派奖-加钱
	TradeUnWin             // 6-派奖撤销-加钱
	TradeActivity          // 7-活动-加钱
	TradeDividends         // 8-分红-加钱
	TradeMistake           // 9-错误款项修正-扣款
	TradeDeficit           // 10-坏账修正加钱
	TradeManageAdd         // 11-后台加钱-加钱
	TradeManageSub         // 12-后台扣钱-加钱
)
const (
	OperateDefault = iota // 0-无限制
	OperateAdd            // 1-加钱
	OperateSub            // 2-扣钱
)
