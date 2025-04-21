package timer

import "time"

type Timer struct{}

func New() *Timer {
	return &Timer{}
}

func (t *Timer) Now() time.Time {
	return time.Now()
}
