package global

import (
	"context"
	"log"
)

type ILog interface {
	Errorf(ctx context.Context, format string, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
}

type consoleLog struct {
}

func (l *consoleLog) Warnf(ctx context.Context, format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *consoleLog) Infof(ctx context.Context, format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *consoleLog) Errorf(ctx context.Context, format string, args ...interface{}) {
	log.Printf(format, args...)
}
