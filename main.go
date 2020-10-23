package main

import (
	_ "beego-user/utils"
	_ "beego-user/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
