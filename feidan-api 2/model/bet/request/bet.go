package request

import "feidan-api/model/common"

type BetReq struct {
	Token   string     `json:"token"`
	Rows    any        `json:"rows"`
	Ce      string     `json:"ce"`
	Data    []*BetData `json:"data" binding:"required"`
	URL     string     `json:"url"`
	Version string     `json:"version"`
}
type BetData struct {
	Code  string                `json:"code"`
	Issue string                `json:"issue"`
	List  []*common.BetDataList `json:"list"`
}

type OutFlyReq struct {
	Rows  any        `json:"rows"`
	Ce    string     `json:"ce"`
	Data  []*BetData `json:"data"`
	Token string     `json:"token" form:"token"`
}
