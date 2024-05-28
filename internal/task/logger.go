package task

import (
	"context"

	"github.com/yituoshiniao/kit/xlog"
	"go.uber.org/zap"
)

// NewLogger 自定义 cron log
func NewLogger(l *zap.Logger, ctx context.Context) *Logger {
	logger := &Logger{
		logger: l,
		Ctx:    ctx,
	}
	return logger
}

type Logger struct {
	logger *zap.Logger
	Ctx    context.Context
}

func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	xlog.S(l.Ctx).Infow(msg, "keysAndValues", keysAndValues)
}

func (l *Logger) Error(err error, msg string, keysAndValues ...interface{}) {
	xlog.S(l.Ctx).Errorw(msg, "err", err, "keysAndValues", keysAndValues)
}
