package middleware

import (
	"net/http"
	"time"

	"github.com/urfave/negroni"
	"go.uber.org/zap"
)

type LoggerEntry struct {
	StartTime string
	Status    int
	Duration  time.Duration
	Hostname  string
	Method    string
	Path      string
	Request   *http.Request
}

type ZapLogger struct {
	log *zap.SugaredLogger
}

func NewZapSDLogger(log *zap.SugaredLogger) *ZapLogger {
	return &ZapLogger{
		log: log,
	}
}

func (h *ZapLogger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()

	next(rw, r)

	res := rw.(negroni.ResponseWriter)

	log := &LoggerEntry{
		Status:   res.Status(),
		Duration: time.Since(start),
		Hostname: r.Host,
		Method:   r.Method,
		Path:     r.URL.Path,
		Request:  r,
	}

	h.log.Infow("request", "ID", r.Header.Get("X-Request-Id"), "method", log.Method, "path", log.Path, "status", log.Status, "duration", log.Duration)
}
