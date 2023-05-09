package enums

// 开工类型
const (
	StartWorkNormal  = 1 //-普通
	StartWorkOperate = 2 //-营运
)

// 开工账房状态类型 1: 空白, 2: 在场 默认, 3: 离场/收工
const (
	AccountStatusOne   = 1 //空白
	AccountStatusTwo   = 2 //在场
	AccountStatusThree = 3 //离场/收工
)

// 账房开工后，默认为入场,  场面状态 1: 入场, 2: 在场, 3: 离场
const (
	SceneStatusOne   = 1 //空白
	SceneStatusTwo   = 2 //在场
	SceneStatusThree = 3 //离场/收工
)

// 场次状态 1 开工 2 收工
const (
	RoundStatusOne = 1
	RoundStatusTwo = 2
)

// 是否暂停开工 1 开工 2 暂停开工
const (
	StartWork = 1 //开工
	StopWork  = 2 //暂停开工
)
