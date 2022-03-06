package tools

import (
	v1 "github.com/AmbitionLover/go-snake/internal/app/api/v1"
	"github.com/gin-gonic/gin"
)

/*
@Date:2022/3/5
@Author [Ambi](https://github.com/AmbitionLover)
密码学研究
*/


type CryptographyRouter struct{}

func (s *CryptographyRouter) InitPortalUserRouter(r *gin.RouterGroup) {
	apiRouter := r.Group("crypto")
	apiRouterApi := v1.ApiGroupApp.ToolsGroup.Cryptography
	{
		apiRouter.POST("md5", apiRouterApi.Encrypt)      //发送请求
		apiRouter.POST("hash", apiRouterApi.HashEncrypt) //发送请求
	}
}
