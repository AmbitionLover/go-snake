package cryptography

import (
	"encoding/hex"
	"github.com/AmbitionLover/go-snake/internal/pkg/model/response"
	"github.com/AmbitionLover/go-snake/pkg/model/crypto"
	cryptoUtil "github.com/AmbitionLover/go-snake/pkg/utils/crypto"
	"github.com/gin-gonic/gin"
)

/*
@Date:2022/2/27
@Author [Ambi](https://github.com/AmbitionLover)

*/

func (d *CryptographyMd5Api) HashEncrypt(c *gin.Context) {
	var md crypto.HashEncrypt
	_ = c.ShouldBindJSON(&md)

	hash, err := cryptoUtil.MatchHash(md.Name)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	hash.Write([]byte(md.Plaintext))

	response.OkWithData(hex.EncodeToString(hash.Sum(nil)), c)
}
