package main

import (
	_ "imgclasses/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.SetStaticPath("/imgdata", "DATA/data")
	beego.Run()
}
