package routers

import (
	"trans/controllers"
	beego "github.com/astaxie/beego/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
