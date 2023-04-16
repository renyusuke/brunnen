package cn

import "brunnen/enums"

var (
	CurrencyPHP  = "菲律賓披索"
	CurrencyHKD  = "港幣"
	CurrencyKRW  = "韓幣"
	CurrencyUSDT = "USDT"
	CurrencyCNY  = "人民幣"
	CurrencyUSD  = "美金"
	CurrencyMap  = map[int64]string{
		enums.CurrencyPHP:  CurrencyPHP,
		enums.CurrencyHKD:  CurrencyHKD,
		enums.CurrencyKRW:  CurrencyKRW,
		enums.CurrencyUSDT: CurrencyUSDT,
		enums.CurrencyCNY:  CurrencyCNY,
		enums.CurrencyUSD:  CurrencyUSD,
	}
)
