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

	beego.Router("/issuedetail/:issue(\\d+)", &controllers.IssueDetailController{})
	beego.Router("/issuedetail/:issue(\\d+)/newlog", &controllers.IssueDetailController{}, "post:SubmitNewLog")
	beego.Router("/issuedetail/:issue(\\d+)/update", &controllers.IssueDetailController{}, "post:UpdateIssue")

	beego.Router("/weatherDemo", &controllers.WeatherDemoController{})
}
