package routers

import (
	"acheme/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/reception/login/judge", &controllers.MainController{},"post:Login")
}
