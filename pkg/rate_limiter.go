package pkg

import (
	"sync"
	"time"
)

type Limiter struct {
	mu       sync.Mutex
	capacity float64
	rate     float64
	tokens   float64
	last     time.Time
}

type RateLimiter struct {
	mu       sync.Mutex
	limiters map[string]*Limiter
}

func NewLimiter(capacity int, rate float64) *Limiter {
	return &Limiter{
		capacity: float64(capacity),
		rate:     rate,
		tokens:   float64(capacity),
		last:     time.Now(),
	}
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		limiters: make(map[string]*Limiter),
	}
}

func (l *Limiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	//calc time elapsed
	now := time.Now()
	elapsed := now.Sub(l.last).Seconds()

	//update limiter
	l.tokens = min(l.tokens+(l.rate*elapsed), l.capacity)
	l.last = now

	//request rejected
	if l.tokens < 1 {
		return false
	}

	//request accepted
	l.tokens--

	return true
}

func (r *RateLimiter) Allow(ip string) bool {
	r.mu.Lock()

	limiter, exists := r.limiters[ip]

	if !exists {
		limiter = NewLimiter(5, 5.0/60.0)
		r.limiters[ip] = limiter
	}

	r.mu.Unlock()

	return limiter.Allow()
}