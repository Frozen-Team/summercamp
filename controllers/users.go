package controllers

import "bitbucket.org/SummerCampDev/summercamp/models/forms"

type Users struct {
	ApplicationController
}

// Register reads the data from the request body into forms.UserReg struct and attempts to save a user to db
// @route POST /users
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

	uc.serveAJAXSuccess(user)
}

// Current serves the info about the current user. If the user is not authorized,
// a error is served.
func (uc *Users) Current() {
	user := uc.authorisedUser()
	if user != nil {
		uc.serveAJAXError(nil, "unauthorized")
		return
	}

	uc.serveAJAXSuccess(user)
}
