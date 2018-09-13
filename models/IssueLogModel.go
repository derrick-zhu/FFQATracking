package models

import (
	"FFQATracking/utils"
	"fmt"
	"sort"
	"strconv"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

const (
	kIssueLogTablePrefix string = "issuelog_"
)

// IssueLogModel comments for issue
type IssueLogModel struct {
	ID        int64 `orm:"index;pk"`
	IssueID   int64
	Content   string `orm:"size(4096)"`
	CreatorID int64
	Time      TimeInterval
}

func init() {
	beego.Info("init()")
	orm.RegisterModelWithPrefix(kIssueLogTablePrefix, new(IssueLogModel))
}

// TableName db table for storage
func (c *IssueLogModel) TableName() string {
	beego.Info("TableName()")
	return tableNameByIssueID(c.IssueID)
}

func tableNameByIssueID(issueID int64) string {
	return strconv.FormatInt(issueID, 10)
}

// AddComment new comment for issue
func AddComment(issueID, creatorID int64, content string) (*IssueLogModel, error) {

	newComment := &IssueLogModel{
		IssueID:   issueID,
		CreatorID: creatorID,
		Content:   content,
		Time:      TimeInterval(utils.TimeIntervalSince1970()),
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

	o, _ := GetQuerySeterWithTable(fmt.Sprintf("%s%s", kIssueLogTablePrefix, tableNameByIssueID(issueID)))

	comm := &IssueLogModel{ID: commentID}
	_, err := o.Delete(comm)

	return err
}

// CommentWithRange fetch comments for issue with its id, in range [low, low+count)
func CommentWithRange(issueID int64, low, count int) ([]IssueLogModel, error) {

	comms := []IssueLogModel{}

	o := GetOrmObject()
	sqlQuery := fmt.Sprintf("SELECT * FROM %s%s LIMIT %d OFFSET %d", kIssueLogTablePrefix, tableNameByIssueID(issueID), count, low)
	rawResult := o.Raw(sqlQuery)

	_, err := rawResult.QueryRows(&comms)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	return comms, nil
}

// AllCommentsForIssue fetch all comments for issue with its id
func AllCommentsForIssue(issueID int64) ([]IssueLogModel, error) {
	return CommentWithRange(issueID, 0, -1)
}

// SortCommentByTime sort the comment by time
func SortCommentByTime(comms *[]IssueLogModel) {

	sort.Slice(comms, func(commA, commB int) bool {
		return (*comms)[commA].Time > (*comms)[commB].Time
	})
}
