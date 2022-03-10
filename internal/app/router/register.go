package router

/*
@Date:2022/2/17
@Author Ambi 赵帅
门户前台路由注册
*/
import (
	"github.com/AmbitionLover/go-snake/internal/app/middleware"
	"github.com/AmbitionLover/go-snake/internal/app/router/apipost"
	"github.com/AmbitionLover/go-snake/internal/app/router/portal"
	"github.com/AmbitionLover/go-snake/internal/app/router/tools"
	global "github.com/AmbitionLover/go-snake/pkg"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	ApiPost apipost.RouterGroup
	Tools   tools.RouterGroup
	Portal  portal.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

func Routers() *gin.Engine {
	Router := gin.Default()
	global.LOG.Info("init Portal router~~~")

	useMiddleware(Router)

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}

	// 获取路由组实例
	// 1、自动化测试
	apiPostGroup(PublicGroup)
	// 2、工具类集合
	toolsGroup(PublicGroup)
	// 3、门户路由
	portalGroup(PublicGroup)

	return Router
}

func portalGroup(PublicGroup *gin.RouterGroup) {
	group := RouterGroupApp.Portal
	{
		group.InitPortalUserRouter(PublicGroup)
	}
}

func toolsGroup(r *gin.RouterGroup) {
	apiRouter := r.Group("tools")
	group := RouterGroupApp.Tools
	{
		// 密码学工具
		group.Cryptography.InitPortalUserRouter(apiRouter)
	}
}

func apiPostGroup(PublicGroup *gin.RouterGroup) {
	group := RouterGroupApp.ApiPost
	{
		group.InitSendRequestRouter(PublicGroup)
	}
}

func useMiddleware(e *gin.Engine) {
	e.Use(middleware.Cors()) // 直接放行全部跨域请求
}
