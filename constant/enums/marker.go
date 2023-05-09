package enums

const (
	MarkerTypeOne   = 1 // 股东
	MarkerTypeTwo   = 2 // 上线
	MarkerTypeThree = 3 // 公司
	MarkerTypeFour  = 4 // 临时
	MarkerTypeFive  = 5 // 股本
)

// 操作类型  1.批额 2.签M 3还M  4还息 5下批 6受批 7回收 8归还 9上分 10利息清除 11加彩 12逾期
const (
	MarkerChangeApproved        int = iota + 1 // 批额
	MarkerChangeSignedMarker                   // 签M
	MarkerChangeBackMarker                     // 还M
	MarkerChangeInterest                       // 还息
	MarkerChangeUnderApproved                  // 下批
	MarkerChangeReceiveApproved                // 受批
	MarkerChangeRecycle                        // 回收
	MarkerChangeReturn                         // 归还
	MarkerChangeTopUp                          // 上分
	MarkerChangeClearInterest                  // 利息清除
	MarkerChangeAddBonus                       // 加彩
	MarkerChangeOverTime                       // 逾期
)

const (
	RepayIsExpiredOne = 1 // 未过期
	RepayIsExpiredTwo = 2 // 过期
)

const (
	RepayMarkerStateOne = 1 // 未还款
	RepayMarkerStateTwo = 2 // 已还款
)

const (
	RepayInterestStateOne   = 1 // 未还款
	RepayInterestStateTwo   = 2 // 免除利息
	RepayInterestStateThree = 3 // 已还款
)

const (
	OperationTypeOne = 1 // 签m
	OperationTypeTwo = 2 // 还m
)

const (
	ApprovalStateOne = 1 // 未回收
	ApprovalStateTwo = 2 // 已回收
)
