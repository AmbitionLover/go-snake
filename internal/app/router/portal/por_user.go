package portal

import (
	v1 "github.com/AmbitionLover/go-snake/internal/app/api/v1"
	"github.com/gin-gonic/gin"
)

/*
@Date:2022/2/9
@Author [Ambi](https://github.com/AmbitionLover)
门户用户
*/

type PortalUserRouter struct{}

func (s *PortalUserRouter) InitPortalUserRouter(Router *gin.RouterGroup) {
	userRouterWithoutRecord := Router.Group("portal")
	baseApi := v1.ApiGroupApp.PortalGroup.PortalUserApi
	{
		userRouterWithoutRecord.POST("login", baseApi.Login)
		userRouterWithoutRecord.POST("verify_token", baseApi.VerifyToken)
	}
}
