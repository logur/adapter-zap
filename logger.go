// Package template provides a Logur adapter for TEMPLATE.
package template

import (
	"github.com/goph/logur"
)

// Logger is a Logur adapter for TEMPLATE.
type Logger struct {
}

// New returns a new Logur logger.
// If logger is nil, a default instance is created.
func New(logger interface{}) *Logger {
	if logger == nil {
		return &Logger{}
	}

	return &Logger{}
}

// Trace implements the Logur Logger interface.
func (l *Logger) Trace(msg string, fields ...map[string]interface{}) {

}

// Debug implements the Logur Logger interface.
func (l *Logger) Debug(msg string, fields ...map[string]interface{}) {

}

// Info implements the Logur Logger interface.
func (l *Logger) Info(msg string, fields ...map[string]interface{}) {

}

// Warn implements the Logur Logger interface.
func (l *Logger) Warn(msg string, fields ...map[string]interface{}) {

}

// Error implements the Logur Logger interface.
func (l *Logger) Error(msg string, fields ...map[string]interface{}) {

}

// LevelEnabled implements the Logur LevelEnabler interface.
func (l *Logger) LevelEnabled(level logur.Level) bool {
	switch level {
	case logur.Trace:
		return true
	case logur.Debug:
		return true
	case logur.Info:
		return true
	case logur.Warn:
		return true
	case logur.Error:
		return true
	}

	return true
}
