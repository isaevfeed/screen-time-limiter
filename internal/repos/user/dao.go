package user

import "time"

type dao struct {
	ID        int64      `db:"id"`
	FirstName string     `db:"first_name"`
	LastName  string     `db:"last_name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}
