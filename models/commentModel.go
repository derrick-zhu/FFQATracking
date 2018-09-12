package models

import (
	"FFQATracking/utils"
	"fmt"
	"sort"
	"strconv"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

// CommentModel comments for issue
type CommentModel struct {
	ID        int64  `orm:"index"`
	IssueID   int64  `orm:"index"`
	Content   string `orm:"size(4096)"`
	CreatorID int64  `orm:"index"`
	Time      TimeInterval
}

func init() {
	orm.RegisterModel(new(CommentModel))
}

// TableName db table for storage
func (c *CommentModel) TableName() string {
	return tableNameByIssueID(c.IssueID)
}

func tableNameByIssueID(issueID int64) string {
	return "comment_" + strconv.FormatInt(issueID, 10)
}

// AddComment new comment for issue
func AddComment(issueID, creatorID int64, content string) (*CommentModel, error) {

	newComment := &CommentModel{
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
	o, _ := GetQuerySeterWithTable(tableNameByIssueID(issueID))

	comm := &CommentModel{ID: commentID}
	_, err := o.Delete(comm)

	return err
}

// CommentWithRange fetch comments for issue with its id, in range [low, low+count)
func CommentWithRange(issueID int64, low, count int) ([]CommentModel, error) {

	comms := []CommentModel{}

	o := GetOrmObject()
	sqlQuery := fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", tableNameByIssueID(issueID), count, low)
	rawResult := o.Raw(sqlQuery)

	_, err := rawResult.QueryRows(&comms)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	return comms, nil
}

// AllCommentsForIssue fetch all comments for issue with its id
func AllCommentsForIssue(issueID int64) ([]CommentModel, error) {
	return CommentWithRange(issueID, 0, -1)
}

// SortComment sort the comment by time
func SortCommentByTime(comms *[]CommentModel) {

	sort.Slice(comms, func(commA, commB int) bool {
		return (*comms)[commA].Time > (*comms)[commB].Time
	})
}
