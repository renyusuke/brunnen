package enums

// ErrAdminTokenExpired Admin 後台相關 0000 - 0999
const (
	RequestSuccess       = 10000
	ErrAdminTokenExpired = 10001 //token 過期
	ErrVerifyFail        = 10002 //驗證錯誤
	ErrAccountFreeze     = 10003 //帳號已凍結
	ErrNameDuplicated    = 20001 //名稱重複
	ErrPassNotEnough     = 20002 //密碼長度不足
	/*40000~ 49999 wallet*/
	ErrClientWalletAlreadyExist = 40001

	ErrSystemError       = 50000 //系統錯誤
	ErrAdminNameNotExist = 50001 //Admin 名稱不存在
	ErrSystemBusy        = 50002 //系統繁忙
)
