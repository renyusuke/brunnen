package cn

import (
	"errors"
	"github.com/renyusuke/brunnen/constant/enums"
)

var (
	RequestSuccess       = errors.New("請求成功")
	ErrAdminTokenExpired = errors.New("token 失效，請重新登入")
	ErrVerifyFail        = errors.New("驗證失敗")
	ErrAccountFreeze     = errors.New("已凍結")
	//商戶
	ErrAdminClientNotExist = errors.New("商戶不存在")

	//共用
	ErrDuplicatedName    = errors.New("名稱重複")
	ErrSystemError       = errors.New("系統錯誤")
	ErrAdminNameNotExist = errors.New("管理者名稱不存在")
	ErrPassNotEnough     = errors.New("密碼長度不夠")
	ErrMissingLanParam   = errors.New("缺少Lan 參數")
	ErrSignInvalid       = errors.New("Sign 錯誤")
	ErrMissingParam      = errors.New("缺少參數")
	/*錢包*/
	ErrClientWalletAlreadyExist = errors.New("商戶已有此錢包")

	ErrSystemBusy = errors.New("系統繁忙")
	ErrMap        = map[int64]error{
		enums.RequestSuccess:       RequestSuccess,
		enums.ErrAdminTokenExpired: ErrAdminTokenExpired,
		enums.ErrVerifyFail:        ErrVerifyFail,
		enums.ErrAccountFreeze:     ErrAccountFreeze,
		enums.ErrNameDuplicated:    ErrDuplicatedName,
		enums.ErrPassNotEnough:     ErrPassNotEnough,
		//商戶
		enums.ErrAdminClientNotExist: ErrAdminClientNotExist,
		/*錢包*/
		enums.ErrClientWalletAlreadyExist: ErrClientWalletAlreadyExist,
		enums.ErrSignInvalid:              ErrSignInvalid,
		/*系統*/
		enums.ErrSystemError:       ErrSystemError,
		enums.ErrAdminNameNotExist: ErrAdminNameNotExist,
		enums.ErrSystemBusy:        ErrSystemBusy,
		enums.ErrMissingLanParam:   ErrMissingLanParam,
		enums.ErrMissingParam:      ErrMissingParam,
	}
)
