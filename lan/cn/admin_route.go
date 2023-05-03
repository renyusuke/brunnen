package cn

import "github.com/renyusuke/brunnen/enums"

var (
	AdminRouteOrder             = "訂單"
	AdminRouteMyWithdrawOrder   = "我的提款單"
	AdminRouteMyDepositOrder    = "我的存款端"
	AdminRouteWithdrawOrder     = "提款新訂單"
	AdminRouteDepositOrder      = "存款新訂單"
	AdminRouteClientUser        = "客戶使用者"
	AdminRouteAdmin             = "admin使用者"
	AdminRouteSecurity          = "風控設置"
	AdminRouteSecurityAdmin     = "admin使用者風控"
	AdminRouteSecurityUser      = "商戶使用者風控"
	AdminRouteSecurityWhiteList = "白名單風控"
	AdminRouteSecurityBlackList = "黑名單風控"
	AdminRouteSecurityClient    = "商戶風空"
	AdminRouteTransactionLog    = "交易紀錄"
	AdminRouteReport            = "報表"
	AdminRouteReportWithdraw    = "提款報表"
	AdminRouteReportDeposit     = "存款報表"
	AdminRouteReportTransaction = "交易報表"
	AdminRouteReportKPI         = "員工績效報表"
	AdminRouteMap               = map[int64]string{
		enums.AdminRouteOrder:             AdminRouteOrder,
		enums.AdminRouteMyWithdrawOrder:   AdminRouteMyWithdrawOrder,
		enums.AdminRouteMyDepositOrder:    AdminRouteMyDepositOrder,
		enums.AdminRouteWithdrawOrder:     AdminRouteWithdrawOrder,
		enums.AdminRouteDepositOrder:      AdminRouteDepositOrder,
		enums.AdminRouteClientUser:        AdminRouteClientUser,
		enums.AdminRouteAdmin:             AdminRouteAdmin,
		enums.AdminRouteSecurity:          AdminRouteSecurity,
		enums.AdminRouteSecurityAdmin:     AdminRouteSecurityAdmin,
		enums.AdminRouteSecurityUser:      AdminRouteSecurityUser,
		enums.AdminRouteSecurityWhiteList: AdminRouteSecurityWhiteList,
		enums.AdminRouteSecurityBlackList: AdminRouteSecurityBlackList,
		enums.AdminRouteSecurityClient:    AdminRouteSecurityClient,
		enums.AdminRouteTransactionLog:    AdminRouteTransactionLog,
		enums.AdminRouteReport:            AdminRouteReport,
		enums.AdminRouteReportWithdraw:    AdminRouteReportWithdraw,
		enums.AdminRouteReportDeposit:     AdminRouteReportDeposit,
		enums.AdminRouteReportTransaction: AdminRouteReportTransaction,
		enums.AdminRouteReportKPI:         AdminRouteReportKPI,
	}
)
