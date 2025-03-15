package model

import "time"

type LimitHistory struct {
	ID         int32     `db:"id"`
	LimitID    int32     `db:"limit_id"`
	TimeAmount int32     `db:"time_amount"`
	SentAt     time.Time `db:"sent_at"`
	LimitDate  time.Time `db:"limit_date"`
}
