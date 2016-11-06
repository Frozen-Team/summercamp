package forms

import (
	"bitbucket.org/SummerCampDev/summercamp/models"
	"github.com/astaxie/beego/validation"
)

type UserPasswordUpdate struct {
	FormModel
	CurrentPassword string `json:"current_password" valid:"Required; MaxSize(64)"`
	Password        string `json:"password" valid:"Required; MaxSize(64)"`
	PasswordConfirm string `json:"password_confirm" valid:"Required; MaxSize(64)"`
}

func (ur *UserPasswordUpdate) Valid(v *validation.Validation) {
	if !isStrongPass(ur.Password) {
		v.SetError("Password", "password-weak")
	}
	if ur.Password != ur.PasswordConfirm {
		v.SetError("PasswordConfirm", "passwords-mismatch")
	}
}

func (upu *UserPasswordUpdate) UpdatePassword(user *models.User) (*models.User, bool) {
	if ok := user.CheckPassword(upu.CurrentPassword); !ok {
		upu.addError("user-login-or-password-incorrect")
		return nil, false
	}

	if ok := user.SetPassword(upu.Password); !ok {
		upu.addError("user-password-set-failed")
		return nil, false
	}

	ok := user.Save()
	if !ok {
		upu.addError("user-save-failed")
	}
	return user, ok
}
