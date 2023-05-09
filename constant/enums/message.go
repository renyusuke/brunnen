package enums

const (
	SmsServiceMaster = 1 // 1 主线第三方服务
	SmsServiceSlave  = 2 // 2 备线第三方服务
)
const (
	SmsSuccess         = int32(10000)
	SmsSystemErr       = int32(10001)
	SmsServiceNotFound = int32(10002)
	SmsParamErr        = int32(10003)
	SmsFail            = int32(10004)
)
const (
	BusyBeeServiceConst int64 = iota + 1
	TexCellServiceConst
	JoeyServiceConst
	CreditBalanceConst
	SendSMSConst
	GetSendStatusConst
)

const (
	TemplateTypeMarket = 1 // 模版 market 类型
)

const (
	MaxRetry = 3 // 短信发送最大重试次数
)
