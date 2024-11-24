package limiter

import (
	"sync"

	"golang.org/x/time/rate"
)

var visitors = make(map[string]*rate.Limiter)
var mu sync.Mutex

func GetVisitor(ip string, r rate.Limit, b int) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exist := visitors[ip]

	if !exist {
		limiter = rate.NewLimiter(r, b)
		visitors[ip] = limiter
	}
	return limiter
}
