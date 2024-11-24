package middleware

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"ratelimit/pkg/limiter"
	response "ratelimit/utils"

	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

func SetMiddlewareRatelimit(next http.HandlerFunc, r rate.Limit, b int) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ip, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			response.WriteError(w, http.StatusInternalServerError, errors.New("internal server error"))
			return
		}
		limiter := limiter.GetVisitor(ip, r, b)
		zap.S().Info(fmt.Sprintf("client IP address= %v", ip))
		if !limiter.Allow() {
			response.WriteError(w, http.StatusTooManyRequests, errors.New("too many requests"))
			zap.S().Warn("too many requests")
			return
		}
		next(w, req)
	}
}
