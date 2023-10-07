package initialize

import (
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	InstallPlugin(Router) // 安装插件
	// systemRouter := router.RouterGroupApp.System
	// {
	// 	systemRouter.
	// }
	return Router
}
