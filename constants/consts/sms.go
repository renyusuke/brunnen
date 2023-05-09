package consts

const (
	BusyBeeServiceConst int64 = iota + 1
	TexCellServiceConst
	JoeyServiceConst
	CreditBalanceConst
	SendSMSConst
	GetSendStatusConst
)

var SMServiceTypeMap = map[int64]map[int64]string{
	BusyBeeServiceConst: BusyBeeRouterMap,
	TexCellServiceConst: TexCellRouterMap,
	JoeyServiceConst:    JoeyRouterMap,
}

const (
	BusyBeeService string = "BusyBee"
	TexCellService string = "Texcell"
	JoeyService    string = "Joey"
	MobilNumberPH  string = "63"
	CurrencyUSD    string = "USD"
	CurrencyRMB    string = "CNY"
	Dash           string = "-"
)

var BusyBeeRouterMap = map[int64]string{
	CreditBalanceConst: "/api/v2/Balance",
	SendSMSConst:       "/api/v2/SendSMS",
	GetSendStatusConst: "/api/v2/MessageStatus",
}

var TexCellRouterMap = map[int64]string{
	CreditBalanceConst: "/1.dox",
	SendSMSConst:       "/14.dox",
}

var JoeyRouterMap = map[int64]string{
	CreditBalanceConst: "/sms/balance/v1",
	SendSMSConst:       "/sms/batch/v2",
}
