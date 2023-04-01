package logs

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var LG *zap.Logger

type LogConfig struct {
	DebugFileName string `json:"debugFileName"`
	InfoFileName  string `json:"infoFileName"`
	WarnFileName  string `json:"warnFileName"`
	MaxSize       int    `json:"maxSize"`
	MaxAge        int    `json:"maxAge"`
	MaxBackups    int    `json:"maxBackups"`
}

func InitLogger(cfg *LogConfig) (err error) {
	writeSyncerDebug := getLogWriter(cfg.DebugFileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	writeSyncerInfo := getLogWriter(cfg.InfoFileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	writeSyncerWarn := getLogWriter(cfg.WarnFileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)

	encoder := getEncoder()

	//文件输出
	debugCore := zapcore.NewCore(encoder, writeSyncerDebug, zapcore.DebugLevel)
	infoCore := zapcore.NewCore(encoder, writeSyncerInfo, zapcore.InfoLevel)
	warnCore := zapcore.NewCore(encoder, writeSyncerWarn, zapcore.WarnLevel)

	//标准输出
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	std := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)
	core := zapcore.NewTee(debugCore, infoCore, warnCore, std)
	LG = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(LG)

	return

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(fileName string, maxSize, maxBackups, maxAge int) zapcore.WriteSyncer {
	logger := &lumberjack.Logger{
		MaxAge:     maxAge,
		MaxSize:    maxSize,
		Filename:   fileName,
		MaxBackups: maxBackups,
	}
	return zapcore.AddSync(logger)
}
