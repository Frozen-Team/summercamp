package controllers

import "bitbucket.org/SummerCampDev/summercamp/models/forms"

type UserSpheres struct {
	ApplicationController
}

// @Title Save
// @Description Create a new user sphere record
// @Param body body string true "if action==add, pass sphere_id, if action==remove, pass only an action"
// @Success 200 {object} if action==add, returns models.UserSphere
// @Success 200 {object} if action==remove, returns nothing
// @Failure 401 Unauthorized
// @Failure 400 errors about what goes wrong during unmarshaling or validation
// @router /projects [post]
func (us *UserSpheres) Save() {
	form := new(forms.UserSphere)

	if ok := us.unmarshalJSON(form); !ok {
		us.serveAJAXInternalServerError()
		return
	}

	if ok := forms.Validate(form); !ok {
		us.serveAJAXBadRequest(form.Errors...)
		return
	}

	form.UserID = us.currentUser.ID

	userSphere, ok := form.Save()
	if !ok {
		us.serveAJAXInternalServerError()
		return
	}

	us.serveAJAXSuccess(userSphere)
}
