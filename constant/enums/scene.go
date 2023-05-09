package enums

// 偷吃户口类型，1 本户口 1 上线 2 顶代 3 其他
const (
	ShareTypeOne   = 1 //本户口
	ShareTypeTwo   = 2 //上线
	ShareTypeThree = 3 //顶代
	ShareTypeFour  = 4 //其他
)

// 占成状态 1 未结算 2 己结算
const (
	ShareStatusOne = 1
	ShareStatusTwo = 2
)

// 开工类型 1: 普通, 2: 营运
const (
	StartWorkTypeOne = 1 //普通
	StartWorkTypeTwo = 2 //营运
)

// 入场类型, 1: 现场, 2: 电投
const (
	AdmissionTypeOne = 1 //现场
	AdmissionTypeTwo = 2 //电投
)

// 出码类型 1. A, 2. B
const (
	TipsTypeOne = 1
	TipsTypeTwo = 2
)

// 本金类型 1. C, 2. M
const (
	PrincipalTypeOne = 1
	PrincipalTypeTwo = 2
)

// 1 暂存 2 取暂存
const (
	TempTypeOne = 1
	TempTypeTwo = 2
)

// 场面资金类型: 1 开工 2 收工 3 回码 4 加彩 5 公水 6 荷水
const (
	SceneTypeOne   = 1
	SceneTypeTwo   = 2
	SceneTypeThree = 3
	SceneTypeFour  = 4
	SceneTypeFive  = 5
	SceneTypeSix   = 6
)

// 转码类型: 1 收工 2 转码 3 回码 4 加彩 5 公水 6 暂存 7 取暂存
const (
	ConvertTypeOne   = 1
	ConvertTypeTwo   = 2
	ConvertTypeThree = 3
	ConvertTypeFour  = 4
	ConvertTypeFive  = 5
	ConvertTypeSix   = 6
	ConvertTypeSeven = 7
)

// 客户类型 1 主客 2 分客
const (
	ClientTypeOne = 1
	ClientTypeTwo = 2
)

// 入场身份 1: 代理, 2: 玩家
const (
	SceneIdentityOne = 1
	SceneIdentityTwo = 2
)

// 加彩类型 1 C码 2 M码
const (
	AddStateOne = 1
	AddStateTwo = 2
)

// 状态 1: 未结算 2: 日结
const (
	CircleStateOne = 1
	CircleStateTwo = 2
)

// 外借银投类型 1 买码 2 营运押金
const (
	LendLedgerTypeOne = 1
	LendLedgerTypeTwo = 2
)

// 外借银投状态 1 加 2 减
const (
	LendLedgerOne = 1
	LendLedgerTwo = 2
)
