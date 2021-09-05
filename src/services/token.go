package services

import (
	"MSC2021/src/global"
	"strconv"
	"time"
)

func GetTokenInWhitelist(userName uint) (redisJWT string, err error) {
	redisJWT, err = global.REDIS.Get(strconv.FormatUint(uint64(userName), 10)).Result()
	return redisJWT, err
}

func PutTokenInWhitelist(userName uint, jwt string) (err error) {
	timer := time.Duration(global.CONFIG.JWT.ExpiresTime) * time.Second
	err = global.REDIS.Set(strconv.FormatUint(uint64(userName), 10), jwt, timer).Err()
	return err
}
