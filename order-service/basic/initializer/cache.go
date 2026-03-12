package initializer

import (
	"fmt"
	"gospacex/order-service/basic/config"

	"github.com/redis/go-redis/v9"
)

func RedisInit() {
	redisConfig := config.GlobalConfig.Redis
	Addr := fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)
	config.Rdb = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.Database, // use default DB
	})
	err := config.Rdb.Ping(config.Ctx).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis连接成功")
}
