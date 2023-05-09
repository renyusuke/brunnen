package enums

// 订单状态, 1 待结算, 2 已结算,3 待退单, 4 已退单, 5 已取消 6 待审批'
const (
	OrderUnSettlement = 1 // 未结算
	OrderSettled      = 2 // 已结算
	OrderUnRefunded   = 3 // 待退单
	OrderRefunded     = 4 // 已退单
	OrderCancelled    = 5 // 已取消
	OrderPending      = 6 // 待审批
)

// 子订单状态,  1-待结算 2-已结算 3-待退单 4-已退单, 5-已取消
const (
	SubOrderUnSettlement = 1 // 未结算
	SubOrderSettled      = 2 // 已结算
	SubOrderUnRefunded   = 3 // 待退单
	SubOrderRefunded     = 4 // 已退单
	SubOrderCancelled    = 5 // 已取消
)

// 结算方式, 1 现金, 2 存卡, 3 佣金, 4 积分, 5 赠送部门 6 赠送部门额度
const (
	PaymentTypeCash       = 1 // 1 现金
	PaymentTypeCard       = 2 // 2 存卡
	PaymentTypeCommission = 3 // 3 佣金
	PaymentTypeIntegral   = 4 // 4 积分
	PaymentTypeGift       = 5 // 5 赠送部门
	PaymentTypeGiftQuota  = 6 // 6 赠送部门额度
)

// 积分-1 礼遇金-2
const (
	FrontTypeIntegral = 1 // 积分-1
	FrontTypeGift     = 2 // 礼遇金-2
)

// 审批标志 1-需要审批

const (
	IsApproval = 1 // 1-需要审批
)

// 记录类型，1增加 2消费 3退款

const (
	GiftsRecordTypeAdd     = 1 // 1-增加
	GiftsRecordTypeConsume = 2 // 2-消费
	GiftsRecordTypeRefund  = 3 // 3-退款
)

// 预支佣金记录类型
const (
	CommissionRecordTypeConsume = 1 // 1-消费
	CommissionRecordTypeRefund  = 2 // 2-退款
)

// 预支佣金记录状态
const (
	CommissionRecordStatusUnsettled = 1 // 1-未结算
	CommissionRecordStatusSettled   = 2 // 2-已结算
	CommissionRecordStatusCancel    = 3 // 3-取消
)
