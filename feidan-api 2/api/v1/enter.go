package v1

import (
	"feidan-api/api/v1/account"
	"feidan-api/api/v1/bet"
)

type ApiGroup struct {
	AccountGroup account.ApiGroup
	BetGroup     bet.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
