package forms

import (
	"errors"

	"bitbucket.org/SummerCampDev/summercamp/models"
	"github.com/astaxie/beego"
)

type UserReg struct {
	Email           string            `json:"email" valid:"Required; Email"`
	Type            models.Speciality `json:"type" valid:"Required;Match(executor|manager|client)"`
	FirstName       string            `json:"first_name" valid:"Required"`
	LastName        string            `json:"last_name" valid:"Required"`
	Password        string            `json:"password" valid:"Required"` //TODO: think about password restrictions
	PasswordConfirm string            `json:"password_confirm" valid:"Required"`
	Country         string            `json:"country" valid:"Required"`
	City            string            `json:"city" valid:"Required"`
	Errors          []error           `json:"-"`
}

// Register validates the input data and if everything is OK, initialize the models.User struct with
// the data from Registration struct and save the record to the db.
func (ur *UserReg) Register() (*models.User, bool) {
	errs, ok := validate(ur)
	if !ok {
		ur.Errors = errs
		return nil, false
	}

	if ur.Password != ur.PasswordConfirm {
		ur.Errors = append(ur.Errors, errors.New("passwords-mismatch"))
		beego.Warning(ur.Errors[0])
		return nil, false
	}

	user := &models.User{Email: ur.Email, Type: ur.Type, FirstName: ur.FirstName, LastName: ur.LastName, Country: ur.Country, City: ur.City}

	ok = user.SetPassword(ur.Password)
	if !ok {
		ur.Errors = append(ur.Errors, errors.New("user-password-set-failed"))
		return nil, false
	}
	ok = user.Save()
	if !ok {
		ur.Errors = append(ur.Errors, errors.New("user-save-failed"))
	}
	return user, ok
}
