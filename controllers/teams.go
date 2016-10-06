package controllers

import (
	"bitbucket.org/SummerCampDev/summercamp/models/forms"
	"strconv"
	"bitbucket.org/SummerCampDev/summercamp/models"
)

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

// @Title GetTeam
// @Description Get info about a team by its id
// @Param id path int true "An id of a team you want to get"
// @Success 200 {object} models.User
// @Failure 400 invalid-id or no-such-user
// @router /:id [get]
func (u *Users) GetTeam() {
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
