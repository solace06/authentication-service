package api

import "time"

type Limiter struct {
	capacity float64
	rate     float64
	tokens   float64
	last     time.Time
}


func NewLimiter(capacity int, rate float64) *Limiter {
	return &Limiter{
		capacity: float64(capacity),
		rate: rate,
		tokens: float64(capacity),
		last: time.Now(),
	}
}

func (l *Limiter) Allow() bool {
	//calc time elapsed
	now:=time.Now()
	elapsed:=now.Sub(l.last).Seconds()

	//update limiter
	l.tokens = min(l.tokens+(l.rate*elapsed),l.capacity)
	l.last=now

	//request rejected
	if l.tokens < 1 {
		return false
	}

	//request accepted
	l.tokens--

	return true
}