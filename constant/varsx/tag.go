package varsx

var (
	//入場類型
	UseFor = map[int64]string{
		1: "现场",
		2: "电投",
	}
	//出碼類型 '"A"=无占成 "B"=有占成 "AO"=A(台面) "AU"=A(台底) "BO"=B(台面) "BU"=B(台底)',
	UseType = map[int64]string{

		1: "A",
		2: "B",
		3: "AO",
		4: "AU",
		5: "BO",
		6: "BU",
	}
	//开工类型
	Mode = map[int64]string{
		1: "普通",
		2: "营运",
	}
	//本金类型
	CapitalType = map[int64]string{

		1: "C",
		2: "M",
	}
	//交易方式 1 现金 2 支票 3 本票 4 Credit 5 电汇 8 转码 0未知
	TradeKind = map[int64]string{
		1: "现金",
		2: "支票",
		3: "本票",
		4: "Credit",
		5: "电汇",
	}
	//操作类型 1=买(现金码没有买操作) 2=转 3=退
	CasinoOperationKind = map[int64]string{
		1: "买",
		2: "转",
		3: "退",
	}
)
