package request

type AddLimit struct {
	UserID int32  `json:"user_id"`
	Amount string `json:"amount"`
}
