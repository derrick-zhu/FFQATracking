package models

// AttachmentModel for bug's attachment files
type AttachmentModel struct {
	ID         int64  `orm:"auto;index"`
	IssueID    int64  `orm:"index"`            // for issue main body
	IssueLogID int64  `orm:"index"`            // for issue's log history item
	FileName   string `orm:"size(256);nonull"` // encoded file's name in back-end
}

// NewAttachmentFile new attachment file's info
func NewAttachmentFile(file string) string {
	return ""
}

// AttachmentForIssue fetch all attachements for the issue with its ID
func AttachmentForIssue(iid int64) (*AttachmentModel, error) {
	return nil, nil
}

// AttachmentForIssueLog fetch all attachments for issue's log with its issue-log ID
func AttachmentForIssueLog(iid, ilid int64) (*AttachmentModel, error) {
	return nil, nil
}

// DeleteAttachments delete all attachments for issue. (Project's bug could be deleted, so its attachments should be deleted too)
func DeleteAttachments(iid int64) error {
	return nil
}
