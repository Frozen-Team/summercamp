package forms

import "github.com/Frozen-Team/summercamp/models"

type UserSphere struct {
	FormModel
	UserID   int `json:"user_id" valid:"Required"`
	SphereID int `json:"sphere_id" valid:"Required"`
}

func (us *UserSphere) Save() (*models.UserSphere, bool) {
	userSphere := models.UserSphere{
		UserID:   us.UserID,
		SphereID: us.SphereID,
	}
	if ok := userSphere.Save(); !ok {
		return nil, false
	}

	return &userSphere, true
}
