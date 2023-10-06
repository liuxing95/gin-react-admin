package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/liuxing95/gin-react-admin/global"
	"github.com/liuxing95/gin-react-admin/middleware"
	"github.com/liuxing95/gin-react-admin/plugin/email"
	"github.com/liuxing95/gin-react-admin/utils/plugin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}

func InstallPlugin(Router *gin.Engine) {
	PublicGroup := Router.Group("")
	fmt.Println("无鉴权插件安装==》", PublicGroup)
	PrivateGroup := Router.Group("")
	fmt.Println("鉴权插件安装==》", PrivateGroup)
	PrivateGroup.Use(middleware.JWTAuth())
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.GRA_CONFIG.Email.To,
		global.GRA_CONFIG.Email.From,
		global.GRA_CONFIG.Email.Host,
		global.GRA_CONFIG.Email.Secret,
		global.GRA_CONFIG.Email.Nickname,
		global.GRA_CONFIG.Email.Port,
		global.GRA_CONFIG.Email.IsSSL,
	))
}
