package controllers

import "bitbucket.org/SummerCampDev/summercamp/models/forms"

// Operations about Projects
type Projects struct {
	ApplicationController
}

// @Title Save
// @Description saves a new project in the db
// @Param body body string true "fields:description, budget, array of sphere_skills"
// @Success 200 {object} models.Project
// @Failure 400 possible errors, nil object
// @Failure 401 unauthorized
// @router / [post]
func (p *Projects) Save() {
	form := new(forms.Project)

	if !p.currentUser.IsClient() {
		p.serveAJAXForbidden()
		return
	}

	if ok := p.unmarshalJSON(form); !ok {
		p.serveAJAXInternalServerError()
		return
	}

	form.ClientID = p.currentUser.ID

	if ok := forms.Validate(form); !ok {
		p.serveAJAXBadRequest(form.Errors...)
		return
	}

	if project, ok := form.Save(); ok {
		p.serveAJAXSuccess(project)
	} else {
		p.serveAJAXInternalServerError()
	}

}
