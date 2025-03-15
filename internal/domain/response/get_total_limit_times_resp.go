package response

type GetTotalLimitTimesResp struct {
	Sum     int  `json:"sum"`
	Expired bool `json:"expired"`
}
