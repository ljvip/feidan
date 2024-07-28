package router

import (
	"feidan-api/router/account"
	"feidan-api/router/bet"
)

type RouterGroup struct {
	Account account.RouterGroup
	Bet     bet.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
