package middleware

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.SetFormatter(&log.JSONFormatter{})
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)
		log.WithFields(log.Fields{
			"code":     wrapper.StatusCode,
			"method":   r.Method,
			"url":      r.URL.Path,
			"execTime": fmt.Sprintf("%dms", time.Since(start).Milliseconds()),
		}).Info("LOGRUS LOG")
	})
}
