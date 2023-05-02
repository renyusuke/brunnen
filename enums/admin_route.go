package enums

// 路由 status
const (
	AdminRouteOrder             = 0  //訂單
	AdminRouteMyWithdrawOrder   = 1  //我的提款單
	AdminRouteMyDepositOrder    = 2  //我的存款端
	AdminRouteWithdrawOrder     = 3  //提款新訂單
	AdminRouteDepositOrder      = 4  //存款新訂單
	AdminRouteClientUser        = 5  //admin客戶的使用者
	AdminRouteAdmin             = 6  //admin的使用者
	AdminRouteSecurity          = 7  //風控設置
	AdminRouteSecurityAdmin     = 8  //admin使用者風控
	AdminRouteSecurityUser      = 9  //商戶使用者風控
	AdminRouteSecurityWhiteList = 10 //白名單風控
	AdminRouteSecurityBlackList = 11 //黑名單風控
	AdminRouteSecurityClient    = 12 //商戶風空
	AdminRouteTransactionLog    = 13 //交易紀錄
	AdminRouteReport            = 14 // 報表
	AdminRouteReportWithdraw    = 15 //提款報表
	AdminRouteReportDeposit     = 16 //存款報表
	AdminRouteReportTransaction = 17 //交易報表
	AdminRouteReportKPI         = 18 //員工績效報表

)
