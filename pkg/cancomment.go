package pkg

import (
	"context"
	"fmt"
	"time"
	"weikang/global"
)

var ctx = context.Background()

func CanComment(userID uint64) bool {
	key := fmt.Sprintf("comment:limit:user:%d", userID)
	// 每次评论，计数加1
	count, _ := global.Rds.Incr(ctx, key).Result()
	if count == 1 {
		// 第一次评论，设置过期时间
		global.Rds.Expire(ctx, key, time.Minute)
	}
	// 1分钟最多3条
	return count <= 3
}

func CanCommentByIP(ip string) bool {
	key := fmt.Sprintf("comment:limit:ip:%s", ip)
	count, _ := global.Rds.Incr(ctx, key).Result()
	if count == 1 {
		global.Rds.Expire(ctx, key, time.Minute)
	}
	return count <= 10 // 1分钟同一IP最多10条
}
