package cryptography

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/AmbitionLover/go-snake/internal/pkg/model/response"
	"github.com/AmbitionLover/go-snake/pkg/model/crypto"
	cryptoUtil "github.com/AmbitionLover/go-snake/pkg/utils/crypto"
	"github.com/gin-gonic/gin"
	"strings"
)

/*
@Date:2022/2/27
@Author [Ambi](https://github.com/AmbitionLover)

*/

func (d *CryptographyMd5Api) HashEncrypt(c *gin.Context) {
	var he crypto.HashEncrypt
	_ = c.ShouldBindJSON(&he)

	hash, err := cryptoUtil.MatchHash(he.Name, he.Salt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	hash.Write([]byte(he.Plaintext))

	response.OkWithData(hex.EncodeToString(hash.Sum(nil)), c)
}

func (d *CryptographyMd5Api) QysUserDe(c *gin.Context) {
	var he crypto.HashEncrypt
	_ = c.ShouldBindJSON(&he)

	one, err := cryptoUtil.MatchHash("HmacSHA256", "5f6db7ec8325a5e4")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	one.Write([]byte(he.Salt))

	he.Plaintext = strings.ReplaceAll(he.Plaintext, "{cipher}", "")

	code2Base64, err := DeAesCode2Base64(he.Plaintext, one.Sum(nil))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(string(code2Base64), c)

}

//base64解码

func DeAesCode2Base64(pwd string, key []byte) ([]byte, error) {
	//解码base64字符串
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return nil, err
	}
	//执行aes解密
	return cryptoUtil.AesDecryptECB(pwdByte, key), nil

}
