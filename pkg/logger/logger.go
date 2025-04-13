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
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var core zapcore.Core
	if cfg.Logger.Production {
		_ = os.Mkdir("logs", 0755)
		logFile := &lumberjack.Logger{
			Filename:   filepath.Join("logs", "logs_"+time.Now().Format("2006-01-02")+".txt"),
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
		consoleWriter := zapcore.Lock(os.Stdout)
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			consoleWriter,
			zap.NewAtomicLevelAt(zapcore.DebugLevel),
		)
	}

	if cfg.Logger.Production {
		log = zap.New(core,
			zap.AddCaller(),
			zap.AddStacktrace(zapcore.ErrorLevel),
		)
	} else {
		log = zap.New(core,
			zap.AddStacktrace(zapcore.WarnLevel),
		)
	}
}

func GetLogger() *zap.Logger {
	return log
}
