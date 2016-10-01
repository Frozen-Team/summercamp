package controllers

import (
	"encoding/json"

	"net/http"

	"bitbucket.org/SummerCampDev/summercamp/models"
	"github.com/astaxie/beego"
)

// keys which must be in meta map (not in external meta maps. External meta maps contains additional data, which is then
// merge with data under keys described below)
const (
	AJAXKeyHasError = "has-error"
	AJAXKeyErrors   = "errors"
)
const (
	SessionKeyUser = "user"
)

// ApplicationController is a base controller for all controllers in the project.
// It contains common helper methods.
type ApplicationController struct {
	beego.Controller
}

// serveAJAXSuccessMeta serve success AJAX as described in ServeSuccess plus some external meta data.
func (a *ApplicationController) serveAJAXSuccessMeta(data interface{}, meta map[string]interface{}) {
	a.serveAJAX(false, data, meta, "")
}

// serveAJAXSuccess response success ajax with the given data into json format
func (a *ApplicationController) serveAJAXSuccess(data interface{}) {
	a.serveAJAXSuccessMeta(data, nil)
}

// serveAJAXErrorMeta serve error AJAX as described in ServeError plus some external meta data
func (a *ApplicationController) serveAJAXErrorMeta(data interface{}, meta map[string]interface{}, errors ...interface{}) {
	a.serveAJAX(true, data, meta, errors...)
}

// serveAJAXError response AJAX error with true "has-error" and "errors" equals to the all occurred errors
// The specified data is passed directly to responseAJAX.
func (a *ApplicationController) serveAJAXError(data interface{}, errors ...interface{}) {
	a.serveAJAXErrorMeta(data, nil, errors...)
}

// serveAJAXUnauthorized is a convenient wrapper above serveAJAXError to serve "unauthorized" error
func (a *ApplicationController) serveAJAXUnauthorized() {
	a.serveAJAXError(nil, "unauthorized")
}

// serveAJAX response with the json with two keys: "meta" and "data".
// result meta consists of argument 'meta' map  merged with error and errorType into one map.
// if 'meta' argument contains key-value pairs which may override specified error and errorType, these values
// are skipped so the original arguments are primer.
func (a *ApplicationController) serveAJAX(hasError bool, data interface{}, meta map[string]interface{}, errors ...interface{}) {
	response := struct {
		Meta map[string]interface{} `json:"meta"`
		Data interface{}            `json:"data"`
	}{
		Meta: map[string]interface{}{
			AJAXKeyHasError: hasError,
			AJAXKeyErrors:   errors,
		},
		Data: data,
	}

	if meta != nil {
		for key, value := range meta {
			if key != AJAXKeyHasError && key != AJAXKeyErrors {
				response.Meta[key] = value
			}
		}
	}

	a.Data["json"] = &response
	a.ServeJSON()
}

// isAuthorized returns true if the user is authorized, false otherwise
func (a *ApplicationController) isAuthorized() bool {
	return a.authorizedUser() != nil
}

// authorizedUser returns authorized user. Returns nil, if the user is not authorized.
func (a *ApplicationController) authorizedUser() *models.User {
	u := a.GetSession(SessionKeyUser)
	if u == nil {
		return nil
	}
	if id, ok := u.(int); ok {
		user, _ := models.Users.FetchByID(id)
		return user
	}
	return nil
}

// authorizeUser set user id to session.
func (a *ApplicationController) authorizeUser(user *models.User) {
	a.SetSession(SessionKeyUser, user.ID)
}

// deauthorizeUser removes user id from the session.
func (a *ApplicationController) deauthorizeUser() {
	a.DelSession(SessionKeyUser)
}

// redirectToSpecialityIndex redirects to index path according to passed Speciality.
func (r *ApplicationController) redirectToSpecialityIndex(s models.Speciality) {
	switch s {
	case models.SpecTypeExecutor:
		r.Redirect(beego.URLFor("Executor.Index"), http.StatusMovedPermanently)
	case models.SpecTypeClient:
		r.Redirect(beego.URLFor("Client.Index"), http.StatusMovedPermanently)
	case models.SpecTypeManager:
		r.Redirect(beego.URLFor("Manager.Index"), http.StatusMovedPermanently)
	}
	beego.BeeLogger.Error("Trying to redirect to bad speciality index path '%d'", s)
}

// unmarshalJSON parses a json object from request body and fills the fields of the v interface.
// If json.Unmarshal fails, a corresponding msg will be added to the log
func (a *ApplicationController) unmarshalJSON(v interface{}) bool {
	err := json.Unmarshal(a.Ctx.Input.RequestBody, v)
	if err != nil {
		beego.BeeLogger.Error("Error while unmarshaling: %v", err)
		return false
	}
	return true
}
