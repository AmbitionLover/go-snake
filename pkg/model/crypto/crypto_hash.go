package crypto

import "fmt"

/*
@Date:2022/2/27
@Author [Ambi](https://github.com/AmbitionLover)

*/

type HashEncrypt struct {
	Plaintext string `json:"plaintext"`
	Name      string `json:"name"`
	Mode      int    `json:"mode"`
}

func (e HashEncrypt) String() string {
	return fmt.Sprintf("plaintext:%s, name:%s, mode:%d", e.Plaintext, e.Name, e.Mode)
}
