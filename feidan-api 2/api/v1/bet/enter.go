package bet

import (
	"feidan-api/service"
)

type ApiGroup struct {
	BetApi
}

var (
	apiService = service.ServiceGroupApp.BetServiceGroup
)
