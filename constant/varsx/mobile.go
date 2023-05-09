package varsx

import "ph-gitlab.vipsroom.net/oa_vip/oa_public/constant/enums"

const (
	Empty          string = ""
	DnsStatusFree         = 0
	DnsStatusBusy         = 1
	Zero                  = 0
	Two                   = 2
	One                   = 1
	Unbind                = "unbind"
	Bind                  = "bind"
	Dash                  = "-"
	PhAreaCode            = "63"
	WsCdrEventPush        = "WS.cdr.event.%s"
	CtiAudio              = "/ctiaudio"
)

var CdrStatusMap = map[string]string{
	"AW": "接听",
	"BS": "对方忙线，对方返回486",
	"CS": "线路不可用，中继返回501,603之类",
	"CL": "主叫放弃",
	"NA": "振铃未接，无人接听",
	"NT": "外拨无权限",
	"VL": "帐户并发受限",
	"TL": "中继并发受限",
	"BK": "黑名单，来电或外拨都有黑名单限制",
	"EX": "授权到期",
	"NM": "帐户余额不足",
	"FL": "未明确错误",
	"UK": "未知错误",
}

var HallNameIdMap = map[string]int{
	"OKADA":   1,
	"COD":     2,
	"HANNA":   3,
	"RESORTS": 4,
	"SOLAIRE": 5,
}

var DepartmentNameIdMap = map[string]int{
	"市场部": 4,
	"帐房部": 2,
	"礼宾部": 3,
	"电投部": 4,
	"场面部": 5,
}

const (
	Dial       string = "Dial"   //拨号接口
	Hangup     string = "Hangup" //挂断电话
	UnHold     string = "UnHold" //取消保留电话
	Hold       string = "Hold"   //保留电话
	Atxfer     string = "atxfer" //转移电话-秘书总台模式
	Cdr        string = "cdr"    //获取座席的某天的话单记录
	Get        string = "Get"    //获取座席信息
	DndOff     string = "DNDOFF" //置闲
	DndOn      string = "DNDON"  //置忙
	Vcfrun     string = "vcfrun" //密码认证入口
	RecDown    string = "recdown"
	DisConnect string = "DisConnect"
	Event      string = "Event"
)

var VoipRouterMap = map[string]string{
	Dial:    "/cti/Call.php",
	Hangup:  "/cti/Call.php",
	UnHold:  "/cti/Call.php",
	Atxfer:  "/cti/Call.php",
	Cdr:     "/cti/Cdr.php",
	Get:     "/cti/Agent.php",
	DndOff:  "/cti/Agent.php",
	DndOn:   "/cti/Agent.php",
	Vcfrun:  "/cti/Call.php",
	RecDown: "/cti/recdown.php",
	Hold:    "/cti/Call.php",
}

const (
	MaxCdrEvent = 3 // 最大重试次数
)

// MessageMap TODO 对应个 JSON 文档最好
var MessageMap = map[int32]string{
	enums.SmsSuccess:         "请求成功",
	enums.SmsSystemErr:       "系统内部错误",
	enums.SmsServiceNotFound: "短信服务商没有找到",
	enums.SmsParamErr:        "参数错误",
	enums.SmsFail:            "请求失败",
}
var SMServiceTypeMap = map[int64]map[int64]string{
	enums.BusyBeeServiceConst: BusyBeeRouterMap,
	enums.TexCellServiceConst: TexcellRouterMap,
	enums.JoeyServiceConst:    JoeyRouterMap,
}

const (
	BusyBeeService string = "BusyBee"
	TexCellService string = "Texcell"
	JoeyService    string = "Joey"
	CurrencyUSD    string = "USD"
	CurrencyRMB    string = "CNY"
)

// AreaCodePH 区号常量
const (
	AreaCodePH = "63"
)

var BusyBeeRouterMap = map[int64]string{
	enums.CreditBalanceConst: "/api/v2/Balance",
	enums.SendSMSConst:       "/api/v2/SendSMS",
	enums.GetSendStatusConst: "/api/v2/MessageStatus",
}

var TexcellRouterMap = map[int64]string{
	enums.CreditBalanceConst: "/1.dox",
	enums.SendSMSConst:       "/14.dox",
}

var JoeyRouterMap = map[int64]string{
	enums.CreditBalanceConst: "/sms/balance/v1",
	enums.SendSMSConst:       "/sms/batch/v2",
}

const (
	DialStatusCall    = "i"
	DialStatusCallOut = "o"
)
