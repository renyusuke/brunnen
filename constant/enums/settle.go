package enums

const (
	SettleTypeImmediately = 1 // 即出
	SettleTypeMonthly     = 2 // 月结
)
const (
	SettleWayCash = 1 // 现金
	SettleWayCard = 2 // 存卡
)

// 1-未结算佣金 2-已结算佣金 3-未结算偷食 4-已结算偷食  5-异常汇率  6-已冻结佣金
const (
	SettleStatusUnSettleCommission   = 1 // 1-未结算佣金
	SettleStatusSettle               = 2 // 2-已结算佣金
	SettleStatusUnSettleSteal        = 3 // 3-未结算偷食
	SettleStatusSettleSteal          = 4 // 5-已结算偷食
	SettleStatusAbnormalExchangeRate = 5 // 4-异常汇率
	SettleStatusFreeze               = 6 // 6-已冻结佣金
)
