package core

import (
	"github.com/liuxing95/gin-react-admin/global"
	"github.com/liuxing95/gin-react-admin/initialize"
)

func RunServer() {
	if global.GRA_CONFIG.System.UseMultipoint || global.GRA_CONFIG.System.UseRedis {
		// 初始化redis服务
		// initialize.Redis()
	}
	// 从db加载jwt数据
	if global.GRA_DB != nil {
		// system.LoadAll()
	}

	// 初始化路由
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")
}
