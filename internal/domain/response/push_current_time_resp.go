package response

type (
	PushCurrentTimeResp struct {
		Data PushCurrentTimeRespData `json:"data"`
	}

	PushCurrentTimeRespData struct {
		Message     string `json:"message"`
		Expired     bool   `json:"expired"`
		TimeBalance int32  `json:"time_balance"`
	}
)
