package services

import (
	"MSC2021/src/global"
	"strconv"
	"time"
)

func GetTokenInWhitelist(userName uint) (redisJWT string, err error) {
	redisJWT, err = global.TOKENBASE.Get(strconv.FormatUint(uint64(userName), 10)).Result()
	return redisJWT, err
}

func PutTokenInWhitelist(userName uint, jwt string) (err error) {
	timer := time.Duration(global.CONFIG.JWT.ExpiresTime) * time.Second
	err = global.TOKENBASE.Set(strconv.FormatUint(uint64(userName), 10), jwt, timer).Err()
	return err
}

func ExpireToken(userName uint) (err error) {
	err = global.TOKENBASE.Del(strconv.FormatUint(uint64(userName), 10)).Err()
	return err
}
