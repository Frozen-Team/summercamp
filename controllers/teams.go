package controllers

import (
	"bitbucket.org/SummerCampDev/summercamp/models"
	"bitbucket.org/SummerCampDev/summercamp/models/forms"
)

// Operations about Teams
type Teams struct {
	ApplicationController
}

// checkIfAllowed is a private method to check if the team specified as a path parameter exists and
// if the current user is the creator of this team.
// The method returns a *models.Team in case if the caller requires some info about the team and bool value
// where true - everything is okay, false - there are troubles and the caller must terminate the action.
func (t *Teams) checkIfAllowed() (*models.Team, bool) {
	id := t.getID()

	team, ok := models.Teams.FetchByID(id)
	if !ok {
		t.serveAJAXBadRequest("no-such-team")
		return nil, false
	}

	currentMember, found := team.GetMember(t.currentUser)
	if !found || !currentMember.IsCreator() {
		t.serveAJAXForbidden()
		return nil, false
	}
	return team, true
}

// @Title Register
// @Description register a new team
// @Param body body string true "Registration info"
// @Success 200 {object} models.Team
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 500 internal-error
// @router / [post]
func (t *Teams) Save() {
	if t.currentUser.Type != models.SpecTypeExecutor {
		t.serveAJAXForbidden()
		return
	}

	form := new(forms.TeamRegistration)

	if ok := t.unmarshalJSON(form); !ok {
		t.serveAJAXInternalServerError()
		return
	}

	if ok := forms.Validate(form); !ok {
		t.serveAJAXBadRequest(form.Errors...)
		return
	}

	if team, ok := form.Register(t.currentUser); ok {
		t.serveAJAXSuccess(team)
	} else {
		t.serveAJAXInternalServerError()
	}
}

// @Title AddMember
// @Description adds a new member to the team
// @Param teamId path int true "the team id"
// @Param body body string true "Team member"
// @Success 200 {object} models.TeamMember
// @Failure 400 invalid-team-id or no-such-team
// @Failure 401 Unauthorized
// @Failure 403 forbidden
// @Failure 500 internal-error
// @router /:id/members [post]
func (t *Teams) AddMember() {
	team, ok := t.checkIfAllowed()
	if !ok {
		return
	}
	//
	//id := t.getID()
	//
	//team, ok := models.Teams.FetchByID(id)
	//if !ok {
	//	t.serveAJAXBadRequest("no-such-team")
	//	return
	//}
	//
	//currentMember, found := team.GetMember(t.currentUser)
	//if !found || !currentMember.IsCreator() {
	//	t.serveAJAXForbidden()
	//	return
	//}

	form := new(forms.TeamMemberAddition)
	if ok := t.unmarshalJSON(form); !ok {
		t.serveAJAXInternalServerError()
		return
	}

	if ok := forms.Validate(form); !ok {
		t.serveAJAXBadRequest(form.Errors...)
		return
	}

	if member, ok := form.AddMember(team); ok {
		t.serveAJAXSuccess(member)
	} else {
		t.serveAJAXInternalServerError()
	}
}

// TODO: if there are members in the team, what then?
// @Title Delete
// @Description Team removal
// @Param teamId path int true "the team id you want to delete"
// @Success 200 OK
// @Failure 400 invalid-team-id or no-such-team
// @Failure 401 Unauthorized
// @Failure 403 forbidden
// @Failure 500 internal-error
// @router /:id [delete]
func (t *Teams) Delete() {
	team, ok := t.checkIfAllowed()
	if !ok {
		return
	}
	//id := t.getID()
	//
	//team, ok := models.Teams.FetchByID(id)
	//if !ok {
	//	t.serveAJAXBadRequest("no-such-team")
	//	return
	//}
	//
	//teamMember, found := team.GetMember(t.currentUser)
	//if !found || !teamMember.IsCreator() {
	//	t.serveAJAXForbidden()
	//	return
	//}

	if team.Delete() {
		t.serveAJAXSuccess(team)
	} else {
		t.serveAJAXInternalServerError()
	}
}

// @Title GetTeam
// @Description Get info about a team by its id
// @Param id path int true "An id of a team you want to get"
// @Success 200 {object} models.Team
// @Failure 400 no-such-team
// @Failure 401 Unauthorized
// @router /:id [get]
func (t *Teams) GetTeam() {
	// TODO: Check if the requested user can be seen (publicly or privately)
	id := t.getID()

	if team, ok := models.Teams.FetchByID(id); ok {
		t.serveAJAXSuccess(team)
	} else {
		t.serveAJAXBadRequest("no-such-team")
	}
}

// @Title AddVacancy
// @Description Add vacancy for a given team
// @Param id path int true "An id of a team you want to get"
// @Success 200 {object} models.Vacancy
// @Failure 400 no-such-team
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @router /:id/vacancies [post]
func (t *Teams) AddVacancy() {
	team, ok := t.checkIfAllowed()
	if !ok {
		return
	}
	//teamID := t.getID()
	//
	//team, ok := models.Teams.FetchByID(teamID)
	//if !ok {
	//	t.serveAJAXBadRequest("no-such-team")
	//	return
	//}
	//
	//teamMember, found := team.GetMember(t.currentUser)
	//if !found || !teamMember.IsCreator() {
	//	t.serveAJAXForbidden()
	//	return
	//}

	form := new(forms.Vacancy)
	if ok := t.unmarshalJSON(form); !ok {
		t.serveAJAXInternalServerError()
		return
	}

	form.TeamID = team.ID
	if ok := forms.Validate(form); !ok {
		t.serveAJAXBadRequest(form.Errors...)
		return
	}

	if vacancy, ok := form.Save(); ok {
		t.serveAJAXSuccess(vacancy)
	} else {
		t.serveAJAXInternalServerError()
	}
}

// @Title RemoveVacancy
// @Description remove vacancy for a given team
// @Param id path int true "An id of a team you want to get"
// @Success 200 {object} models.Vacancy
// @Failure 400 no-such-team
// @Failure 400 no-v_id
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @router /:id/vacancies/:v_id [delete]
func (t *Teams) RemoveVacancy() {
	_, ok := t.checkIfAllowed()
	if !ok {
		return
	}

	vacancyID, err := t.getUrlIntValue(URLParamVacancyID)
	if err != nil {
		t.serveAJAXBadRequest("no-" + URLParamVacancyID)
		return
	}

	if ok := models.Vacancies.DeleteByID(vacancyID); ok {
		t.serveAJAXSuccess(nil)
	} else {
		t.serveAJAXInternalServerError()
	}
}
