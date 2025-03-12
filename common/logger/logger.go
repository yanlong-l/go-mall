package logger

import (
	"context"
	"fmt"
	"path"
	"runtime"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once sync.Once
	l    *Logger
)

type Logger struct {
	_logger *zap.Logger
}

func getLogInstance() *Logger {
	once.Do(func() {
		l = &Logger{
			_logger: _logger,
		}
	})
	return l
}

func (l *Logger) Log(lvl zapcore.Level, msg string, kv ...interface{}) {
	// 保证要打印的日志信息成对出现
	if len(kv)%2 != 0 {
		kv = append(kv, "keysAndValues is not a pair")
	}
	// 在log中携带上调用日志方法的信息
	funcName, file, line := l.getLoggerCallerInfo()
	kv = append(kv, "funcName", funcName, "file", file, "line", line)

	// 处理kv，把它变成[]zap.Field类型
	fields := make([]zap.Field, 0, len(kv)/2)
	for i := 0; i < len(kv); i += 2 {
		fields = append(fields, zap.Any(fmt.Sprint(kv[i]), kv[i+1]))
	}
	// 调用zap的方法打印日志
	if ce := l._logger.Check(lvl, msg); ce != nil {
		ce.Write(fields...)
	}
}

// getLoggerCallerInfo 日志调用者信息 -- 方法名, 文件名, 行号
func (l *Logger) getLoggerCallerInfo() (funcName, file string, line int) {
	pc, file, line, ok := runtime.Caller(3) // 回溯拿调用日志方法的业务函数的信息
	if !ok {
		return
	}
	file = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	return
}

func Debug(ctx context.Context, msg string, kv ...interface{}) {
	kv = append(kv, "traceid", ctx.Value("traceid"), "spanid", ctx.Value("spanid"), "pspanid", ctx.Value("pspanid"))
	l = getLogInstance()
	l.Log(zap.DebugLevel, msg, kv...)
}

func Info(ctx context.Context, msg string, kv ...interface{}) {
	kv = append(kv, "traceid", ctx.Value("traceid"), "spanid", ctx.Value("spanid"), "pspanid", ctx.Value("pspanid"))
	l = getLogInstance()
	l.Log(zap.InfoLevel, msg, kv...)
}

func Warn(ctx context.Context, msg string, kv ...interface{}) {
	kv = append(kv, "traceid", ctx.Value("traceid"), "spanid", ctx.Value("spanid"), "pspanid", ctx.Value("pspanid"))
	l = getLogInstance()
	l.Log(zap.WarnLevel, msg, kv...)
}

func Error(ctx context.Context, msg string, kv ...interface{}) {
	kv = append(kv, "traceid", ctx.Value("traceid"), "spanid", ctx.Value("spanid"), "pspanid", ctx.Value("pspanid"))
	l = getLogInstance()
	l.Log(zap.ErrorLevel, msg, kv...)
}
