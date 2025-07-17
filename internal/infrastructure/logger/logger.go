package logger

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

// Logger interface defines the logging methods
type Logger interface {
	Info(message string, args ...interface{})
	Error(message string, args ...interface{})
	Debug(message string, args ...interface{})
	Warn(message string, args ...interface{})
}

// loggerImpl implements the Logger interface
type loggerImpl struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
	warnLogger  *log.Logger
}

// NewLogger creates a new logger instance
func NewLogger() Logger {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	if env == "production" {
		return newFileLogger()
	}
	return newConsoleLogger()
}

// newConsoleLogger creates a console logger for development
func newConsoleLogger() Logger {
	return &loggerImpl{
		infoLogger:  log.New(os.Stdout, "[INFO] ", log.LstdFlags),
		errorLogger: log.New(os.Stderr, "[ERROR] ", log.LstdFlags),
		debugLogger: log.New(os.Stdout, "[DEBUG] ", log.LstdFlags),
		warnLogger:  log.New(os.Stdout, "[WARN] ", log.LstdFlags),
	}
}

// newFileLogger creates a file logger for production
func newFileLogger() Logger {
	// Create logs directory if it doesn't exist
	logsDir := "logs"
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		log.Fatalf("Failed to create logs directory: %v", err)
	}

	// Create log files with timestamp
	timestamp := time.Now().Format("2006-01-02")

	infoFile, err := os.OpenFile(filepath.Join(logsDir, "info-"+timestamp+".log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open info log file: %v", err)
	}

	errorFile, err := os.OpenFile(filepath.Join(logsDir, "error-"+timestamp+".log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open error log file: %v", err)
	}

	debugFile, err := os.OpenFile(filepath.Join(logsDir, "debug-"+timestamp+".log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open debug log file: %v", err)
	}

	warnFile, err := os.OpenFile(filepath.Join(logsDir, "warn-"+timestamp+".log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open warn log file: %v", err)
	}

	return &loggerImpl{
		infoLogger:  log.New(infoFile, "[INFO] ", log.LstdFlags),
		errorLogger: log.New(errorFile, "[ERROR] ", log.LstdFlags),
		debugLogger: log.New(debugFile, "[DEBUG] ", log.LstdFlags),
		warnLogger:  log.New(warnFile, "[WARN] ", log.LstdFlags),
	}
}

// Info logs an info message
func (l *loggerImpl) Info(message string, args ...interface{}) {
	if len(args) > 0 {
		l.infoLogger.Printf(message, args...)
	} else {
		l.infoLogger.Println(message)
	}
}

// Error logs an error message
func (l *loggerImpl) Error(message string, args ...interface{}) {
	if len(args) > 0 {
		l.errorLogger.Printf(message, args...)
	} else {
		l.errorLogger.Println(message)
	}
}

// Debug logs a debug message
func (l *loggerImpl) Debug(message string, args ...interface{}) {
	if len(args) > 0 {
		l.debugLogger.Printf(message, args...)
	} else {
		l.debugLogger.Println(message)
	}
}

// Warn logs a warning message
func (l *loggerImpl) Warn(message string, args ...interface{}) {
	if len(args) > 0 {
		l.warnLogger.Printf(message, args...)
	} else {
		l.warnLogger.Println(message)
	}
}
