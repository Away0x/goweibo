package core

import (
	"fmt"
  "goweibo/core/pkg/timeutils"
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
	core := zapcore.NewCore(encoder, writeSyncer, getLevel())
	logger := zap.New(core, zap.AddCaller())
	appLog = logger.Sugar()

	fmt.Printf("\nLogger initialization successful: in %s, level is %s\n", GetConfig().String("LOG.FOLDER"), getLevel())
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

  prefix += "(" + string(GetConfig().AppRunMode()) + ")"

	timeStr := timeutils.FormatDate(time.Now())
	filename := path.Join(GetConfig().String("LOG.FOLDER"), prefix+timeStr+".log")
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    GetConfig().Int("LOG.MAXSIZE"),
		MaxBackups: GetConfig().Int("LOG.MAXBACKUPS"),
		MaxAge:     GetConfig().Int("LOG.MAXAGES"),
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getLevel() zapcore.Level {
	level := GetConfig().String("LOG.LEVEL")
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
