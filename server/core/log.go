package core

import (
	"fmt"
	"path"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// SetupLog 初始化日志
func SetupLog() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	appLog = logger.Sugar()
	fmt.Println("zap logger init success")
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	prefix := GetConfig().String("LOG.PREFIX")

	if prefix == "" {
		_ = fmt.Errorf("logger prefix not found")
	}
	// timeStr := time.Now().Format("2006-01-02-13-04-05")
	timeStr := time.Now().Format("2006-01-02")
	filename := path.Join(GetConfig().String("LOG.Folder"), prefix+timeStr+".log")
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    GetConfig().Int("LOG.MAXSIZE"),
		MaxBackups: GetConfig().Int("LOG.MAXBACKUPS"),
		MaxAge:     GetConfig().Int("LOG.MAXAGES"),
	}
	return zapcore.AddSync(lumberJackLogger)
}
