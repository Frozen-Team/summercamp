package forms

import (
	"bitbucket.org/SummerCampDev/summercamp/models"
	"github.com/astaxie/beego/validation"
)

type action string

const (
	UserSphereAdd    action = "add"
	UserSphereRemove action = "remove"
)

type UserSphere struct {
	FormModel
	ID       int    `json:"id"`
	Action   action `json:"action" valid:"Required; Match(add|remove)"`
	UserID   int    `json:"user_id"`
	SphereID int    `json:"sphere_id"`
}

func (us *UserSphere) Valid(v *validation.Validation) {
	switch us.Action {
	case UserSphereAdd:
		v.Required(us.UserID, "user_id")
		v.Required(us.SphereID, "sphere_id")
	case UserSphereRemove:
		v.Required(us.ID, "id")
	}
}

func (us *UserSphere) Save() (*models.UserSphere, bool) {
	if ok := us.validate(us); !ok {
		return nil, false
	}

	switch us.Action {
	case UserSphereAdd:
		userSphere := models.UserSphere{
			UserID:   us.UserID,
			SphereID: us.SphereID,
		}
		if ok := userSphere.Save(); !ok {
			return nil, false
		}

		return &userSphere, true

	case UserSphereRemove:
		userSphere := models.UserSphere{
			ID: us.ID,
		}
		return nil, userSphere.Delete()
	}

	return nil, false
}
