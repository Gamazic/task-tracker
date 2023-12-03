package microframework

import (
	"net/http"
	"time"
	"tracker_backend/src/infrastructure"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (l *loggingResponseWriter) WriteHeader(code int) {
	l.statusCode = code
	l.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler, logger infrastructure.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		loggingW := &loggingResponseWriter{ResponseWriter: w}
		next.ServeHTTP(loggingW, req)
		logger.Infof("%s %s %d %s",
			req.Method,
			req.RequestURI,
			loggingW.statusCode,
			time.Since(start),
		)
	})
}

func MaxBytes(next http.Handler, maxBytes int64) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
		next.ServeHTTP(w, r)
	})
}