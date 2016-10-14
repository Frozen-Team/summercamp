package forms

import (
	"bitbucket.org/SummerCampDev/summercamp/models"
	u "bitbucket.org/SummerCampDev/summercamp/models/forms/utils"

	"github.com/astaxie/beego/validation"
)

type UserSkill struct {
	FormModel
	ID      int      `json:"id"`
	Action  u.Action `json:"action" valid:"Required; Match(add|remove)"`
	UserID  int      `json:"user_id"`
	SkillID int      `json:"skill_id"`
}

func (us *UserSkill) Valid(v *validation.Validation) {
	switch us.Action {
	case u.ActionAdd:
		v.Required(us.UserID, "user_id")
		v.Required(us.SkillID, "skill_id")
	case u.ActionRemove:
		v.Required(us.ID, "id")
	}
}

func (us *UserSkill) Save() (*models.UserSkill, bool) {
	switch us.Action {
	case u.ActionAdd:
		userSkill := models.UserSkill{
			UserID:  us.UserID,
			SkillID: us.SkillID,
		}
		if ok := userSkill.Save(); !ok {
			return nil, false
		}

		return &userSkill, true

	case u.ActionRemove:
		userSkill := models.UserSkill{
			ID: us.ID,
		}
		return nil, userSkill.Delete()
	}

	return nil, false
}
