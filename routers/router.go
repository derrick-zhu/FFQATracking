package routers

import (
	"FFQATracking/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.AutoRouter(&controllers.LoginController{})

	beego.Router("/register", &controllers.RegisterController{})

	beego.Router("/issuelist", &controllers.IssueListController{})

	beego.Router("/issue", &controllers.IssueController{})
	beego.AutoRouter(&controllers.IssueController{})

	beego.Router("/weatherDemo", &controllers.WeatherDemoController{})
}
