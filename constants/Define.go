package constants

const (
	Title       = "Title"
	IsLogin     = "IsLogin"
	AccountData = "AccountData"

	/// login & signup

	KeyUID         = "uid"
	KeyUNAME       = "uname"
	KeyEMAIL       = "email"
	KeyPWD         = "pwd"
	KeyPermissions = "PersmissionRules"

	/// new issue

	KeyIssueHTMLValue   = "issueHTMLValue"
	KeyIssueAttachments = "issueAttachments"

	/// issue detail

	KeyIssueData       = "issueData"
	KeyIssueLogHistory = "issueLogHistory"

	// KeyIsHome flag for show HOME page.
	KeyIsHome       = "IsHome"
	KeyIsBlackBoard = "IsBlackBoard"
)

const (
	// ServerUploadDir local upload dir
	ServerUploadDir = "static/upload/"
)

const (
	// MAXINT max integer number
	MAXINT = 1<<32 - 1
)
