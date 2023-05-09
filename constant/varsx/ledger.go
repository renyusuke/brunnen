package varsx

import "ph-gitlab.vipsroom.net/oa_vip/oa_public/constant/enums"

var (
	ChipsRecordType = map[int64]string{
		enums.CasinoBuy:       "大廠買碼",
		enums.CasinoRolling:   "大廠轉碼",
		enums.CasinoRefund:    "大廠退碼",
		enums.OpenChipsBuy:    "開工買碼",
		enums.OpenChipsAddBuy: "開工加彩",
		enums.ChipsRolling:    "轉碼",
		enums.ChipsRefund:     "回碼",
		enums.PublicTips:      "公水",
		enums.DepositChips:    "暫存",
		enums.WithdrawChips:   "取暫存",
		enums.CloseWork:       "收工退碼",
	}
)
