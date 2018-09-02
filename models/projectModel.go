package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	ProjectDBTable string = "ProjectModel"
)

type ProjectType int64

// ProjectModel class for bug list project
type ProjectModel struct {
	ID      IndexType `orm:"index"`
	Title   string    `orm:"index;size(128)"`
	Creator IndexType
	Owner   IndexType
	Create  time.Time
	Type    ProjectType `orm:"index"`
}

func init() {
	orm.RegisterModel(new(ProjectModel))
}

func (c *ProjectModel) TableName() string {
	return ProjectDBTable
}

func (c *ProjectModel) AddProject(title string, creator IndexType, prjType ProjectType) {
	return
}

func (c *ProjectModel) ProjectWithID(id IndexType) (*ProjectModel, error) {
	return nil, nil
}

func (c *ProjectModel) UpdateProject(id IndexType, params map[string]interface{}) error {
	return nil
}

func (c *ProjectModel) DeleteProject(id IndexType) error {
	return nil
}
