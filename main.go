package main

import (
	_ "trans/routers"
	beego "github.com/astaxie/beego/server/web"
)

func main() {
	beego.Run()
}

