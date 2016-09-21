package models

import "bitbucket.org/SummerCampDev/summercamp/models/utils"

type Skill struct {
	ID   int    `json:"id" orm:"column(id)"`
	Name string `json:"name" orm:"column(name)"`
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
// which can be useful while working with TeamMember model and are not directly relate to it
type skillsAPI struct{}

// TeamMembers is an object via which we can access helper methods for the TeamMember model
var Skills *skillsAPI

// FetchByID fetch a team from the teams table by id
func (s *skillsAPI) FetchByID(id int) (*Skill, bool) {
	skill := Skill{ID: id}
	err := DB.Read(&skill)
	return &skill, utils.ProcessError(err, "fetch the skill by id")
}

// NewSkill is a wrapper to initialize a new skill object
func (s *skillsAPI) NewSkill(name string) *Skill {
	return &Skill{
		Name: name,
	}
}
