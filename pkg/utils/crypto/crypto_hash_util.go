package crypto

/*
@Date:2022/2/27
@Author [Ambi](https://github.com/AmbitionLover)

*/

import (
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"strconv"
	"strings"
)

const _Hash_name = "MD4MD5SHA1SHA224SHA256SHA384SHA512MD5SHA1RIPEMD160SHA3_224SHA3_256SHA3_384SHA3_512SHA512_224SHA512_256BLAKE2s_256BLAKE2b_256BLAKE2b_384BLAKE2b_512"

var _Hash_index = [...]uint8{0, 3, 6, 10, 16, 22, 28, 34, 41, 50, 58, 66, 74, 82, 92, 102, 113, 124, 135, 146}

func MatchHash(hashName, salt string) (h hash.Hash, err error) {

	if strings.HasPrefix(hashName, "Hmac") {
		if strings.HasPrefix(hashName, "HmacSHA") {
			hacHashNo := hashName[len("HmacSHA"):]
			atomNo, err := strconv.Atoi(hacHashNo)
			if err != nil {
				return nil, err
			}
			switch atomNo {
			case 1:
				h := hmac.New(sha256.New, []byte(salt))
				return h, nil
			case 256:
				h := hmac.New(sha256.New, []byte(salt))
				return h, nil
			}

		}
		return nil, fmt.Errorf("hash error by HmacSHAname : %s", hashName)
	}

	index := strings.Index(_Hash_name, hashName)
	for i, v := range _Hash_index {
		if int(v) == index {
			h := crypto.Hash(i + 1)
			if h.Available() {
				return h.New(), nil
			}
			break
		}
	}

	return nil, fmt.Errorf("hash error by name : %s", hashName)
}

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
