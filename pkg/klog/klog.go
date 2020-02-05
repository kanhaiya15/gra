package klog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// SLogger
var (
	SLogger *zap.SugaredLogger
)

// NewLogger generates a zap configuration object
// The single source of truth for the app configuration.
func NewLogger() *zap.SugaredLogger {
	if SLogger == nil {
		writerSyncer := getLogWriter()
		encoder := getEncoder()
		core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
		zlogger := zap.New(core, zap.AddCaller())
		SLogger = zlogger.Sugar()
	}
	return SLogger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./gra.log",
		MaxSize:    10,
		MaxAge:     15,
		MaxBackups: 15,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
