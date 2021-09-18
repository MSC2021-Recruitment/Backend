package global

import (
	"MSC2021/src/models/config"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DATABASE  *gorm.DB
	TOKENBASE *redis.Client
	CONFIG    config.ServerConfig
	VIPER     *viper.Viper
	LOGGER    *zap.Logger
)
