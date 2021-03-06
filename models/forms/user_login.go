package forms

import (
	"github.com/Frozen-Team/summercamp/models"
)

type UserLogin struct {
	FormModel
	Email    string `json:"email" valid:"Required; Email"`
	Password string `json:"password" valid:"Required; MaxSize(64)"`
}

// Login validates the input data and if everything is OK, fetches a User struct by email. with
// the data from Registration struct and save the record to the db.
func (ul *UserLogin) LogIn() (*models.User, bool) {
	user, ok := models.Users.FetchByEmail(ul.Email)
	if !ok {
		ul.addError("user-login-or-password-incorrect")
		return nil, false
	}

	if ok := user.CheckPassword(ul.Password); !ok {
		ul.addError("user-login-or-password-incorrect")
		return nil, false
	}

	return user, true
}
