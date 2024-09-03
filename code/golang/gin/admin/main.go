package main

import (
	_ "go.uber.org/automaxprocs"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	//global.GVA_DB = db.Gorm()
	//fmt.Println(global.GVA_DB)

}
