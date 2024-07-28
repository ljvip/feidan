package account

import (
	"feidan-api/api/v1"
	"github.com/gin-gonic/gin"
)

type AccountRouter struct{}

func (s *AccountRouter) InitAccountRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("")
	apiV1Router := Router.Group("v1")
	apiAccountV1Router := apiV1Router.Group("account")
	v1ApiRouterApi := v1.ApiGroupApp.AccountGroup.AccountApi
	{
		apiAccountV1Router.Any("getToken", v1ApiRouterApi.GetToken)
		apiAccountV1Router.Any("getUserInfo", v1ApiRouterApi.GetUserInfo)
		apiV1Router.Any("getList", v1ApiRouterApi.GetList)
		apiRouter.Any("auto_login", v1ApiRouterApi.AutoLogin)
	}
}
