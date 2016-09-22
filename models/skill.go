package models

import (
	"bitbucket.org/SummerCampDev/summercamp/models/utils"
	"github.com/astaxie/beego"
)

type Skill struct {
	ID       int    `json:"id" orm:"column(id)"`
	SphereID int    `json:"sphere_id" orm:"column(sphere_id)"`
	Name     string `json:"name" orm:"column(name)"`
}

func (s *Skill) TableName() string {
	return "skills"
}

// Save insert a new record to the db if ID field is of default value. Otherwise an existing
// record is updated.
func (s *Skill) Save() bool {
	var err error
	var action string

	if s.ID == 0 {
		_, err = DB.Insert(s)
		action = "create"
	} else {
		_, err = DB.Update(s)
		action = "update"
	}

	return utils.ProcessError(err, action+" skill")
}

// Delete deletes a record from the db. If the record is successfully deleted, the return value
// is true, false - otherwise.
func (s *Skill) Delete() bool {
	if s.ID == 0 {
		return false
	}
	_, err := DB.Delete(s)

	return utils.ProcessError(err, " delete skill")
}

// skillsAPI is an empty struct which is a receiver of helper methods
// which can be useful while working with Skill model.
type skillsAPI struct{}

// Skills is an object via which we can access helper methods for the Skill model.
var Skills *skillsAPI

// FetchByID fetch a a skill from the skills table by id
func (s *skillsAPI) FetchByID(id int) (*Skill, bool) {
	skill := Skill{ID: id}
	err := DB.Read(&skill)
	return &skill, utils.ProcessError(err, "fetch the skill by id")
}

// NewSkill is a wrapper to initialize a new skill object.
func (s *skillsAPI) NewSkill(name string, sphereID int) *Skill {
	return &Skill{
		Name:     name,
		SphereID: sphereID,
	}
}

// FetchAllByNames fetch skills by their names. For some possible use, the 'name' field
// of the model is also initialized. If some of the skills are not present in the db, the method
// still returns true and those skills, which were successfully fetched.
func (s *skillsAPI) FetchAllByNames(skillNames ...string) ([]Skill, bool) {
	if len(skillNames) == 0 {
		beego.Warning("Empty skill names are passed to FetchIDsByNames")
		return nil, false
	}

	var skills []Skill

	_, err := DB.QueryTable(SkillModel).Filter("name__in", skillNames).All(&skills)
	if err != nil {
		return nil, utils.ProcessError(err, " fetch skills by names")
	}

	if len(skills) != len(skillNames) {
		beego.BeeLogger.Warning("the following skills '%#v' cannot be fetched from the db",
			findMissing(skills, skillNames))
	}

	return skills, true
}

// FetchAllBySphere fetch skills by a sphere id specified as sphereID.
func (s *skillsAPI) FetchAllBySphere(sphereID int) ([]Skill, bool) {
	var skills []Skill
	_, err := DB.QueryTable(SkillModel).Filter("sphere_id", sphereID).All(&skills)
	return skills, utils.ProcessError(err, " fetch skills by sphere")
}

// findMissing find that skill names from skillNames which are not in skills.
func findMissing(skills []Skill, skillNames []string) []string {
	var results []string
	for _, skillName := range skillNames {
		for _, skill := range skills {
			if skillName == skill.Name {
				break
			}
		}
		results = append(results, skillName)
	}
	return results
}
