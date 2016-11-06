package models

import (
	"bitbucket.org/SummerCampDev/summercamp/models/utils"
	"github.com/astaxie/beego"
)

type UserSkill struct {
	ID      int `json:"id" orm:"column(id)"`
	UserID  int `json:"user_id" orm:"column(user_id)"`
	SkillID int `json:"skill_id" orm:"column(skill_id)"`
}

// TableName specify the table name for UserSkill model. This name is used in the orm RegisterModel
func (us *UserSkill) TableName() string {
	return "user_skills"
}

// Save inserts a new or updates an existing user's skill record in the DB.
func (us *UserSkill) Save() bool {
	_, err := DB.InsertOrUpdate(us)
	return utils.ProcessError(err, "insert or update user's skill")
}

// Delete deletes a record from the db. If the record is successfully deleted, the return value
// is true, false - otherwise.
func (us *UserSkill) Delete() bool {
	if us.ID == 0 {
		return false
	}
	_, err := DB.Delete(us)

	return utils.ProcessError(err, " delete a user`s skill")
}

type UserSkillsAPI struct{}

var UserSkills *UserSkillsAPI

// SaveSkillsForUser create a new UserSkill record for each skillID from skillIDs and userID pair.
// If each record is successfully saved to the db, the func return false
func (us *UserSkillsAPI) SaveSkillsForUser(userID int, skillIDs ...int) bool {
	if len(skillIDs) == 0 {
		beego.BeeLogger.Warning("Empty skills list is passed to SaveSkillsForUser")
		return false
	}

	var failedSkills []int
	for _, skillID := range skillIDs {
		userSkill := UserSkill{
			UserID:  userID,
			SkillID: skillID,
		}
		if ok := userSkill.Save(); !ok {
			failedSkills = append(failedSkills, skillID)
		}
	}
	ok := len(failedSkills) == 0
	if !ok {
		beego.BeeLogger.Warning("Failed to save user skills for skills with ids: '%v'", failedSkills)
	}
	return ok
}

// FetchSkillsByUser fetch all skills for a given user
func (us *UserSkillsAPI) FetchSkillsByUser(userID int) ([]Skill, bool) {
	var skills []Skill
	_, err := DB.Raw(`
	SELECT skills.id,
	       skills.name,
	       skills.sphere_id
	FROM user_skills us
	LEFT OUTER JOIN skills ON skills.id=us.skill_id
	WHERE us.user_id=$1;`, userID).QueryRows(&skills)
	return skills, utils.ProcessError(err, " fetch skills by a user id")
}

// FetchUsersBySkills fetch all users for a given skill id
func (us *UserSkillsAPI) FetchUsersBySkill(skillID int) ([]User, bool) {
	var users []User
	_, err := DB.Raw(`
	SELECT users.id,
	       users.type,
	       users.first_name,
	       users.last_name,
	       users.balance,
	       users.bid,
	       users.braintree_id,
	       users.country,
	       users.city,
	       users.timezone,
	       users.create_time,
	       users.update_time
	FROM user_skills us
	LEFT OUTER JOIN users ON users.id=us.user_id
	WHERE us.skill_id=$1;`, skillID).QueryRows(&users)
	return users, utils.ProcessError(err, " fetch users by a skill id")
}

func (us *UserSkillsAPI) SkillsCountByUser(userID int) (int, bool) {
	count, err := DB.QueryTable(UserSkillModel).Filter("user_id", userID).Count()
	return int(count), utils.ProcessError(err, " fetch skills count for user")
}
