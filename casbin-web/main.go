package main

import (
	"github.com/astaxie/beego"
	_ "zct/casbin-web/routers"
)

func main() {
	beego.Run()
}

