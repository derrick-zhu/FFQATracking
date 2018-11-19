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
	beego.Router("/blackboard", &controllers.BlackboardController{}) // 这个通用入口恐怕可以允许数据默认置顶的project id, mile stone id等参数
	beego.Router("/blackboard/newissue", &controllers.BlackboardController{}, "post:SubmitNewIssue")
	beego.Router("/blackboard/newinitiative", &controllers.BlackboardController{}, "post:SubmitNewProject")
	beego.Router("/blackboard/newmilestone", &controllers.BlackboardController{}, "post:SubmitNewMilestone")
	// beego.Router("/blackboard/filter/change/", &controllers.BlackboardController{}, "get:FilterChanged")

	// issue详情页的控制器
	beego.Router("/issuedetail/:issue(\\d+)", &controllers.IssueDetailController{})
	beego.Router("/issuedetail/:issue(\\d+)/newlog", &controllers.IssueDetailController{}, "post:SubmitNewLog")
	beego.Router("/issuedetail/:issue(\\d+)/update", &controllers.IssueDetailController{}, "post:UpdateIssue")
	beego.Router("/issuedetail/:issue(\\d+)/newattach", &controllers.IssueDetailController{}, "post:NewAttachment")
	beego.Router("/issuedetail/:issue(\\d+)/deletecomment", &controllers.IssueDetailController{}, "post:DeleteComment")
}
