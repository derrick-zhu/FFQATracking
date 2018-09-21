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

	beego.Router("/issue", &controllers.IssueController{})
	beego.AutoRouter(&controllers.IssueController{})

	beego.Router("/issuedetail/:issue(\\d+)", &controllers.IssueDetailController{})
	beego.Router("/issuedetail/:issue(\\d+)/newlog", &controllers.IssueDetailController{}, "post:SubmitNewLog")
	beego.Router("/issuedetail/:issue(\\d+)/update", &controllers.IssueDetailController{}, "post:UpdateIssue")

	beego.Router("/weatherDemo", &controllers.WeatherDemoController{})
	beego.Router("/upload", &controllers.ReadController{})
	// beego.Router("/upload/PageData", &controllers.UserController{})
	// beego.Router("/upload/PageNextData", &controllers.YonghuController{})
	// beego.Router("/upload/Index", &controllers.PageController{})
	// beego.Router("/upload/EasyUI", &controllers.EasyUIController{})
	// beego.Router("/upload/EasyUIData", &controllers.EasyUIDataController{})
	// beego.Router("/upload/FileOpt", &controllers.FileOptUploadController{})
	// beego.Router("/upload/FileDown", &controllers.FileOptDownloadController{})
	beego.Router("/upload/FileRead", &controllers.ReadController{})
	beego.Router("/upload/FileWrite", &controllers.WriteController{})
	beego.Router("/upload/FileCreate", &controllers.CreateController{})
	beego.Router("/upload/FileDelete", &controllers.DeleteController{})
}
