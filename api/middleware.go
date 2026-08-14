package api

import "time"

type Limiter struct {
	Capacity float64
	Rate     float64
	Tokens   float64
	Last     time.Time
}


func NewLimiter(capacity int, rate float64) *Limiter {
	return &Limiter{
		Capacity: float64(capacity),
		Rate: rate,
		Tokens: float64(capacity),
		Last: time.Now(),
	}
}

func (l *Limiter) Allow() bool {
	//calc time elapsed
	now:=time.Now()
	elapsed:=now.Sub(l.Last).Seconds()

	//update limiter
	l.Tokens = min(l.Tokens+(l.Rate*elapsed),l.Capacity)
	l.Last=now

	//request rejected
	if l.Tokens < 1 {
		return false
	}

	//request accepted
	l.Tokens--

	return true
}