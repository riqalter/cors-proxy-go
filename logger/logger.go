package logger

import (
	"log"
	"os"
	"time"
)

type Logger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

var DefaultLogger *Logger

func init() {
	DefaultLogger = New()
}

// New creates a new logger instance
func New() *Logger {
	return &Logger{
		infoLogger:    log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warningLogger: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger:   log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info logs an informational message
func (l *Logger) Info(v ...interface{}) {
	l.infoLogger.Println(v...)
}

// Infof logs a formatted informational message
func (l *Logger) Infof(format string, v ...interface{}) {
	l.infoLogger.Printf(format, v...)
}

// Warning logs a warning message
func (l *Logger) Warning(v ...interface{}) {
	l.warningLogger.Println(v...)
}

// Warningf logs a formatted warning message
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warningLogger.Printf(format, v...)
}

// Error logs an error message
func (l *Logger) Error(v ...interface{}) {
	l.errorLogger.Println(v...)
}

// Errorf logs a formatted error message
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
}

// Convenience functions using the default logger
func Info(v ...interface{}) {
	DefaultLogger.Info(v...)
}

func Infof(format string, v ...interface{}) {
	DefaultLogger.Infof(format, v...)
}

func Warning(v ...interface{}) {
	DefaultLogger.Warning(v...)
}

func Warningf(format string, v ...interface{}) {
	DefaultLogger.Warningf(format, v...)
}

func Error(v ...interface{}) {
	DefaultLogger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	DefaultLogger.Errorf(format, v...)
}

// LogRequest logs details about an HTTP request
func LogRequest(method, url, remoteAddr, userAgent string, duration time.Duration) {
	Infof("Request: %s %s from %s [%s] took %v", 
		method, url, remoteAddr, userAgent, duration)
}

// LogProxyRequest logs details about a proxy request
func LogProxyRequest(targetURL, method, remoteAddr string, statusCode int, duration time.Duration) {
	if statusCode >= 400 {
		Warningf("Proxy: %s %s from %s -> Status: %d, Duration: %v", 
			method, targetURL, remoteAddr, statusCode, duration)
	} else {
		Infof("Proxy: %s %s from %s -> Status: %d, Duration: %v", 
			method, targetURL, remoteAddr, statusCode, duration)
	}
}
