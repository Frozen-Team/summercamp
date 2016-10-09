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
// @Failure 401 Unauthorized
// @router / [post]
func (t *Teams) Register() {
	if t.currentUser.Type != models.SpecTypeExecutor {
		t.serveAJAXMethodNotAllowed()
		return
	}
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

// AddMember adds new member to the team
// @Title AddMember
// @Description Team member addition
// @Param teamId path int true "the team id"
// @Param body body string true "Team member"
// @Success 200 {object} models.TeamMember
// @Failure 400 invalid-team-id or no-such-team
// @Failure 401 Unauthorized
// @router /:teamId/members [post]
func (t *Teams) AddMember() {
	if id, err := strconv.Atoi(t.Ctx.Input.Param(":id")); err != nil {
		t.serveAJAXBadRequest("invalid-team-id")
		return
	} else {
		form := new(forms.TeamMemberAddition)
		if ok := t.unmarshalJSON(form); !ok {
			t.serveAJAXBadRequest()
			return
		}
		if team, ok := models.Teams.FetchByID(id); ok {
			if currentMember, found := team.IsMember(t.currentUser); found {
				if currentMember.Access == models.LevelCreator {
					if member, ok := form.AddMember(team); ok {
						t.serveAJAXSuccess(member)
					} else {
						t.serveAJAXBadRequest()
					}
				} else {
					t.serveAJAXForbidden()
				}
			} else {
				t.serveAJAXForbidden()
			}
		} else {
			t.serveAJAXBadRequest("no-such-team")
		}
	}
}

// @Title Delete
// @Description Team removal
// @Param teamId path int true "the team id you want to delete"
// @Success 200 OK
// @Failure 400 invalid-team-id or no-such-team
// @Failure 401 Unauthorized
// @router /:teamId [delete]
func (t *Teams) Delete() {
	if id, err := strconv.Atoi(t.Ctx.Input.Param(":id")); err != nil {
		t.serveAJAXBadRequest("invalid-team-id")
		return
	} else {
		if team, ok := models.Teams.FetchByID(id); ok {
			if teamMember, found := team.IsMember(t.currentUser); found && teamMember.Access == models.LevelCreator {
				if team.Delete() {
					t.serveAJAXSuccess(team)
				} else {
					t.serveAJAXInternalServerError()
				}
			} else {
				t.serveAJAXForbidden()
			}
		} else {
			t.serveAJAXBadRequest("no-such-team")
		}
	}
}

// @Title GetTeam
// @Description Get info about a team by its id
// @Param id path int true "An id of a team you want to get"
// @Success 200 {object} models.Team
// @Failure 400 invalid-team-id or no-such-team
// @Failure 401 Unauthorized
// @router /:id [get]
func (t *Teams) GetTeam() {
	// TODO: Check if the requested user can be seen (publicly or privately)
	if id, err := strconv.Atoi(t.Ctx.Input.Param(":id")); err != nil {
		t.serveAJAXBadRequest("invalid-team-id")
	} else {
		team, ok := models.Teams.FetchByID(id)
		if ok {
			t.serveAJAXSuccess(team)
		} else {
			t.serveAJAXBadRequest("no-such-team")
		}
	}
}
