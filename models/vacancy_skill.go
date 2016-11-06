package models

import (
	"bitbucket.org/SummerCampDev/summercamp/models/utils"
	"github.com/astaxie/beego"
)

type VacancySkill struct {
	ID        int `json:"id" orm:"column(id)"`
	VacancyID int `json:"vacancy_id" orm:"column(vacancy_id)"`
	SkillID   int `json:"skill_id" orm:"column(skill_id)"`
}

// TableName specify the table name for VacancySkill model. This name is used in the orm RegisterModel
func (ps *VacancySkill) TableName() string {
	return "vacancy_skills"
}

// Save insert a new record to the db if ID field is of default value. Otherwise an existing
// record is updated.
func (ps *VacancySkill) Save() bool {
	var err error
	var action string

	if ps.ID == 0 {
		_, err = DB.Insert(ps)
		action = "create"
	} else {
		_, err = DB.Update(ps)
		action = "update"
	}
	return utils.ProcessError(err, action+" a vacancy`s skill")
}

// Delete deletes a record from the db. If the record is successfully deleted, the return value
// is true, false - otherwise.
func (ps *VacancySkill) Delete() bool {
	if ps.ID == 0 {
		return false
	}
	_, err := DB.Delete(ps)

	return utils.ProcessError(err, " delete a vacancy`s skill")
}

type VacancySkillsAPI struct{}

var VacancySkills *VacancySkillsAPI

// SaveSkillsForVacancy create a new VacancySkill record for each skillID from skillIDs and vacancyID pair.
// If the DB inserter successfully closes, the function returns true, false - otherwise.
func (ps *VacancySkillsAPI) SaveSkillsForVacancy(vacancyID int, skillIDs ...int) bool {
	if len(skillIDs) == 0 {
		beego.BeeLogger.Warning("Empty skills list is passed to SaveSkillsForVacancy")
		return false
	}

	i, err := DB.QueryTable(VacancySkillModel).PrepareInsert()
	if err != nil {
		return utils.ProcessError(err, " create an inserter")
	}

	var failedSkills []int

	for _, skillID := range skillIDs {
		_, err = i.Insert(&VacancySkill{
			VacancyID: vacancyID,
			SkillID:   skillID,
		})
		if err != nil {
			failedSkills = append(failedSkills, skillID)
		}
	}

	ok := len(failedSkills) == 0
	if !ok {
		beego.BeeLogger.Warning("Failed to save vacancy skills for skills with ids: '%v'", failedSkills)
	}

	err = i.Close()
	return utils.ProcessError(err, " insert multiple vacancy skills")
}
