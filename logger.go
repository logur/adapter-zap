// Package zap provides a Logur adapter for Uber's Zap.
package zap

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"logur.dev/logur"

	"logur.dev/adapter/zap/internal/keyvals"
)

// Logger is a Logur adapter for Uber's Zap.
type Logger struct {
	logger *zap.SugaredLogger
	core   zapcore.Core
}

// New returns a new Logur logger.
// If logger is nil, a default instance is created.
func New(logger *zap.Logger) *Logger {
	if logger == nil {
		logger = zap.L()
	}

	return &Logger{
		logger: logger.Sugar(),
		core:   logger.Core(),
	}
}

// Trace implements the Logur Logger interface.
func (l *Logger) Trace(msg string, fields ...map[string]interface{}) {
	// Fall back to Debug
	l.Debug(msg, fields...)
}

// Debug implements the Logur Logger interface.
func (l *Logger) Debug(msg string, fields ...map[string]interface{}) {
	if !l.core.Enabled(zap.DebugLevel) {
		return
	}

	l.logger.Debugw(msg, l.keyvals(fields)...)
}

// Info implements the Logur Logger interface.
func (l *Logger) Info(msg string, fields ...map[string]interface{}) {
	if !l.core.Enabled(zap.InfoLevel) {
		return
	}

	l.logger.Infow(msg, l.keyvals(fields)...)
}

// Warn implements the Logur Logger interface.
func (l *Logger) Warn(msg string, fields ...map[string]interface{}) {
	if !l.core.Enabled(zap.WarnLevel) {
		return
	}

	l.logger.Warnw(msg, l.keyvals(fields)...)
}

// Error implements the Logur Logger interface.
func (l *Logger) Error(msg string, fields ...map[string]interface{}) {
	if !l.core.Enabled(zap.ErrorLevel) {
		return
	}

	l.logger.Errorw(msg, l.keyvals(fields)...)
}

func (l *Logger) keyvals(fields []map[string]interface{}) []interface{} {
	var kvs []interface{}
	if len(fields) > 0 {
		kvs = keyvals.FromMap(fields[0])
	}

	return kvs
}

func (l *Logger) TraceContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Trace(msg, fields...)
}

func (l *Logger) DebugContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Debug(msg, fields...)
}

func (l *Logger) InfoContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Info(msg, fields...)
}

func (l *Logger) WarnContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Warn(msg, fields...)
}

func (l *Logger) ErrorContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Error(msg, fields...)
}

// LevelEnabled implements the Logur LevelEnabler interface.
func (l *Logger) LevelEnabled(level logur.Level) bool {
	switch level {
	case logur.Trace:
		return l.core.Enabled(zap.DebugLevel)
	case logur.Debug:
		return l.core.Enabled(zap.DebugLevel)
	case logur.Info:
		return l.core.Enabled(zap.InfoLevel)
	case logur.Warn:
		return l.core.Enabled(zap.WarnLevel)
	case logur.Error:
		return l.core.Enabled(zap.ErrorLevel)
	}

	return true
}
