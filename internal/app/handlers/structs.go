package handlers

import "time"

type (
	timer interface {
		Now() time.Time
	}
)
