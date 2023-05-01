package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"time"

	"context"
)

var (
	key        = "public:lock:%v"
	expiration = 10 * time.Second
)

func SetLock(ctx *gin.Context, r *redis.Client, val interface{}) (res bool, err error) {

	k := fmt.Sprintf(key, val)
	err = r.SetNX(ctx, key, 0, expiration).Err()
	if err != nil {
		return false, err.Error()
	}

	go func(k string) {
		defer r.Del(context.Background(), k).Err()
		for {
			select {
			case <-ctx.Done():
				return

			case <-time.After(expiration):
				return
			}
		}
	}(k)

	return true, err.Error()
}
