package ctxtime

import (
	"testing"
	"time"
)

func TestClock_Now(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		c := NewClock()
		got := c.Now().Format("2006-01-02")
		want := time.Now().Format("2006-01-02")
		if got != want {
			t.Errorf("Clock.Now() = %s, want %s", got, want)
		}
	})
	t.Run("set hands of the clock", func(t *testing.T) {
		base := time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)
		c := NewClock(WithBase(base))
		got := c.Now().Format("2006-01-02")
		want := base.Format("2006-01-02")
		if got != want {
			t.Errorf("Clock.Now() = %s, want %s", got, want)
		}
	})

}
