package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	ctx     context.Context
	traceId string
	spanId  string
	pSpanId string
	_logger *zap.Logger
}

func NewLogger(ctx context.Context) *Logger {
	var traceId string
	var spanId string
	var pSpanId string
	if ctx.Value("traceId") != nil {
		traceId = ctx.Value("traceId").(string)
	}
	if ctx.Value("spanId") != nil {
		spanId = ctx.Value("spanId").(string)
	}
	if ctx.Value("pSpanId") != nil {
		pSpanId = ctx.Value("pSpanId").(string)
	}
	return &Logger{
		ctx:     ctx,
		traceId: traceId,
		spanId:  spanId,
		pSpanId: pSpanId,
		_logger: _logger,
	}
}

func (l *Logger) log(lvl zapcore.Level, msg string, kv ...interface{}) {
	// 保证要打印的日志信息成对出现
	if len(kv)%2 != 0 {
		kv = append(kv, "keysAndValues is not a pair")
	}
	fields := make([]zap.Field, 0, len(kv)/2)
	for i := 0; i < len(kv); i += 2 {
		fields = append(fields, zap.Any(fmt.Sprint(kv[i]), kv[i+1]))
	}
	if ce := l._logger.Check(lvl, msg); ce != nil {
		ce.Write(fields...)
	}
}
func (l *Logger) Debug(msg string, kv ...interface{}) {
	l.log(zap.DebugLevel, msg, kv...)
}

func (l *Logger) Info(msg string, kv ...interface{}) {
	l.log(zap.InfoLevel, msg, kv...)
}

func (l *Logger) Warn(msg string, kv ...interface{}) {
	l.log(zap.WarnLevel, msg, kv...)
}

func (l *Logger) Error(msg string, kv ...interface{}) {
	l.log(zap.ErrorLevel, msg, kv...)
}
