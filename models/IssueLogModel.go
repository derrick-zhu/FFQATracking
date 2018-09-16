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
)

// IssueLogModel comments for issue
type IssueLogModel struct {
	ID        int64  `orm:"pk;index;auto"`
	IssueID   int64  `orm:"index"`
	Content   string `orm:"size(4096)"`
	CreatorID int64
	Time      int64
	PrvStatus int64 // 老的issue状态
	NewStatus int64 // 新的issue状态
}

func init() {
	beego.Info("init()")
	orm.RegisterModel(new(IssueLogModel))
}

// TableName db table for storage
func (c *IssueLogModel) TableName() string {
	return issueLogTableName
}

// AddComment new comment for issue
func AddComment(issueID, creatorID, prvStatus, newStatus int64, content string) (*IssueLogModel, error) {

	newComment := &IssueLogModel{
		IssueID:   issueID,
		CreatorID: creatorID,
		Content:   content,
		Time:      utils.TimeTickSince1970(),

		PrvStatus: prvStatus,
		NewStatus: newStatus,
	}

	o := GetOrmObject()
	_, err := o.Insert(newComment)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	return newComment, nil
}

// RemoveComment delete comment with comment id in issue
func RemoveComment(issueID, commentID int64) error {

	o, _ := GetQuerySeterWithTable(issueLogTableName)

	comm := &IssueLogModel{ID: commentID}
	_, err := o.Delete(comm)

	return err
}

// CommentWithRange fetch comments for issue with its id, in range [low, low+count)
func CommentWithRange(issueID int64, low, count int) (*[]IssueLogModel, error) {

	comms := &[]IssueLogModel{}

	o := GetOrmObject()
	sqlQuery :=
		fmt.Sprintf("SELECT * FROM %s WHERE issue_i_d == %d LIMIT %d OFFSET %d",
			issueLogTableName,
			issueID,
			count,
			low)
	rawResult := o.Raw(sqlQuery)

	_, err := rawResult.QueryRows(comms)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	beego.Debug(comms)

	return comms, nil
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
