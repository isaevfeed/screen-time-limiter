package response

type GetTotalLimitTimesResp struct {
	Sum         int32 `json:"sum"`
	TimeBalance int32 `json:"time_balance"`
	Expired     bool  `json:"expired"`
}
