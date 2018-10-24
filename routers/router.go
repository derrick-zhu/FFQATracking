package routers

import (
	"FFQATracking/controllers"

	"github.com/astaxie/beego"
)

func init() {

	// router to main page
	beego.Router("/", &controllers.MainController{})

	// router to register account
	beego.Router("/register", &controllers.RegisterController{})
	// router for helping user to sign in
	beego.Router("/login", &controllers.LoginController{})
	beego.AutoRouter(&controllers.LoginController{})

	// router for personal page or account page
	beego.Router("/person/:uid", &controllers.PersonController{})
	beego.Router("/account/:uid", &controllers.AccountController{})

	// router to blackboard, control center
	beego.Router("/blackboard", &controllers.BlackboardController{})
	beego.Router("/blackboard/newissue", &controllers.BlackboardController{}, "post:SubmitNewIssue")
	beego.Router("/blackboard/newinitiative", &controllers.BlackboardController{}, "post:SubmitNewProject")

	// // router for helping user to create new issue
	// beego.Router("/issue", &controllers.IssueNewController{})
	// beego.AutoRouter(&controllers.IssueNewController{})

	// 工程相关的控制器
	// beego.Router("/initiative", &controllers.InitiativeNewController{})
	// beego.Router("/initiative/new", &controllers.InitiativeNewController{}, "post:SubmitNewProject")

	// issue详情页的控制器
	beego.Router("/issuedetail/:issue(\\d+)", &controllers.IssueDetailController{})
	beego.Router("/issuedetail/:issue(\\d+)/newlog", &controllers.IssueDetailController{}, "post:SubmitNewLog")
	beego.Router("/issuedetail/:issue(\\d+)/update", &controllers.IssueDetailController{}, "post:UpdateIssue")
	beego.Router("/issuedetail/:issue(\\d+)/newattach", &controllers.IssueDetailController{}, "post:NewAttachment")
	beego.Router("/issuedetail/:issue(\\d+)/deletecomment", &controllers.IssueDetailController{}, "post:DeleteComment")
}
