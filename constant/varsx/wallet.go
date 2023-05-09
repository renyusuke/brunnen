package varsx

import "ph-gitlab.vipsroom.net/oa_vip/oa_public/constant/enums"

// WalletDepositSubMp 账变子类型映射
var WalletDepositSubMp = map[int64]string{
	enums.WalletDepositSubCashIn:          "账户存款",
	enums.WalletDepositSubTransferIn:      "转账入款",
	enums.WalletDepositSubChangeIn:        "货币兑入",
	enums.WalletDepositSubReturnCode:      "回码存入",
	enums.WalletDepositSubConsume:         "消费退款",
	enums.WalletDepositSubGame:            "游戏下分",
	enums.WalletDepositSubCommission:      "佣金存卡",
	enums.WalletDepositSubStopWork:        "收工存入",
	enums.WalletDepositSubOperate:         "营运存入",
	enums.WalletDepositSubIncome:          "占成收益",
	enums.WalletWithdrawSubCashOut:        "账户取款",
	enums.WalletWithdrawSubAddWinnings:    "取款加彩",
	enums.WalletWithdrawSubTransferOut:    "转账出款",
	enums.WalletWithdrawSubChangeOut:      "货币兑出",
	enums.WalletWithdrawSubStartWork:      "开工取款",
	enums.WalletWithdrawSubGame:           "游戏上分",
	enums.WalletWithdrawSubConsume:        "消费结算",
	enums.WalletWithdrawSubIncome:         "占成找数",
	enums.WalletWithdrawSubOperate:        "营运支出",
	enums.WalletWithdrawSubCommission:     "占成佣金找数",
	enums.WalletWithdrawSubCommissionCash: "佣金现金",
	enums.WalletWithdrawSubMarker:         "归还Marker",
	enums.WalletFreezeSubBail:             "保证金冻结",
	enums.WalletFreezeSubManual:           "手动冻结",
	enums.WalletFreezeSubConsume:          "消费冻结",
	enums.WalletUnFreezeSubBail:           "保证金解冻",
	enums.WalletUnFreezeSubManual:         "手动解冻",
	enums.WalletUnFreezeSubConsume:        "消费解冻",
}
