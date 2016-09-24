package models

import (
	"bitbucket.org/SummerCampDev/summercamp/models/utils"
	"github.com/astaxie/beego"
)

type ProjectSkill struct {
	ID        int `json:"id" orm:"column(id)"`
	ProjectID int `json:"project_id" orm:"column(project_id)"`
	SkillID   int `json:"skill_id" orm:"column(column_id)"`
}

// TableName specify the table name for ProjectSkill model. This name is used in the orm RegisterModel
func (ps *ProjectSkill) TableName() string {
	return "project_skills"
}

// Save insert a new record to the db if ID field is of default value. Otherwise an existing
// record is updated.
func (ps *ProjectSkill) Save() bool {
	var err error
	var action string

	if ps.ID == 0 {
		_, err = DB.Insert(ps)
		action = "create"
	} else {
		_, err = DB.Update(ps)
		action = "update"
	}

	return utils.ProcessError(err, action+" a project`s skill")
}

// Delete deletes a record from the db. If the record is successfully deleted, the return value
// is true, false - otherwise.
func (ps *ProjectSkill) Delete() bool {
	if ps.ID == 0 {
		return false
	}
	_, err := DB.Delete(ps)

	return utils.ProcessError(err, " delete a project`s skill")
}

type ProjectSkillsAPI struct{}

var ProjectSkills *ProjectSkillsAPI

// SaveSkillsForProject create a new ProjectSkill record for each skillID from skillIDs and projectID pair.
// If each record is successfully saved to the db, the func return false
func (ps *ProjectSkillsAPI) SaveSkillsForProject(projectID int, skillIDs ...int) bool {
	if len(skillIDs) == 0 {
		beego.BeeLogger.Warning("Empty skills list is passed to SaveSkillsForProject")
		return false
	}

	var failedSkills []int
	for _, skillID := range skillIDs {
		projectSkill := ProjectSkill{
			ProjectID: projectID,
			SkillID:   skillID,
		}
		if ok := projectSkill.Save(); !ok {
			failedSkills = append(failedSkills, skillID)
		}
	}
	ok := len(failedSkills) == 0
	if !ok {
		beego.BeeLogger.Warning("Failed to save project skills for skills with ids: '%v'", failedSkills)
	}
	return ok
}
