package varsx

import "ph-gitlab.vipsroom.net/oa_vip/oa_public/constant/enums"

const (
	InternalCardNameYunYin  = "YUNYIN_CARD"    // 1-运营卡名称
	InternalCardNamGS       = "GS_CARD"        // 2-过数卡名称
	InternalCardNamRate     = "RATE"           // 3-RATE卡名称
	InternalCardNamY        = "Y_CARD"         // 4-Y卡名称
	InternalCardNamFAXI     = "FAXI"           // 5-息卡名称
	InternalCardNamScenChai = "SCEN_CHAIKA"    // 6-场面柴卡名称
	InternalCardNamYongJin  = "YONGJIN_CARD"   // 7-佣金卡名称
	InternalCardNamB        = "B_CARD"         // 8-B佣卡名称
	InternalCardNamY0       = "PROFI0"         // 9-盈0卡名称
	InternalCardNamY1       = "PROFI1"         // 10-盈1卡名称
	InternalCardNamY2       = "PROFI2"         // 11-盈2卡名称
	InternalCardNamY8       = "PROFI8"         // 12-盈8卡名称
	InternalCardNamY9       = "PROFI9"         // 13-盈9卡名称
	InternalCardNamAgentA   = "PROXY_WALLET_A" // 14-代理钱包A卡名称
	InternalCardNamAgentB   = "PROXY_WALLET_B" // 15-代理钱包B卡名称
)

var InternalCardNameMapById = map[int64]string{
	enums.InternalCardIdYunYin:   InternalCardNameYunYin,
	enums.InternalCardIdGS:       InternalCardNamGS,
	enums.InternalCardIdRate:     InternalCardNamRate,
	enums.InternalCardIdY:        InternalCardNamY,
	enums.InternalCardIdFAXI:     InternalCardNamFAXI,
	enums.InternalCardIdScenChai: InternalCardNamScenChai,
	enums.InternalCardIdYongJin:  InternalCardNamYongJin,
	enums.InternalCardIdB:        InternalCardNamB,
	enums.InternalCardIdY0:       InternalCardNamY0,
	enums.InternalCardIdY1:       InternalCardNamY1,
	enums.InternalCardIdY2:       InternalCardNamY2,
	enums.InternalCardIdY8:       InternalCardNamY8,
	enums.InternalCardIdY9:       InternalCardNamY9,
	enums.InternalCardIdAgentA:   InternalCardNamAgentA,
	enums.InternalCardIdAgentB:   InternalCardNamAgentB,
}

var InternalCardIdMapByName = map[string]int64{
	InternalCardNameYunYin:  enums.InternalCardIdYunYin,
	InternalCardNamGS:       enums.InternalCardIdGS,
	InternalCardNamRate:     enums.InternalCardIdRate,
	InternalCardNamY:        enums.InternalCardIdY,
	InternalCardNamFAXI:     enums.InternalCardIdFAXI,
	InternalCardNamScenChai: enums.InternalCardIdScenChai,
	InternalCardNamYongJin:  enums.InternalCardIdYongJin,
	InternalCardNamB:        enums.InternalCardIdB,
	InternalCardNamY0:       enums.InternalCardIdY0,
	InternalCardNamY1:       enums.InternalCardIdY1,
	InternalCardNamY2:       enums.InternalCardIdY2,
	InternalCardNamY8:       enums.InternalCardIdY8,
	InternalCardNamY9:       enums.InternalCardIdY9,
	InternalCardNamAgentA:   enums.InternalCardIdAgentA,
	InternalCardNamAgentB:   enums.InternalCardIdAgentB,
}
