package slog

import (
	"fmt"
	"log/slog"

	"github.com/go-kratos/kratos/v2/log"
)

var _ log.Logger = (*Logger)(nil)

type Logger struct {
	log *slog.Logger
}

func NewLogger(log *slog.Logger) *Logger {
	return &Logger{log: log}
}

func (l Logger) Log(level log.Level, keyvals ...interface{}) error {
	keylen := len(keyvals)
	if keylen == 0 || keylen%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}
	data := make([]any, 0, (keylen/2)+1)
	for i := 0; i < keylen; i += 2 {
		data = append(data, fmt.Sprint(keyvals[i], keyvals[i+1]))
	}
	switch level {
	case log.LevelDebug:
		l.log.Debug("", data)
	case log.LevelInfo:
		l.log.Info("", data...)
	case log.LevelWarn:
		l.log.Warn("", data...)
	case log.LevelError, log.LevelFatal:
		l.log.Error("", data...)
	default:
		l.log.Debug("", data...)
	}
	return nil
}
