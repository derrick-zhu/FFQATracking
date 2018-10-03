package routers

import (
	"FFQATracking/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	beego.Router("/register", &controllers.RegisterController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.AutoRouter(&controllers.LoginController{})

	beego.Router("/person/:uid", &controllers.PersonController{})
	beego.Router("/account/:uid", &controllers.AccountController{})

	beego.Router("/issuelist", &controllers.IssueListController{})

	beego.Router("/issue", &controllers.IssueNewController{})
	beego.AutoRouter(&controllers.IssueNewController{})

	beego.Router("/issuedetail/:issue(\\d+)", &controllers.IssueDetailController{})
	beego.Router("/issuedetail/:issue(\\d+)/newlog", &controllers.IssueDetailController{}, "post:SubmitNewLog")
	beego.Router("/issuedetail/:issue(\\d+)/update", &controllers.IssueDetailController{}, "post:UpdateIssue")
	beego.Router("/issuedetail/:issue(\\d+)/newattach", &controllers.IssueDetailController{}, "post:NewAttachment")

	// beego.Router("/weather", &controllers.WeatherDemoController{})
}
