package logger

import (
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	// Test development mode
	os.Setenv("APP_ENV", "development")
	logger := NewLogger()

	// Test info logging
	logger.Info("Test info message")
	logger.Error("Test error message")
	logger.Debug("Test debug message")
	logger.Warn("Test warning message")
}

func TestNewLoggerProduction(t *testing.T) {
	// Test production mode
	os.Setenv("APP_ENV", "production")
	logger := NewLogger()

	// Test info logging
	logger.Info("Test production info message")
	logger.Error("Test production error message")
	logger.Debug("Test production debug message")
	logger.Warn("Test production warning message")
}

func TestLoggerInterface(t *testing.T) {
	// Test that logger implements the interface
	var _ Logger = NewLogger()
}
