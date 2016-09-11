package forms

import (
	"unicode"

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
	Errors          []string          `json:"-"`
}

// Register validates the input data and if everything is OK, initialize the models.User struct with
// the data from Registration struct and save the record to the db.
func (ur *UserReg) Register() (*models.User, bool) {
	errs, ok := validate(ur)
	if !ok {
		ur.Errors = errs
		return nil, false
	}

	if !isStrongPass(ur.Password) {
		ur.Errors = append(ur.Errors, "password-weak")
		return nil, false
	}

	if ur.Password != ur.PasswordConfirm {
		ur.Errors = append(ur.Errors, "passwords-mismatch")
		beego.Warning(ur.Errors[0])
		return nil, false
	}

	user := &models.User{
		Email:     ur.Email,
		Type:      ur.Type,
		FirstName: ur.FirstName,
		LastName:  ur.LastName,
		Country:   ur.Country,
		City:      ur.City,
	}

	ok = user.SetPassword(ur.Password)
	if !ok {
		ur.Errors = append(ur.Errors, "user-password-set-failed")
		return nil, false
	}
	ok = user.Save()
	if !ok {
		ur.Errors = append(ur.Errors, "user-save-failed")
	}
	return user, ok
}

// isStringPass checks the password if it is strong enough. The minimum length is 5 symbols.
// The password must contain at least 2 of the following: upper letter, lower letter or number OR it must contain
// some special symbol which is not any of the described before.
func isStrongPass(p string) bool {
	len := len(p)
	var num, upper, lower byte
	if len < 5 {
		return false
	}
	for _, c := range p {
		if unicode.IsNumber(c) {
			num = 1
			continue
		} else if unicode.IsUpper(c) {
			upper = 1
			continue
		}
		if unicode.IsLower(c) {
			lower = 1
			continue
		}
		if !(unicode.IsLetter(c) || unicode.IsNumber(c)) {
			num = 2
			break
		}
	}
	return (num + upper + lower) >= 2
}
