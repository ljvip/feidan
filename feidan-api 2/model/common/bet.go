package common

type BetDataList struct {
	Pum    any    `json:"pum"`
	Amount any    `json:"amount"`
	Log    string `json:"log"`
	Error  string `json:"error,omitempty"`
}
