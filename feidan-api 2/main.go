package main

import (
	core2 "feidan-api/core"
	"feidan-api/global"
	initialize2 "feidan-api/initialize"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	global.GVA_VP = core2.Viper() // 初始化Viper
	global.GVA_LOG = core2.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize2.Gorm() // gorm连接数据库
	initialize2.DBList()
	if global.GVA_DB != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core2.RunWindowsServer()
}
