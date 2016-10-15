package forms

import "bitbucket.org/SummerCampDev/summercamp/models"

type UserSkill struct {
	FormModel
	UserID  int `json:"user_id" valid:"Required"`
	SkillID int `json:"skill_id" valid:"Required"`
}

func (us *UserSkill) Process() (*models.UserSkill, bool) {
	userSkill := models.UserSkill{
		UserID:  us.UserID,
		SkillID: us.SkillID,
	}
	if ok := userSkill.Save(); !ok {
		return nil, false
	}

	return &userSkill, true
}
