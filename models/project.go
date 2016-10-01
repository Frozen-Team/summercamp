package models

import (
	"time"

	"bitbucket.org/SummerCampDev/summercamp/models/utils"
)

type Project struct {
	ID          int       `orm:"column(id)"`
	Description string    `orm:"column(description)"`
	Budget      int       `orm:"column(budget)"`
	ClientID    int       `orm:"column(client_id)"`
	CreateTime  time.Time `orm:"column(create_time);auto_now_add;type(datetime)"`
	UpdateTime  time.Time `orm:"column(update_time);auto_now;type(datetime)"`
}

// TableName specifies the table name of the model Project
func (p *Project) TableName() string {
	return "projects"
}

// Save creates a project record in the db
func (p *Project) Save() bool {
	var err error
	var action string

	if p.ID == 0 {
		_, err = DB.Insert(p)
		action = "create"
	} else {
		_, err = DB.Update(p)
		action = "update"
	}
	return utils.ProcessError(err, action+" project")
}

// Delete deletes the project record from the db
func (p *Project) Delete() bool {
	_, err := DB.Delete(p)
	return utils.ProcessError(err, " delete project")
}

// Skills is a wrapper for a method of ProjectSkillsAPI to fetch skills by a project id.
func (p *Project) Skills() ([]Skill, bool) {
	return ProjectSkills.FetchSkillsByProject(p.ID)
}

type projectsAPI struct{}

var Projects *projectsAPI

// FetchByID fetches a project from the projects table by id.
func (t *projectsAPI) FetchByID(id int) (*Project, bool) {
	project := Project{ID: id}
	err := DB.Read(&project)
	return &project, utils.ProcessError(err, " fetch the project by id")
}

// FetchAllByClient fetches all projects for the given user.
func (p *projectsAPI) FetchAllByClient(clientID int) ([]Project, bool) {
	var projects []Project
	_, err := DB.QueryTable(ProjectModel).Filter("client_id", clientID).All(&projects)
	return projects, utils.ProcessError(err, " fetch projects by client")
}
