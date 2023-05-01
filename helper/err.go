package helper

import (
	"github.com/renyusuke/brunnen/lan/cn"
	"github.com/renyusuke/brunnen/vars"
)

var (
	ErrMap = map[int64]map[int64]error{
		vars.CN: cn.ErrMap,
	}
)

func init() {
}
func GetErr(errorCode int64, landConde int64) (errDesc string) {
	language := DefaultLan
	return ErrMap[language][errorCode].Error()

}
