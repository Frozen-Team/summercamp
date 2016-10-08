package forms

import "bitbucket.org/SummerCampDev/summercamp/models"

type UserEmailUpdate struct {
	FormModel
	Email string `json:"email" valid:"Required; Email"`
}

func (eu *UserEmailUpdate) UpdateEmail(user *models.User) (*models.User, bool) {
	if ok := eu.validate(eu); !ok {
		return nil, false
	}
	user.Email = eu.Email
	ok := user.Save()
	if !ok {
		eu.addError("user-save-failed")
	}
	return user, ok
}