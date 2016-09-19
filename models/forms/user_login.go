package forms

import (
	"bitbucket.org/SummerCampDev/summercamp/models"
)

type UserLogin struct {
	FormModel
	Email    string `json:"email" valid:"Required; Email"`
	Password string `json:"password" valid:"Required; MaxSize(64)"`
}

// Login validates the input data and if everything is OK, fetches a User struct by email. with
// the data from Registration struct and save the record to the db.
// Return table:
// User	         bool    Description
// nil or User   false   Error occurred
// nil           true    Email and\or password is incorrect
// User          true    The User is logged on
//
func (ul *UserLogin) Login() (*models.User, bool) {
	if ok := ul.validate(ul); !ok {
		return nil, false
	}

	user, ok := models.Users.FetchByEmail(ul.Email)
	if !ok {
		return nil, true
	}

	match, ok := user.CheckPassword(ul.Password)
	if !ok {
		ul.addError("user-login-failed")
		return nil, false
	}
	if !match {
		return nil, true
	}

	return user, true
}
