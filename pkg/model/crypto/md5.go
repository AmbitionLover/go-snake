package crypto

/*
@Date:2022/2/18
@Author [Ambi](https://github.com/AmbitionLover)

*/
import "fmt"

// EncMode 加密模式
type EncMode int

const (
	Bit16Small EncMode = iota
	Bit16Big
	Bit32Small
	Bit32Big
)

type Md5Encrypt struct {
	Plaintext string  `json:"plaintext"`
	Mode      EncMode `json:"mode"`
}

func (e Md5Encrypt) String() string {
	return fmt.Sprintf("data:%s, mode:%d", e.Plaintext, e.Mode)
}
