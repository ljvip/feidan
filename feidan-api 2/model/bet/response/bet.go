package response

import "feidan-api/model/common"

type BetRsp struct {
	Data *BetRspData `json:"data"`
	Info *BetInfo    `json:"info"`
}

type BetInfo struct {
	Balance  float64 `json:"balance"`
	Betting  float64 `json:"betting"`
	MaxLimit float64 `json:"maxLimit"`
	Result   float64 `json:"result"`
	Type     int32   `json:"type"`
}

type BetRspData struct {
	State   int32        `json:"state"`
	Success []*BetResult `json:"success"`
	Failure []*BetResult `json:"failure"`
}

type BetResult struct {
	Code  string                `json:"code"`
	Issue string                `json:"issue"`
	List  []*common.BetDataList `json:"list"`
}
