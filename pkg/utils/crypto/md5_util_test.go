package crypto

import (
	"github.com/AmbitionLover/go-snake/pkg/model/crypto"
	"github.com/AmbitionLover/go-snake/pkg/utils"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/*
@Date:2022/2/18
@Author [Ambi](https://github.com/AmbitionLover)

*/

func Test_Md5(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Test md5", t, func() {
		encrypt := crypto.Md5Encrypt{Plaintext: "123123", Mode: 0}

		Convey("16Small", func() {
			c := crypto.Md5Encrypt{}
			utils.StructAssign(&c, &encrypt)
			So(Md5De(c), ShouldEqual, "13955235245b2497")
		})
		Convey("16Big", func() {
			c := crypto.Md5Encrypt{}
			utils.StructAssign(&c, &encrypt)
			c.Mode = 1
			So(Md5De(c), ShouldEqual, "13955235245B2497")
		})
		Convey("32Small", func() {
			c := crypto.Md5Encrypt{}
			utils.StructAssign(&c, &encrypt)
			c.Mode = 2
			So(Md5De(c), ShouldEqual, "4297f44b13955235245b2497399d7a93")
		})
		Convey("32Big", func() {
			c := crypto.Md5Encrypt{}
			utils.StructAssign(&c, &encrypt)
			c.Mode = 3
			So(Md5De(c), ShouldEqual, "4297F44B13955235245B2497399D7A93")
		})
	})
}

func Test_1(t *testing.T) {
	encrypt := crypto.Md5Encrypt{Plaintext: "123123", Mode: 0}
	want := "13955235245b2497"
	if Md5De(encrypt) == want {
		t.Log("ok")
	}
	//var hashType crypto1.Hash
	//of := reflect.ValueOf(hashType)
	//t.Log(of)

}
