package request

type PushCurrentTime struct {
	LimitID    int32  `json:"limit_id"`
	TimeAmount string `json:"time_amount"`
}
