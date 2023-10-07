package system

import "github.com/liuxing95/gin-react-admin/service"

type ApiGroup struct {
	JwtApi
	SystemApi
	AuthorityBtnApi
}

var (
	apiService          = service.ServiceGroupApp.SystemServiceGroup.ApiService
	jwtService          = service.ServiceGroupApp.SystemServiceGroup.JwtService
	authorityBtnService = service.ServiceGroupApp.SystemServiceGroup.AuthorityBtnService
)
