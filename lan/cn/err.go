package cn

import (
	"errors"
	enums "github.com/renyusuke/brunnen/enums"
)

var (
	RequestSuccess       = errors.New("請求成功")
	ErrAdminTokenExpired = errors.New("token 失效，請重新登入")
	ErrVerifyFail        = errors.New("驗證失敗")
	ErrAccountFreeze     = errors.New("已凍結")

	ErrDuplicatedName    = errors.New("名稱重複")
	ErrSystemError       = errors.New("系統錯誤")
	ErrAdminNameNotExist = errors.New("管理者名稱不存在")
	ErrPassNotEnough     = errors.New("密碼長度不夠")

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

		/*錢包*/
		enums.ErrClientWalletAlreadyExist: ErrClientWalletAlreadyExist,

		/*系統*/
		enums.ErrSystemError:       ErrSystemError,
		enums.ErrAdminNameNotExist: ErrAdminNameNotExist,
		enums.ErrSystemBusy:        ErrSystemBusy,
	}
)
