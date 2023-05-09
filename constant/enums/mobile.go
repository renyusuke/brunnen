package enums

// Vip status
const (
	VipUnknownStatus int64 = iota //未知
	VipIdleStatus                 //空闲
	VipInuseStatus                //在用
	VipBusyStatus                 //忙
	VipInvalidStatus              //无效
	VipOfflineStatus              //离线
	VipRingStatus                 //振铃
	VipRingingStatus              //回铃
	VipHoldStatus                 //保留中
)

// 绑定类型，1-验证密码，2-验证授权人密码
const (
	BindTypePwd  = 1 // 1-验证密码
	BindTypeAuth = 2 // 2-验证授权人密码
)
const (
	SeatBind   = 1 // 电话坐席绑定
	SeatUnBind = 2 // 电话坐席解绑
)

const (
	PhoneReserve    = 1 // 预留电话
	PhoneNotReserve = 2 // 非预留电话
)

const (
	SendMsgAll  = 1 // 发送消息全部联系方式
	SendMsgPart = 2 // 发送消息部分联系方式
)
const (
	DialStatusClick    = 1 // 	点击拨号
	DialStatusControl  = 2 // 	手动拨号
	DialStatusCall     = 3 // 	来电
	DialStatusCallBack = 4 // 	回拨
	DialStatusCallOut  = 5 // 	外呼/呼出
	DialStatusBehalf   = 6 // 	系统代拨来电
	DialStatusRelay    = 7 // 	中继外拨回叫，仅计费用
)
