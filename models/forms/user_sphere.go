package forms

import (
	"bitbucket.org/SummerCampDev/summercamp/models"
	u "bitbucket.org/SummerCampDev/summercamp/models/forms/utils"
	"github.com/astaxie/beego/validation"
)

type UserSphere struct {
	FormModel
	ID       int      `json:"id"`
	Action   u.Action `json:"action" valid:"Required; Match(add|remove)"`
	UserID   int      `json:"user_id"`
	SphereID int      `json:"sphere_id"`
}

func (us *UserSphere) Valid(v *validation.Validation) {
	switch us.Action {
	case u.ActionAdd:
		v.Required(us.UserID, "user_id")
		v.Required(us.SphereID, "sphere_id")
	case u.ActionRemove:
		v.Required(us.ID, "id")
	}
}

func (us *UserSphere) Save() (*models.UserSphere, bool) {
	if ok := us.validate(us); !ok {
		return nil, false
	}

	switch us.Action {
	case u.ActionAdd:
		userSphere := models.UserSphere{
			UserID:   us.UserID,
			SphereID: us.SphereID,
		}
		if ok := userSphere.Save(); !ok {
			return nil, false
		}

		return &userSphere, true

	case u.ActionRemove:
		userSphere := models.UserSphere{
			ID: us.ID,
		}
		return nil, userSphere.Delete()
	}

	return nil, false
}
