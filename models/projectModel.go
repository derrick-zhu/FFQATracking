package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	projectDBTable string = "ProjectModel"
)

type ProjectType int64

// ProjectModel class for bug list project
type ProjectModel struct {
	ID      int64  `orm:"index;pk;auto"`
	Title   string `orm:"index;size(128)"`
	Creator int64
	Owner   int64
	Create  time.Time
	Type    ProjectType `orm:"index"`
}

func init() {
	orm.RegisterModel(new(ProjectModel))
}

func (c *ProjectModel) TableName() string {
	return projectDBTable
}

func (c *ProjectModel) AddProject(title string, creator int64, prjType ProjectType) {
	return
}

func (c *ProjectModel) ProjectWithID(id int64) (*ProjectModel, error) {
	return nil, nil
}

func (c *ProjectModel) UpdateProject(id int64, params map[string]interface{}) error {
	return nil
}

func (c *ProjectModel) DeleteProject(id int64) error {
	return nil
}
