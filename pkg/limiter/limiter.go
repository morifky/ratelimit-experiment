package limiter

import (
	"sync"
)

type RateLimiter struct {
	visitors map[string]*TokenBucket
	mu       sync.Mutex
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		visitors: make(map[string]*TokenBucket),
	}
}

func (rl *RateLimiter) GetVisitor(ip string, rate float64, capacity float64) *TokenBucket {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	limiter, exists := rl.visitors[ip]
	if !exists {
		limiter = NewTokenBucket(rate, capacity)
		rl.visitors[ip] = limiter
	}
	return limiter
}
