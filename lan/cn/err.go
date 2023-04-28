package cn

import (
	"brunnen/enums"
	"errors"
)

var (
	ErrAdminTokenExpired = errors.New("token 失效，請重新登入")
	ErrVerifyFail        = errors.New("驗證失敗")
	ErrAccountFreeze     = errors.New("已凍結")

	ErrDuplicatedName = errors.New("名稱重複")
	ErrPassNotEnough  = errors.New("密碼長度不夠")

	ErrMap = map[int64]error{
		enums.ErrAdminTokenExpired: ErrAdminTokenExpired,
		enums.ErrVerifyFail:        ErrVerifyFail,
		enums.ErrAccountFreeze:     ErrAccountFreeze,
		enums.ErrNameDuplicated:    ErrDuplicatedName,
		enums.ErrPassNotEnough:     ErrPassNotEnough,
	}
)
