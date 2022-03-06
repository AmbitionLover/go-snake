package apipost

import (
	"github.com/AmbitionLover/go-snake/internal/pkg/model/response"
	"github.com/gin-gonic/gin"
)

/*
@Date:2022/2/7
@Author Ambi 赵帅

*/

type DefaultSendRequestApi struct{}

func (d *DefaultSendRequestApi) SendRequest(c *gin.Context) {
	response.OkWithMessage("创建成功", c)
}
