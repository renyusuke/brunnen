package varsx

import "ph-gitlab.vipsroom.net/oa_vip/oa_public/constant/enums"

// MarkerType 信贷类型 1 股东 2 上线 3 公司 4 临时 5 股本
var MarkerType = []int{1, 2, 3, 4, 5}

// MarkerTypeMp M类型映射
var MarkerTypeMp = map[int64]string{
	enums.MarkerTypeOne:   "股东M",
	enums.MarkerTypeTwo:   "上线M",
	enums.MarkerTypeThree: "公司M",
	enums.MarkerTypeFour:  "临时M",
	enums.MarkerTypeFive:  "股本M",
}
