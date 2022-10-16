package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

var SugaredLogger *zap.SugaredLogger
var Logger *zap.Logger

func InitLogger() {
	// 初始化 配置写入线程和日期编码器
	log.Println("zapLogger init start")
	writeSyncer := getLogWrite()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	// 配置zaplogger
	Logger = zap.New(core, zap.AddCaller())
	SugaredLogger = Logger.Sugar()
	log.Println("zapLogger init end")
}

func getEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(config)
}

func getLogWrite() zapcore.WriteSyncer {
	lumberjack := &lumberjack.Logger{
		Filename:   "./log.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberjack)
}
