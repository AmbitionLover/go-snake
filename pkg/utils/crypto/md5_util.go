package crypto

import (
	"crypto/md5"
	"fmt"
	"github.com/AmbitionLover/go-snake/pkg/model/crypto"
	"strings"
)

/*
@Date:2022/2/18
@Author [Ambi](https://github.com/AmbitionLover)

*/

func Md5De(encrypt crypto.Md5Encrypt) (res string) {
	fmt.Printf("%s", encrypt)
	original := fmt.Sprintf("%x", md5.Sum([]byte(encrypt.Plaintext)))
	switch encrypt.Mode {
	case crypto.Bit16Small:
		res = strings.ToLower(original[8:24])
	case crypto.Bit16Big:
		res = strings.ToUpper(original[8:24])
	case crypto.Bit32Small:
		res = strings.ToLower(original)
	case crypto.Bit32Big:
		res = strings.ToUpper(original)
	}
	return
}
