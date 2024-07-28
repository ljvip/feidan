package bet

import (
	"feidan-api/global"
	"feidan-api/model/bet/request"
	"feidan-api/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BetApi struct{}

func (s *BetApi) Bet(c *gin.Context) {
	var req request.BetReq
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	rsp, err := apiService.BetService.Bet(c, &req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(rsp, "", c)
}

func (s *BetApi) OutFly(c *gin.Context) {
	var req request.OutFlyReq
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	global.GVA_LOG.Info("飞单参数:", zap.Any("param", req))
	req.Token = c.Query("token")
	rsp, err := apiService.BetService.OutFly(c, &req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithCustomData(rsp, c)
}
