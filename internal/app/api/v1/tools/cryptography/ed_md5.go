package cryptography

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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

func hmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
