package forms

import "bitbucket.org/SummerCampDev/summercamp/models"

type UserSummary struct {
	FormModel
	Summary string `json:"summary" valid:"Required; MaxSize(64)"` //TODO: decide constraint on the summary
}

func (us *UserSummary) Update(user *models.User) bool {
	if !us.validate(us) {
		return false
	}

	user.Summary = us.Summary

	ok := user.Save()
	if !ok {
		us.addError("user-update-failed")
	}

	return ok
}
