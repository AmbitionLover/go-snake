package crypto

import (
	"encoding/base64"
	"encoding/hex"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/*
@Date:2022/3/8
@Author [Ambi](https://github.com/AmbitionLover)

*/

func TestAesEncrypt(t *testing.T) {
	Convey("Test Aes", t, func() {
		origDataString := "Hello World"
		origData := []byte(origDataString) // 待加密的数据
		key := []byte("ABCDEFGHIJKLMNOP")  // 加密的密钥
		t.Log("原文：", string(origData))

		Convey("Aes CBC", func() {
			encrypted := AesEncryptCBC(origData, key)
			Convey("密文(hex)：", func() {
				So(hex.EncodeToString(encrypted), ShouldEqual, "20c7cbfe27e4baf919c06fefd9b9fa07")
			})
			Convey("密文(base64)：", func() {
				So(base64.StdEncoding.EncodeToString(encrypted), ShouldEqual, "IMfL/ifkuvkZwG/v2bn6Bw==")
			})
			decrypted := AesDecryptCBC(encrypted, key)
			Convey("解密结果", func() {
				So(string(decrypted), ShouldEqual, origDataString)
			})
		})

		Convey("Aes ECB", func() {
			encrypted := AesEncryptECB(origData, key)
			Convey("密文(hex)：", func() {
				So(hex.EncodeToString(encrypted), ShouldEqual, "dc6e0e53ad9d33dcd78646d7cbca66b8")
			})
			Convey("密文(base64)：", func() {
				So(base64.StdEncoding.EncodeToString(encrypted), ShouldEqual, "3G4OU62dM9zXhkbXy8pmuA==")
			})
			decrypted := AesDecryptECB(encrypted, key)
			Convey("解密结果", func() {
				So(string(decrypted), ShouldEqual, origDataString)
			})
		})

		Convey("Aes CFB", func() {
			encrypted := AesEncryptCFB(origData, key)
			t.Log("Aes CFB 密文(hex)：", hex.EncodeToString(encrypted))
			t.Log("Aes CFB 密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
			decrypted := AesDecryptCFB(encrypted, key)
			Convey("解密结果", func() {
				So(string(decrypted), ShouldEqual, origDataString)
			})
		})

	})
}
