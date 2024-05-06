package http

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func loggingMiddleware(logger *zap.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Wrap the response writer to capture data for logging
			wrappedWriter := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			// Processing the request
			next.ServeHTTP(wrappedWriter, r)

			// After request is processed
			duration := time.Since(start)

			// Log details
			logger.Info("request",
				zap.String("method", r.Method),
				zap.String("uri", r.RequestURI),
				zap.String("host", r.Host),
				zap.String("remote_ip", r.RemoteAddr),
				zap.Int("status", wrappedWriter.statusCode),
				zap.Duration("latency", duration),
				zap.Int64("response_size", wrappedWriter.size),
				zap.String("protocol", r.Proto),
			)
		})
	}
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int64
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += int64(size)
	return size, err
}
