package service

import (
	"feidan-api/service/account"
	"feidan-api/service/bet"
)

type ServiceGroup struct {
	AccountServiceGroup account.ServiceGroup
	BetServiceGroup     bet.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
