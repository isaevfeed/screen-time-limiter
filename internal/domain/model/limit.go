package model

type Limit struct {
	ID     int32
	UserID int32
	Amount int32
}

func (l Limit) Expired(amount int32) bool {
	return l.Amount <= amount
}
