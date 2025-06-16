package middleware

import (
	"net/http"
	"time"

	"cors-proxy/logger"
)

// LoggingResponseWriter wraps http.ResponseWriter to capture status code
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// Logging middleware logs all incoming HTTP requests
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Wrap the response writer to capture status code
		lrw := NewLoggingResponseWriter(w)
		
		// Process the request
		next.ServeHTTP(lrw, r)
		
		// Log the request details
		duration := time.Since(start)
		userAgent := r.Header.Get("User-Agent")
		if userAgent == "" {
			userAgent = "Unknown"
		}
		
		// Log with different levels based on status code
		if lrw.statusCode >= 400 {
			logger.Warningf("Request: %s %s from %s [%s] -> Status: %d, Duration: %v",
				r.Method, r.URL.Path, r.RemoteAddr, userAgent, lrw.statusCode, duration)
		} else {
			logger.Infof("Request: %s %s from %s [%s] -> Status: %d, Duration: %v",
				r.Method, r.URL.Path, r.RemoteAddr, userAgent, lrw.statusCode, duration)
		}
	})
}
