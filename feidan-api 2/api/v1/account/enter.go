package account

import (
	"feidan-api/service"
)

type ApiGroup struct {
	AccountApi
}

var (
	apiService = service.ServiceGroupApp.AccountServiceGroup
)
