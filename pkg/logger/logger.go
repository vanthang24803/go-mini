package logger

import (
	"os"
	"path/filepath"
	"time"

	"github.com/vanthang24803/mini/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.Logger

func Init() {
	cfg := config.New()

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var core zapcore.Core
	if cfg.Logger.Production {
		_ = os.Mkdir("logs", 0755)
		logFile := &lumberjack.Logger{
			Filename:   filepath.Join("logs", "logs_"+time.Now().Format("2006-01-02")+".log"),
			MaxAge:     cfg.Logger.MaxAge,
			MaxBackups: 30,
			MaxSize:    100,
			Compress:   true,
		}

		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(logFile),
			zap.NewAtomicLevelAt(zapcore.InfoLevel),
		)
	} else {
		// In development, only write to console
		consoleWriter := zapcore.Lock(os.Stdout)
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			consoleWriter,
			zap.NewAtomicLevelAt(zapcore.DebugLevel),
		)
	}

	log = zap.New(core, zap.AddCaller())
}

func GetLogger() *zap.Logger {
	return log
}
