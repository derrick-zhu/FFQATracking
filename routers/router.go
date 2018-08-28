package routers

import (
	"FFQATracking/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	beego.Router("/login", &controllers.LoginController{})
	//beego.AutoRouter(&controllers.LoginController{})
}
