package system

import "github.com/liuxing95/gin-react-admin/service"

type ApiGroup struct {
	JwtApi
	SystemApi
}

var (
	apiService = service.ServiceGroupApp.SystemServiceGroup.ApiService
	jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
)
