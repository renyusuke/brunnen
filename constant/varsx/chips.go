package varsx

import "ph-gitlab.vipsroom.net/oa_vip/oa_public/constant/enums"

// ChipTypeMp 出码类型映射
var ChipTypeMp = map[int64]string{
	enums.ChipA:     "A",
	enums.ChipB:     "B",
	enums.ChipUpA:   "A(台面)",
	enums.ChipDownA: "A(台底)",
	enums.ChipUpB:   "B(台面)",
	enums.ChipDownB: "B(台底)",
}
