package middleware

import (
	"fmt"
	"net/http"

	"github.com/car12o/audio-mastery/pkg/logger"
	"github.com/felixge/httpsnoop"
)

type RequestLog struct {
	log logger.Service
}

type httpRequestInfo struct {
	method   string
	uri      string
	code     int
	size     int64
	duration int64
}

func NewRequestLog(log logger.Service) *RequestLog {
	return &RequestLog{log}
}

func (rl *RequestLog) Handle(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpReqInfo := httpRequestInfo{
			method: r.Method,
			uri:    r.URL.String(),
		}

		m := httpsnoop.CaptureMetrics(handler, w, r)
		httpReqInfo.code = m.Code
		httpReqInfo.size = m.Written
		httpReqInfo.duration = m.Duration.Milliseconds()

		rl.log.Info(fmt.Sprintf(
			"%s %s %d | %dms | %db",
			httpReqInfo.method,
			httpReqInfo.uri,
			httpReqInfo.code,
			httpReqInfo.duration,
			httpReqInfo.size,
		))
	})
}
