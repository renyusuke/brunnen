package cn

import "errors"

var (
	ErrMap = map[int64]error{
		10001: errors.New("token 失效，請重新登入"),
	}
	ErrAdminTokenExpired = errors.New("token 失效，請重新登入")
)
