package request

type GetTokenReq struct {
	Url      string `form:"url" binding:"required"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type GetUserInfoReq struct {
	Url   string `form:"url" binding:"required"`
	Token string `form:"token" binding:"required"`
}

type GetListReq struct {
	Url   string `form:"url" binding:"required"`
	Token string `form:"token" binding:"required"`
}
