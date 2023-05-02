package helper

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/renyusuke/brunnen/lan/cn"
	"github.com/renyusuke/brunnen/tbl"
	"github.com/renyusuke/brunnen/vars"
	"net/http"
)

var (
	AdminStatusMap = map[int64]map[int64]string{
		vars.CN: cn.AdminStatusMap,
	}
	AdminTypeMap = map[int64]map[int64]string{
		vars.CN: cn.AdminTypeMap,
	}
)

func init() {
}

// GetAdminStatusMap 拿使用者狀態
func GetAdminStatusMap(status int64, lanCode int64) (desc string) {
	return AdminStatusMap[lanCode][status]
}

// GetAdminTypeMap 拿使用者職位
func GetAdminTypeMap(userType int64, lanCode int64) (desc string) {
	return AdminTypeMap[lanCode][userType]
}

type Base struct {
	*tbl.Admin
	IP                string `json:"ip"`
	Path              string `json:"path"`
	Token             string `json:"token"`
	FirstValidatorId  int64  `json:"first_validator_id"`  // 一次验证人ID
	FirstValidator    string `json:"first_validator"`     // 一次验证人账号
	SecondValidatorId int64  `json:"second_validator_id"` // 二次验证人ID
	SecondValidator   string `json:"second_validator"`    // 二次验证人账号
	Validator         string `json:"validator"`           // 验证人
}

func GetBase(ctx *gin.Context) (*Base, error) {

	var out Base
	base := ctx.Value("base")
	err := json.Unmarshal([]byte(fmt.Sprintf("%v", base)), &out)
	if err != nil {
		return nil, err
	}

	if out.FirstValidator != "" {
		out.Validator = out.FirstValidator
	}
	if out.SecondValidator != "" {
		if out.Validator == "" {
			out.Validator = out.SecondValidator
		} else {
			out.Validator += "/" + out.SecondValidator
		}
	}

	return &out, nil
}

func GetBaseByRequest(r *http.Request) (a *Base, err error) {

	base := r.Header.Get("base")
	err = json.Unmarshal([]byte(base), a)
	if err != nil {
		return
	}

	return
}
