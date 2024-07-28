package bet

import (
	"feidan-api/api/v1"
	"github.com/gin-gonic/gin"
)

type BetRouter struct{}

func (s *BetRouter) InitBetRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("")
	apiV1Router := Router.Group("v1")
	v1ApiRouterApi := v1.ApiGroupApp.BetGroup.BetApi
	{
		apiV1Router.POST("bet", v1ApiRouterApi.Bet)
		apiRouter.Any("outFly", v1ApiRouterApi.OutFly)
	}
}
