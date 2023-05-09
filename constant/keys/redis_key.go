package keys

// 后台账号相关

var (
	RedisKeyAdminToken   = "oa:admin:token:%v"   // 后台账号Token
	RedisKeyAdminData    = "oa:admin:data:%v"    // 后台账号缓存数据
	RedisKeyAdminSession = "oa:admin:session:%v" // 后台账号Session
)

// 会员相关

var (
	RedisKeyMemberToken       = "oa:member:token:%v"               // H5会员账号Token
	RedisKeyMemberUnavailable = "oa:member:unavailable"            // 会员账号是否被禁用
	RedisFrozenShareDeposit   = "SCENE:FROZEN:SHARE:DEPOSIT:%s:%d" // redis 占成保证金
)

// 授权相关

var (
	RedisKeySysAuth      = "oa:sys:auth:%v"       // 后台账号授权码Key Value 为后台账号ID
	RedisKeySysAuthPath  = "oa:sys:auth:path:%v"  // 后台路由校验部门
	RedisKeySysAuthData  = "oa:sys:auth:data:%v"  // 二次校验的数据保存Key
	RedisKeySysAuthAdmin = "oa:sys:auth:admin:%v" // 二次校验的校验人信息
)

// 通讯相关

var (
	RedisKeyNewsletterEventMax   = "oa:newsletter:event:%v" // 电话最大重试次数Key
	RedisKeyNewsletterSmsService = "oa:newsletter:sms:%v"   // 短信支持的区号服务商
)

// RedisKeyLedgerShift 更次 需要带场馆ID
var (
	RedisKeyLedgerShift = "oa:sys:shift:%v"
)

var (
	RedisKeyHallMonth = "oa:sys:month:%v"
)

// 消息code生成存储redis key
var (
	RedisKeyInformationCodeTz = "oa:information:tz:code:%v" //通知code

	RedisKeyInformationCodeHd = "oa:information:hd:code:%v" //活动code

	RedisKeyInformationCodeZx = "oa:information:zx:code:%v" //资讯code

)
