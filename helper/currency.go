package helper

import (
	"github.com/renyusuke/brunnen/lan/cn"
	"github.com/renyusuke/brunnen/vars"
)

var (
	CurrencyMap = map[int64]map[int64]string{
		vars.CN: cn.CurrencyMap,
	}
)

func init() {
}

// GetCurrencyDesc 拿錢包
func GetCurrencyDesc(currencyId int64, lanCode int64) (desc string) {
	return CurrencyMap[lanCode][currencyId]
}
