package ctxtime

import (
	"context"
	"time"
)

type timeKey struct{}

func WithTime(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, timeKey{}, t)
}

func WithClock(ctx context.Context, c *Clock) context.Context {
	return context.WithValue(ctx, timeKey{}, c)
}

func Now(ctx context.Context) time.Time {
	switch v := ctx.Value(timeKey{}).(type) {
	case time.Time:
		return v
	case *Clock:
		return v.Now()
	}
	return time.Now()
}
