package controllers

import "bitbucket.org/SummerCampDev/summercamp/models/forms"

type UsersController struct {
	ApplicationController
}

// Register reads the data from the request body into forms.UserReg struct and attempts to save a user to db
// @route POST /users
func (uc *UsersController) Register() {
	regForm := new(forms.UserRegistration)

	if ok := uc.unmarshalJSON(regForm); !ok {
		uc.ServeAJAXError(nil, "bad-data")
		return
	}

	user, ok := regForm.Register()
	if !ok {
		uc.ServeAJAXError(nil, regForm.Errors)
		return
	}

	uc.ServeAJAXSuccess(user)
}
