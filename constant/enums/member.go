package enums

const (
	MemberFlagNormal   = 1 // 普通账户
	MemberFlagCompany  = 2 // 公司账户
	MemberFlagInternal = 3 // 内部账号
)

const (
	MemberTypeNormal      = 1 // 普通类型
	MemberTypeShareholder = 2 // 股东类型
)

const (
	MemberQueryByMember      = 1 // 户口表查询
	MemberQueryByPhone       = 2 // 联系方式表查询
	MemberQueryByCertificate = 3 // 证件表表查询
)
const (
	ProposalStatusValid   = 1 // 1-有效订单
	ProposalStatusInvalid = 2 // 2-无效订单
)

// 性别 1.男  2.女
const (
	GenderOne = 1
	GenderTwo = 2
)
