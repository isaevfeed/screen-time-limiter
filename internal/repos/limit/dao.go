package limit

import "time"

type dao struct {
	ID       int32      `db:"id"`
	UserID   int32      `db:"user_id"`
	Amount   int32      `db:"time_limit"`
	CreateAt time.Time  `db:"created_at"`
	UpdateAt *time.Time `db:"updated_at"`
}
