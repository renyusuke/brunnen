package helper

import (
	"brunnen/lan/cn"
	"brunnen/vars"
)

var (
	ErrMap map[int64]map[int64]error
)

func init() {
	ErrMap[vars.CN] = cn.ErrMap
}
func GetErr(errorCode int64, landConde int64) (errDesc error) {
	language := defaultLan
	return ErrMap[language][errorCode]

}
