package utils

import (
	"MSC2021/src/global"
	prostates "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := prostates.New(
		path.Join(global.CONFIG.Zap.Director, "%Y-%m-%d.log"),
		prostates.WithLinkName(global.CONFIG.Zap.LinkName),
		prostates.WithMaxAge(7*24*time.Hour),
		prostates.WithRotationTime(24*time.Hour),
	)
	if global.CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
