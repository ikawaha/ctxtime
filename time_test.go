package ctxtime

import (
	"context"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	t.Run("default (time.Now())", func(t *testing.T) {
		ctx := context.TODO()
		got := Now(ctx).Format("2006-01-02")
		want := time.Now().Format("2006-01-02")
		if got != want {
			t.Errorf("Now() = %s, want %s", got, want)
		}
	})
	t.Run("set with the time", func(t *testing.T) {
		ctx := context.TODO()
		want := time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)
		ctx = WithTime(ctx, want)
		got := Now(ctx)
		if got, want := got.Format("2006-01-02"), want.Format("2006-01-02"); got != want {
			t.Errorf("Now() = %s, want %s", got, want)
		}
		if !got.Equal(want) {
			t.Errorf("Now() = %v, want %v = %v", got, got, want)
		}
	})
	t.Run("set with the clock", func(t *testing.T) {
		ctx := context.TODO()
		want := time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)
		c := NewClock(WithBase(want))
		ctx = WithClock(ctx, c)
		got := Now(ctx)
		if got, want := got.Format("2006-01-02"), want.Format("2006-01-02"); got != want {
			t.Errorf("Now() = %s, want %s", got, want)
		}
		if !got.After(want) {
			t.Errorf("Now() = %v, want %v", got, want)
		}
	})
}
