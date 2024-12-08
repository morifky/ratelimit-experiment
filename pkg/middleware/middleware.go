package middleware

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"ratelimit/pkg/limiter"
	response "ratelimit/utils"

	"go.uber.org/zap"
)

func SetMiddlewareRatelimit(next http.HandlerFunc, rate float64, capacity float64, limiter *limiter.RateLimiter) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ip, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			response.WriteError(w, http.StatusInternalServerError, errors.New("internal server error"))
			return
		}
		lim := limiter.GetVisitor(ip, rate, capacity)
		zap.S().Info(fmt.Sprintf("client IP address= %v", ip))
		if !lim.Allow() {
			response.WriteError(w, http.StatusTooManyRequests, errors.New("too many requests"))
			zap.S().Warn("too many requests")
			return
		}
		next(w, req)
	}
}
