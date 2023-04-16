package helper

import (
	"brunnen/lan/cn"
	"brunnen/vars"
)

var (
	CurrencyMap map[int64]map[int64]string
)

func init() {
	CurrencyMap[vars.CN] = cn.CurrencyMap
}

// GetCurrencyDesc 拿錢包
func GetCurrencyDesc(currencyId int64, lanCode int64) (desc string) {
	return CurrencyMap[lanCode][currencyId]
}
