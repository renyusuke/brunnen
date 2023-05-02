package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	expireTim = time.Minute * 2
)

func GetToken(ctx *gin.Context, rdb *redis.Client, key string) (res bool, err error, token string) {
	token, err = rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	if err != nil {
		return false, err, ""
	} else {
		err = SetToken(ctx, rdb, key, token)
		if err != nil {
			return false, err, ""
		}
	}
	return true, err, token
}

func SetToken(ctx *gin.Context, rdb *redis.Client, key string, value any) (err error) {
	err = rdb.Set(ctx, key, value, expireTim).Err()
	return err
}
