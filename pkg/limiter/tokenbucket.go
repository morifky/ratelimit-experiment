package limiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	rate         float64   // token per second
	capacity     float64   // max capacity
	tokens       float64   // current token
	lastFillTime time.Time // last time token were added
	mu           sync.Mutex
}

func NewTokenBucket(rate float64, capacity float64) *TokenBucket {
	return &TokenBucket{
		rate:     rate,
		capacity: capacity,
	}
}

func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsedTime := now.Sub(tb.lastFillTime).Seconds()
	newTokens := elapsedTime * tb.rate

	tb.tokens = tb.tokens + newTokens

	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	tb.lastFillTime = now
}

func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()
	if tb.tokens > -1 {
		tb.tokens--
		return true
	}
	return false
}
