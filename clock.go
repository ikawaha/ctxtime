package ctxtime

import "time"

type Clock struct {
	start time.Time
	base  time.Time
}

type ClockOption func(*Clock)

func WithBase(t time.Time) ClockOption {
	return func(c *Clock) {
		c.base = t
	}
}

func NewClock(opts ...ClockOption) *Clock {
	now := time.Now()
	ret := Clock{
		start: now,
		base:  now,
	}
	for _, opt := range opts {
		opt(&ret)
	}
	return &ret
}

func (c *Clock) Now() time.Time {
	if c.start == c.base {
		return time.Now()
	}
	return c.base.Add(time.Since(c.start))
}
