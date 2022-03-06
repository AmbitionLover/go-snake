package init

import (
	"fmt"
	portal "github.com/AmbitionLover/go-snake/internal/app/router"
	global "github.com/AmbitionLover/go-snake/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

/*
@Date:2022/2/7
@Author [Ambi](https://github.com/AmbitionLover)

*/

type server interface {
	ListenAndServe() error
}

func RunServe() {
	// 1、启动admin
	initAdmin()
	// 2、启动门户
	initPortal()
}

// 暂时未提供
func initAdmin() {
	global.LOG.Info("Admin 暂不支持")
}

func initPortal() {
	routers := portal.Routers()
	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(address, routers)
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("Portal run success on ", zap.String("address", address))
	global.LOG.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
