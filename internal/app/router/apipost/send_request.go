package apipost

/*
@Date:2022/2/7
@Author [Ambi](https://github.com/AmbitionLover)
发送 http 请求
*/
import (
	v1 "github.com/AmbitionLover/go-snake/internal/app/api/v1"
	"github.com/gin-gonic/gin"
)

type SendRequestRouter struct{}

func (s *SendRequestRouter) InitSendRequestRouter(r *gin.RouterGroup) {
	apiRouter := r.Group("request")
	apiRouterApi := v1.ApiGroupApp.ApiPostGroup.DefaultSendRequestApi
	{
		apiRouter.GET("send", apiRouterApi.SendRequest) //发送请求
	}
}
