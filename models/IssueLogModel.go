package models

import (
	"FFQATracking/utils"
	"fmt"
	"sort"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

const (
	issueLogTableName string = "issuelog"

	// LogTypeComment for the comment change type of IssueLogModel
	LogTypeComment = 0
	// LogTypeStatus for the status change of IssueLogModel
	LogTypeStatus = 1
)

// IssueLogModel comments for issue
type IssueLogModel struct {
	ID          int64  `orm:"pk;index;auto"`
	IssueID     int64  `orm:"index"`
	Type        int64  // Type为LogTypeComment时， content是comment内容; LogTypeStatus时，content无效
	Content     string `orm:"size(4096)"`
	CreatorID   int64
	Time        int64
	StatusTitle string
	PrvStatus   int64 // 老的issue状态
	NewStatus   int64 // 新的issue状态
}

func init() {
	orm.RegisterModel(new(IssueLogModel))
}

// TableName db table for storage
func (c *IssueLogModel) TableName() string {
	return issueLogTableName
}

// AddLogComment new comment log for issue
func AddLogComment(issueID, creatorID int64, content string) (*IssueLogModel, error) {

	newComment := &IssueLogModel{
		IssueID:   issueID,
		CreatorID: creatorID,
		Type:      LogTypeComment,
		Content:   content,
		Time:      utils.TimeTickSince1970(),
		PrvStatus: -1,
		NewStatus: -1,
	}

	o := GetOrmObject()
	if _, err := o.Insert(newComment); err != nil {
		beego.Error(err)
		o.Rollback()
		return nil, err
	}

	return newComment, nil
}

// AddLogStatus new status log for issue
func AddLogStatus(issueID, creatorID int64, statusTitle string, prvStatus, newStatus int64) (*IssueLogModel, error) {

	newComment := &IssueLogModel{
		IssueID:     issueID,
		CreatorID:   creatorID,
		Type:        LogTypeStatus,
		Time:        utils.TimeTickSince1970(),
		StatusTitle: statusTitle,
		PrvStatus:   prvStatus,
		NewStatus:   newStatus,
	}

	o := GetOrmObject()
	if _, err := o.Insert(newComment); err != nil {
		beego.Error(err)
		o.Rollback()
		return nil, err
	}

	return newComment, nil
}

// RemoveComment delete comment with comment id in issue
func RemoveComment(issueID, commentID int64) error {

	o, _ := GetQuerySeterWithTable(issueLogTableName)

	var err error
	comm := &IssueLogModel{ID: commentID}

	if _, err = o.Delete(comm); err != nil {
		beego.Error(err)
		o.Rollback()
	}

	return err
}

// CommentWithRange fetch comments for issue with its id, in range [low, low+count)
func CommentWithRange(issueID int64, low, count int) (*[]IssueLogModel, error) {

	comms := []IssueLogModel{}

	o := GetOrmObject()
	sqlQuery :=
		fmt.Sprintf("SELECT * FROM %s WHERE issue_i_d == %d LIMIT %d OFFSET %d",
			issueLogTableName,
			issueID,
			count,
			low)
	rawResult := o.Raw(sqlQuery)

	_, err := rawResult.QueryRows(&comms)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	return &comms, nil
}

// AllCommentsForIssue fetch all comments for issue with its id
func AllCommentsForIssue(issueID int64) (*[]IssueLogModel, error) {
	return CommentWithRange(issueID, 0, -1)
}

// SortCommentByTime sort the comment by time
func SortCommentByTime(comms *[]IssueLogModel) {

	sort.Slice((*comms), func(commA, commB int) bool {
		return (*comms)[commA].Time > (*comms)[commB].Time
	})
}
