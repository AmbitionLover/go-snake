package utils

import (
	global "github.com/AmbitionLover/go-snake/pkg"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
	"time"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetWriteSyncer
//@description: zap logger中加入file-rotatelogs
//@return: zapcore.WriteSyncer, error

func GetWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := timeDivisionWriter(file)

	if global.CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}

func sizeDivisionWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename, // 日志文件的位置
		MaxSize:    10,       // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 200,      // 保留旧文件的最大个数
		MaxAge:     30,       // 保留旧文件的最大天数
		Compress:   true,     // 是否压缩/归档旧文件
		LocalTime:  true,
	}
}

func timeDivisionWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(
		strings.Replace(filename, ".log", "", -1)+"_%Y-%m-%d.log",
		rotatelogs.WithMaxAge(time.Duration(int64(24*time.Hour)*int64(30))),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
