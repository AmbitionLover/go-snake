package portal

import (
	"github.com/AmbitionLover/go-snake/pkg/model"
	uuid "github.com/satori/go.uuid"
)

/*
@Date:2022/2/9
@Author [Ambi](https://github.com/AmbitionLover)

*/

type PorUser struct {
	model.BaseModel
	UUID      uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`                                                           // 用户UUID
	Username  string    `json:"userName" gorm:"comment:用户登录名"`                                                        // 用户登录名
	Password  string    `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	NickName  string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	HeaderImg string    `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
}
