package system

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/liuxing95/gin-react-admin/api/v1"
)

type AuthorityBtnRouter struct{}

func (r *AuthorityBtnRouter) InitAuthorityBtnRouter(Router *gin.RouterGroup) {
	authorityRouterWithoutRecord := Router.Group("authorityBtn")
	authorityBtnApi := v1.ApiGroupApp.SystemApiGroup.AuthorityBtnApi
	{
		authorityRouterWithoutRecord.POST("getAuthorityBtn", authorityBtnApi.GetAuthorityBtn)
		authorityRouterWithoutRecord.POST("setAuthorityBtn", authorityBtnApi.SetAuthorityBtn)
		authorityRouterWithoutRecord.POST("canRemoveAuthorityBtn", authorityBtnApi.CanRemoveAuthorityBtn)
	}
}
