package models

import (
	"github.com/Frozen-Team/summercamp/models/utils"
	"github.com/astaxie/beego"
)

type ProjectSkill struct {
	ID        int `json:"id" orm:"column(id)"`
	ProjectID int `json:"project_id" orm:"column(project_id)"`
	SkillID   int `json:"skill_id" orm:"column(skill_id)"`
}

// TableName specify the table name for ProjectSkill model. This name is used in the orm RegisterModel
func (ps *ProjectSkill) TableName() string {
	return "project_skills"
}

// Save inserts a new or updates an existing project's skill record in the DB.
func (ps *ProjectSkill) Save() bool {
	_, err := DB.InsertOrUpdate(ps)
	return utils.ProcessError(err, "insert or update project's skill")
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

// FetchSkillsByProject fetch all skills for a given project
func (ps *ProjectSkillsAPI) FetchSkillsByProject(projectID int) ([]Skill, bool) {
	var skills []Skill
	_, err := DB.Raw(`
	SELECT skills.id,
	       skills.name,
	       skills.sphere_id
	FROM project_skills ps
	LEFT OUTER JOIN skills ON skills.id=ps.skill_id
	WHERE ps.project_id=$1;`, projectID).QueryRows(&skills)
	return skills, utils.ProcessError(err, " fetch skills by a project id")
}

// FetchProjectsBySkills fetch all projects for a given skill id
func (ps *ProjectSkillsAPI) FetchProjectsBySkill(skillID int) ([]Project, bool) {
	var projects []Project
	_, err := DB.Raw(`
	SELECT projects.id,
	       projects.description,
	       projects.budget,
	       projects.client_id,
	       projects.create_time,
	       projects.update_time
	FROM project_skills ps
	LEFT OUTER JOIN projects ON projects.id=ps.project_id
	WHERE ps.skill_id=$1;`, skillID).QueryRows(&projects)
	return projects, utils.ProcessError(err, " fetch projects by a skill id")
}
