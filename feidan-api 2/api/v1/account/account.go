package account

import (
	"feidan-api/model/account/request"
	accountRes "feidan-api/model/account/response"
	"feidan-api/model/common/response"
	"github.com/gin-gonic/gin"
)

type AccountApi struct{}

func (s *AccountApi) GetToken(c *gin.Context) {
	var req request.GetTokenReq
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	token, err := apiService.AccountService.GetToken(c, req.Url, req.Username, req.Password)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(&accountRes.GetTokenRsp{Token: token}, "登陆成功", c)
}

func (s *AccountApi) GetUserInfo(c *gin.Context) {
	var req request.GetUserInfoReq
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	info, err := apiService.AccountService.GetUserInfo(c, req.Url, req.Token)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(info, "获取用户信息成功", c)
}

func (s *AccountApi) GetList(c *gin.Context) {
	var req request.GetListReq
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, err := apiService.AccountService.GetList(c, req.Url, req.Token)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(list, "获取成功", c)
}

func (s *AccountApi) AutoLogin(c *gin.Context) {
	result, err := apiService.AccountService.AutoLogin(c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.HtmlString(result, c)
}
