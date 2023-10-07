package router

import "github.com/liuxing95/gin-react-admin/router/system"

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
