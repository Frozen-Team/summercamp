package controllers

import "bitbucket.org/SummerCampDev/summercamp/models/forms"

type Projects struct {
	ApplicationController
}

// @Title Save
// @Description saves a new project in the db
// @Param body body string true "fields:client_id, description, budget, array of sphere_skills"
// @Success 200 {object} models.Project
// @Failure 400 possible errors, nil object
// @Failure 401 unauthorized
// @router /projects [post]
func (p *Projects) Save() {
	newProjectForm := new(forms.Project)

	if ok := p.unmarshalJSON(newProjectForm); !ok {
		p.serveAJAXBadRequest()
		return
	}

	project, ok := newProjectForm.Save()

	if !ok {
		p.serveAJAXBadRequest(newProjectForm.Errors...)
		return
	}

	p.serveAJAXSuccess(project)
}
