package controllers

import "bitbucket.org/SummerCampDev/summercamp/models/forms"

// Operations about Teams
type Teams struct {
	ApplicationController
}

// Register reads the data from the request body into forms.TeamRegistration struct and attempts to save a team to db
// @Title Register
// @Description Team registration
// @Param body body string true "Registration info"
// @Success 200 {object} models.Team
// @router / [post]
func (t *Teams) Register() {
	form := new(forms.TeamRegistration)

	if ok := t.unmarshalJSON(form); !ok {
		t.serveAJAXBadRequest()
		return
	}

	team, ok := form.Register(t.currentUser)
	if !ok {
		t.serveAJAXBadRequest(form.Errors...)
		return
	}
	t.serveAJAXSuccess(team)
}

// @Title Delete
// @Description Team removal
// @Param objectId path int true "the team id you want to get"
// @Success 200 {object} models.Team
// @router /:objectId [delete]
func (t *Teams) Delete() {
	form := new(forms.TeamRegistration)

	if ok := t.unmarshalJSON(form); !ok {
		t.serveAJAXBadRequest()
		return

	}

	team, ok := form.Register(t.currentUser)
	if !ok {
		t.serveAJAXBadRequest(form.Errors...)
		return
	}
	t.serveAJAXSuccess(team)
}
