package enums

// 路由 status
const (
	AdminRouteOrder             = 100000 //訂單
	AdminRouteMyWithdrawOrder   = 110000 //我的提款單
	AdminRouteMyDepositOrder    = 120000 //我的存款端
	AdminRouteWithdrawOrder     = 130000 //提款新訂單
	AdminRouteDepositOrder      = 140000 //存款新訂單
	AdminRouteClientUser        = 200000 //admin客戶的使用者
	AdminRouteAdmin             = 300000 //admin的使用者
	AdminRouteSecurity          = 400000 //風控設置
	AdminRouteSecurityAdmin     = 410000 //admin使用者風控
	AdminRouteSecurityUser      = 420000 //商戶使用者風控
	AdminRouteSecurityWhiteList = 430000 //白名單風控
	AdminRouteSecurityBlackList = 440000 //黑名單風控
	AdminRouteSecurityClient    = 450000 //商戶風空
	AdminRouteReport            = 500000 // 報表
	AdminRouteTransactionLog    = 510000 //交易紀錄
	AdminRouteReportWithdraw    = 520000 //提款報表
	AdminRouteReportDeposit     = 530000 //存款報表
	AdminRouteReportTransaction = 540000 //交易報表
	AdminRouteReportKPI         = 550000 //員工績效報表

)
