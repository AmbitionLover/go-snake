package crypto

/*
@Date:2022/2/27
@Author [Ambi](https://github.com/AmbitionLover)

*/

import (
	"crypto"
	"fmt"
	hashEncrypt "github.com/AmbitionLover/go-snake/pkg/model/crypto"
	"hash"
	"strings"
)

const _Hash_name = "MD4MD5SHA1SHA224SHA256SHA384SHA512MD5SHA1RIPEMD160SHA3_224SHA3_256SHA3_384SHA3_512SHA512_224SHA512_256BLAKE2s_256BLAKE2b_256BLAKE2b_384BLAKE2b_512"

var _Hash_index = [...]uint8{0, 3, 6, 10, 16, 22, 28, 34, 41, 50, 58, 66, 74, 82, 92, 102, 113, 124, 135, 146}

func MatchHash(hashName string) (h hash.Hash, err error) {

	index := strings.Index(_Hash_name, hashName)
	for i, v := range _Hash_index {
		if int(v) == index {
			return crypto.Hash(i + 1).New(), nil
		}
	}

	return nil, fmt.Errorf("hash error by name : %s", hashName)
}

func Sum(r hashEncrypt.HashEncrypt) {

}
