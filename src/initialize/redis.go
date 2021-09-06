package initialize

import (
	"MSC2021/src/global"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"os"
)

func Redis() *redis.Client {
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	result, err := client.Ping().Result()
	global.LOGGER.Info("Trying to connect Redis on " + redisCfg.Addr)
	if err != nil {
		global.LOGGER.Error("Redis connection ping failed, err:", zap.Any("err", err))
		os.Exit(1)
	} else {
		global.LOGGER.Info("Successfully connected to Redis:", zap.String("ping", result))
	}
	return client
}
