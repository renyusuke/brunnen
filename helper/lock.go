package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"time"

	"context"
)

var (
	key        = "public:lock:%v:"
	expiration = 10 * time.Second
)

func SetLock(ctx *gin.Context, r *redis.Client, val interface{}) (res bool, err error) {

	k := fmt.Sprintf(key, val)
	err = r.SetNX(ctx, key, 0, expiration).Err()
	if err != nil {
		return false, err
	}

	go func(k string) {
		defer r.Del(context.Background(), k).Err()
		for {
			select {
			case <-ctx.Done():
				log.Println("上下文執行結束 監聽結束")
				return

			case <-time.After(expiration):
				log.Println("鎖到期")
				return
			}
		}
	}(k)

	return true, err
}
