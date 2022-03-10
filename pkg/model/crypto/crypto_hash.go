package crypto

import "fmt"

/*
@Date:2022/2/27
@Author [Ambi](https://github.com/AmbitionLover)

*/

//go:generate stringer -type=HashEncrypt
type HashEncrypt struct {
	Plaintext string `json:"plaintext"`
	Name      string `json:"name"`
	Mode      int    `json:"mode"`
	Salt      string `json:"salt"`
}

func (e HashEncrypt) String() string {
	return fmt.Sprintf("plaintext:%s,\t name:%s,\t mode:%d,\t salt:%s", e.Plaintext, e.Name, e.Mode, e.Salt)
}
