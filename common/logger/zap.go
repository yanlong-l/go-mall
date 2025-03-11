package logger

import (
	"os"

	"github.com/yanlong-l/go-mall/common/enum"
	"github.com/yanlong-l/go-mall/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var _logger *zap.Logger

func ZapLoggerTest(data interface{}) {
	_logger.Info("test for zap init",
		zap.Any("app", config.App),
		zap.Any("database", config.Database),
		zap.Any("data", "快乐池塘栽种了梦想就变成海洋\n鼓的眼睛大嘴巴同样唱的响亮\n借我一双小翅膀就能飞向太阳\n我相信奇迹就在身上\n啦......\n有你相伴 leap frog\n啦......\n自信成长有你相伴 leap frog\n快乐的一只小青蛙 leap frog\n快乐的一只小青蛙 leap frog\n(rap)快乐的池塘里面有只小青蛙\n它跳起舞来就像被王子附体了\n酷酷的眼神,没有哪只青蛙能比美\n总有一天它会被公主唤醒了"),
	)
}

func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	fileWriteSyncer := getFileLogWriter()

	var cores []zapcore.Core
	switch config.App.Env {
	case enum.ModeTest, enum.ModeProd:
		// 测试环境和生产环境的日志输出到文件中
		cores = append(cores, zapcore.NewCore(encoder, fileWriteSyncer, zapcore.InfoLevel))
	case enum.ModeDev:
		// 开发环境同时向控制台和文件输出日志， Debug级别的日志也会被输出
		cores = append(
			cores,
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
			zapcore.NewCore(encoder, fileWriteSyncer, zapcore.DebugLevel),
		)

	}
	core := zapcore.NewTee(cores...)
	_logger = zap.New(core)
}

func getFileLogWriter() (writeSyncer zapcore.WriteSyncer) {
	// 使用 lumberjack 实现 logger rotate
	lumberJackLogger := &lumberjack.Logger{
		Filename:  config.App.Log.FilePath,
		MaxSize:   config.App.Log.FileMaxSize,      // 文件最大 100 M
		MaxAge:    config.App.Log.BackUpFileMaxAge, // 旧文件最多保留90天
		Compress:  false,
		LocalTime: true,
	}

	return zapcore.AddSync(lumberJackLogger)
}
