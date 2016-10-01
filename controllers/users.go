package controllers

import (
	"bitbucket.org/SummerCampDev/summercamp/models"
	"bitbucket.org/SummerCampDev/summercamp/models/forms"
	"net/http"
	"strconv"
)

// Operations about Users
type Users struct {
	ApplicationController
}

func (u *Users) Prepare() {
	u.SkipAuthorizationActions("Register", "Login")
	u.ApplicationController.Prepare()
}

// Register reads the data from the request body into forms.UserReg struct and attempts to save a user to db
// @Title Register
// @Description User registration
// @Param body body string true "Registration info"
// @Success 200 {object} models.User
// @Failure 200 Nil object and error tag
// @router /users/ [post]
func (uc *Users) Register() {
	regForm := new(forms.UserRegistration)

	if ok := uc.unmarshalJSON(regForm); !ok {
		uc.serveAJAXBadRequest()
		return
	}

	user, ok := regForm.Register()
	if !ok {
		uc.serveAJAXBadRequest(regForm.Errors...)
		return
	}
	uc.authorizeUser(user)
	uc.serveAJAXSuccess(user)
}

// Login reads the data from the request body into forms.UserLogin struct, attempts to query a user from the db
// by email and checks password. In case of success the user is authorized
// @Title Login
// @Description Login a user to the system
// @Param body body string true "Json body message with user credentials"
// @Success 200 {object} models.User
// @Failure 200 nil object
// @router /users/login [post]
func (u *Users) Login() {
	loginForm := new(forms.UserLogin)

	if ok := u.unmarshalJSON(loginForm); !ok {
		u.serveAJAXBadRequest()
		return
	}

	user, ok := loginForm.Login()
	if !ok {
		u.serveAJAXBadRequest(loginForm.Errors...)
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
// @router /users/logout [post]
func (u *Users) Logout() {
	u.deauthorizeUser()
	u.serveAJAXSuccess(nil)
}

// @Title Current
// @Description Get info about the currently logged in user
// @Success 200 {object} models.User
// @Failure 200 bad-request
// @router /users/current [get]
func (u *Users) Current() {
	u.serveAJAXSuccess(u.currentUser)
}

// @Title UpdateField
// @Description Updates user field
// @Param body body string true "A body that should contain a field name and new value"
// @Success 200 {object} models.User
// @Failure 401 Unauthorized
// @Failure 400 bad-data
// @router /users/update_field [post]
func (u *Users) UpdateField() {
	form := &forms.UserUpdate{}

	if ok := u.unmarshalJSON(form); !ok {
		u.serveAJAXBadRequest()
		return
	}
	if _, ok := form.Update(u.currentUser); ok {
		u.serveAJAXSuccess(u.currentUser)
		return
	}
	u.serveAJAXError(nil, http.StatusInternalServerError, form.Errors...)
}

// @Title GetUser
// @Description Get info about a user by its id
// @Param id path int true "An id of a user you want to get"
// @Success 200 {object} models.User
// @Failure 400 invalid-id or no-such-user
// @router /users/:id [get]
func (u *Users) GetUser() {
	// TODO: Check if the requested user can be seen (publicly or privately)
	if id, err := strconv.Atoi(u.Ctx.Input.Param(":id")); err != nil {
		u.serveAJAXBadRequest("invalid-id")
	} else {
		user, ok := models.Users.FetchByID(id)
		if ok {
			u.serveAJAXSuccess(user)
		} else {
			u.serveAJAXBadRequest("no-such-user")
		}
	}
}
