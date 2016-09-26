package controllers

import "bitbucket.org/SummerCampDev/summercamp/models/forms"

// Operations about Users
type Users struct {
	ApplicationController
}

// Register reads the data from the request body into forms.UserReg struct and attempts to save a user to db
// @Title Register
// @Description User registration
// @Param body body string true "Registration info"
// @Success 200 {object} models.User
// @router / [post]
func (uc *Users) Register() {
	regForm := new(forms.UserRegistration)

	if ok := uc.unmarshalJSON(regForm); !ok {
		uc.serveAJAXError(nil, "bad-data")
		return
	}

	user, ok := regForm.Register()
	if !ok {
		uc.serveAJAXError(nil, regForm.Errors)
		return
	}
	uc.authorizeUser(user)
	uc.serveAJAXSuccess(user)
}

// Login reads the data from the request body into forms.UserLogin struct, attempts to query a user from the db
// by email and checks password. In case of success the user is authorized
// @Title Login
// @Description Login a user to the system
// @Param body body string true "Body message"
// @Success 200 {object} models.User
// @Failure 200 nil object
// @router /login [post]
func (u *Users) Login() {
	if u.isAuthorized() {
		u.serveAJAXUnauthorized()
	}
	loginForm := new(forms.UserLogin)

	if ok := u.unmarshalJSON(loginForm); !ok {
		u.serveAJAXError(nil, "bad-data")
		return
	}

	user, ok := loginForm.Login()
	if !ok {
		u.serveAJAXError(nil, loginForm.Errors)
		return
	}
	u.authorizeUser(user)
	u.serveAJAXSuccess(user)
}

// Logout deauthorizes logged in User otherwise responses "bad-request"
// @Title Logout
// @Description Logout a user from the system
// @Success 200 {object} models.User
// @Failure 200 bad-request
// @router /logout [post]
func (u *Users) Logout() {
	user := u.authorizedUser()
	if user == nil {
		u.serveAJAXUnauthorized()
		return
	}
	u.deauthorizeUser(user)
	u.serveAJAXSuccess(nil)
}

// @Title Current
// @Description Get info about the currently logged in user
// @Success 200 {object} models.User
// @Failure 200 bad-request
// @router /current [get]
func (u *Users) Current() {
	user := u.authorizedUser()
	if user == nil {
		u.serveAJAXUnauthorized()
		return
	}

	u.serveAJAXSuccess(user)
}
