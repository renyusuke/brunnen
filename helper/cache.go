package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/renyusuke/brunnen/tool"
	"strconv"
	"time"
)

var (
	expireTim = time.Minute * 2
)

func GetToken(ctx *gin.Context, rdb *redis.Client, key string) (res bool, err error, token string) {
	err = rdb.Get(ctx, key).Err()
	if err != nil {
		return false, err, ""
	} else {
		w := tool.Worker{}
		tokenInt64 := w.GetId()
		token = strconv.FormatInt(tokenInt64, 10)
		err = SetToken(ctx, rdb, key, token)
		if err != nil {
			return false, err, ""
		}
	}
	return true, err, token
}

func SetToken(ctx *gin.Context, rdb *redis.Client, key string, value string) (err error) {
	err = rdb.Set(ctx, key, value, expireTim).Err()
	return err
}
