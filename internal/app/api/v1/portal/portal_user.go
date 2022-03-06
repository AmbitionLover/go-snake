package portal

/*
@Date:2022/2/8
@Author [Ambi](https://github.com/AmbitionLover)
门户网站用户
*/
import (
	"github.com/AmbitionLover/go-snake/internal/app/model/portal"
	"github.com/AmbitionLover/go-snake/internal/app/model/portal/request"
	portalRes "github.com/AmbitionLover/go-snake/internal/app/model/portal/response"
	"github.com/AmbitionLover/go-snake/internal/pkg/model/response"
	"github.com/AmbitionLover/go-snake/internal/pkg/utils"
	global "github.com/AmbitionLover/go-snake/pkg"
	"github.com/AmbitionLover/go-snake/pkg/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type PortalUserApi struct {
}

// Login @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body systemReq.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /portal/login [post]
func (p *PortalUserApi) Login(c *gin.Context) {
	// 先模拟一个人
	p.tokenNext(c, portal.PorUser{
		BaseModel: model.BaseModel{ID: 1, CreatedAt: time.Time{}},
		Username:  "Ambi",
		NickName:  "赵帅",
	})
}

func (p *PortalUserApi) VerifyToken(c *gin.Context) {
	if _, err := utils.GetClaims(c); err != nil {
		response.FailWithMessage("Failed to get token!", c)
		return
	} else {
		//response.OkWithDetailed(portalRes.LoginResponse{
		//	PortalUserResponse:portalRes.PortalUserResponse{User:  portal.PorUser{UUID: claims.UUID}},
		//	ApiToken:           token,
		//	ExpiresAt:          claims.StandardClaims.ExpiresAt * 1000,
		//}, "登录成功", c)
		response.Ok(c)
	}
}

// 登录以后签发jwt
func (p *PortalUserApi) tokenNext(c *gin.Context, user portal.PorUser) {
	j := utils.NewJWT()
	claims := j.CreateClaims(request.BaseClaims{
		ID:       user.ID,
		UUID:     user.UUID,
		NickName: user.NickName,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.LOG.Error("Failed to get token!", zap.Error(err))
		response.FailWithMessage("Failed to get token!", c)
		return
	}
	response.OkWithDetailed(portalRes.LoginResponse{
		PortalUserResponse: portalRes.PortalUserResponse{User: user},
		ApiToken:           token,
		ExpiresAt:          claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
}
