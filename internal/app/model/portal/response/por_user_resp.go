package response

import "github.com/AmbitionLover/go-snake/internal/app/model/portal"

/*
@Date:2022/2/9
@Author [Ambi](https://github.com/AmbitionLover)
用户登录返回信息
*/
type PortalUserResponse struct {
	User portal.PorUser `json:"user"`
}

type LoginResponse struct {
	PortalUserResponse
	ApiToken  string `json:"api_token"`
	ExpiresAt int64  `json:"expiresAt"`
}
