package initialize

import (
	"common/global"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var logSlice = []string{"err", "sql", "info"}

func InitLogger() {
	encoder := getEncoder()

	for _, s := range logSlice {
		fileName := fmt.Sprintf("./logs/%s/", s)

		writeSyncer := getLogWriter(fileName)
		core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

		loggers := zap.New(core, zap.AddCaller())
		global.Logger[s] = loggers.Sugar()
	}

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filename string) zapcore.WriteSyncer {

	l, _ := rotatelogs.New(
		filename+"%Y%m%d.log",
		rotatelogs.WithMaxAge(30*24*time.Hour),    // 最长保存30天
		rotatelogs.WithRotationTime(time.Hour*24), // 24小时切割一次
	)

	return zapcore.AddSync(l)
}
