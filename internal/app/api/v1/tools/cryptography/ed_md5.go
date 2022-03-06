package cryptography

import (
	"github.com/AmbitionLover/go-snake/internal/pkg/model/response"
	"github.com/AmbitionLover/go-snake/pkg/model/crypto"
	cryptoUtil "github.com/AmbitionLover/go-snake/pkg/utils/crypto"
	"github.com/gin-gonic/gin"
)

/*
@Date:2022/2/8
@Author [Ambi](https://github.com/AmbitionLover)
MD5 加解密
*/

type CryptographyMd5Api struct{}

func (d *CryptographyMd5Api) Encrypt(c *gin.Context) {
	var md crypto.Md5Encrypt
	_ = c.ShouldBindJSON(&md)
	response.OkWithData(cryptoUtil.Md5De(md), c)
}
