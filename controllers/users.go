package controllers

import "bitbucket.org/SummerCampDev/summercamp/models/forms"

type UsersController struct {
	ApplicationController
}

// Register reads the data from the request body into forms.UserReg struct and attempts to save a user to db
// POST /users
func (uc *UsersController) Register() {
	regForm := new(forms.UserReg)

	if ok := uc.unmarshalJSON(regForm); !ok {
		uc.ServeAJAXError([]string{"bad-data"}, nil)
		return
	}

	user, ok := regForm.Register()
	if !ok {
		uc.ServeAJAXError(regForm.Errors, nil)
		return
	}

	uc.ServeAJAXSuccess(user)
}
