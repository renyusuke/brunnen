package topics

const (
	TopicNatsWorkStop = "topic:nats:work:stop" // 收工
)

var (
	//AdminTokenKey 登入 token
	AdminTokenKey = "oa:admin:token:%v"
	//AdminRouteKey 接口過濾
	AdminRouteKey = "oa:admin:route:%v" //v 是 path
	//ClientWallet 錢包分布式鎖
	ClientWallet = "ct:wallet:%v:client:%v"

	//ClientMerchantNo 的 Prefix
	ClientMerchantNo = "cl"
)
