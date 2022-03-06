package main

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download


import (
	_ "github.com/AmbitionLover/go-snake/init"
)

func main() {

}
