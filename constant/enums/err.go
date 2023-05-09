package enums

// 1000-1999 系统相关错误
const (
	// 跟随公共包
	ErrSysBadRequest   int64 = 1000 // 错误请求
	ErrSysTokenExpired int64 = 1001 // Token失效
	ErrSysAuthFailed   int64 = 1002 // 权限校验失败
	ErrRequestLimit    int64 = 1003 // 请求频率限制
	ErrTimeout         int64 = 1004 // 请求超时
	ErrIPLimit         int64 = 1005 // IP限制
	ErrImageSizeLimit  int64 = 1006 // 图片大小限制
	ErrImageSuffix     int64 = 1007 // 图片后缀错误

	// 项目内部定义
	ErrMissAcctPwd  int64 = 1100 // 账号或密码错误
	ErrAcctDisabled int64 = 1101 // 账号已停用
)

// 2000-2999 通讯相关错误
const (
	ErrNewsCallBinding             int64 = 2000
	ErrNewsNoSupportedCarrier      int64 = 2001
	ErrNewsNoSupportedMsg          int64 = 2002
	ErrNewsBoundCustomerService    int64 = 2003
	ErrNoMsgService                int64 = 2004 // 没有对应短信服务商
	ErrTheBussAlreadyBound         int64 = 2005 // 该坐席已经绑定,请先解绑
	ErrTheBussCannotBindOther      int64 = 2006 // 不能解绑他人的坐席
	ErrTheServiceMustOffer         int64 = 2007 // 切换为备用线路的服务商必须为未开启状态
	ErrTheServiceStatusDifferent   int64 = 2008 // 数据库和三方服务状态不一致，请重试
	ErrTheServiceRequestFailed     int64 = 2009 // 请求第三方接口失败,置忙置闲失败
	ErrTheAuthFailed               int64 = 2010 // 校验失败
	ErrIllegalAction               int64 = 2011 // 非法动作，请检查您的action
	ErrMainServiceCanNotBeDisabled int64 = 2012 // 主线服务无法禁用
	ErrNoPhoneNumber               int64 = 2013 // 没有电话号码
	ErrSendMessageFailed           int64 = 2014 // 发送消息失败
	ErrCheckBalanceBusyBee         int64 = 2015 // 检测BusyBee失败
	ErrCheckBalanceTexCell         int64 = 2016 // 检测TexCell失败
	ErrCheckBalanceJoey            int64 = 2017 // 检测Joey失败

)

// 场面相关 3000-3999
const ()

// 户口相关 4000-4999
const (
	ErrMemberMistakeOldPwd = 4000
	ErrMemberNotParent     = 4001
	ErrMemberMistakeFlag   = 4002 // 用户Flag错误
)

// 银头相关 5000-5999
const (
	ErrLedgerChipsNameDuplicated   = 5001
	ErrLedgerJunketsChipsNotEnough = 5002
	ErrLedgerChipsNotExist         = 5003
	ErrLedgerHallNotExist          = 5004
	ErrLedgerRegularChipsNotEnough = 5005
	ErrLedgerRoundStillNotClose    = 5006 //還有未收工場次
	ErrLedgerIdsNecessary          = 5007 //綁定交收方案必須至少傳入一筆籌碼id
	ErrLedgerChipsLockFail         = 5008
	ErrLedgerHallLockFail          = 5009
	ErrLedgerDoesNotExist          = 5010 // 银头不存在
)

// 礼物积分服务相关 6000-6999
const ()

// 报表相关 7000-7999
const ()

// 钱包或者Marker相关 8000-8999
const (
	ErrMarkerCreditSubHas        int64 = 8000 // 该子户口存在未归还的上级授信
	ErrWalletQuery               int64 = 8001 // 钱包查询失败
	ErrWalletDoesNotExist        int64 = 8002 // 该钱包不存在
	ErrMarkerLimit               int64 = 8003 // 负额上限不足
	ErrWalletInsufficientBalance int64 = 8004 // 钱包余额不足
)

// 权限校验 9000-9999
const (
	ErrAuthFail int64 = 9001
)
