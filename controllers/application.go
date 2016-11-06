package controllers

import (
	"encoding/json"

	"net/http"

	"strconv"

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

const (
	URLParamID        = ":id"
	URLParamVacancyID = ":v_id"
)

// ApplicationController is a base controller for all controllers in the project.
// It contains common helper methods.
type ApplicationController struct {
	beego.Controller
	currentUser              *models.User
	skipAuthorizationActions []string
}

func (a *ApplicationController) SkipAuthorizationActions(action ...string) {
	a.skipAuthorizationActions = append(a.skipAuthorizationActions, action...)
}

func (a *ApplicationController) Prepare() {
	_, action := a.GetControllerAndAction()
	for _, a := range a.skipAuthorizationActions {
		if a == action {
			return
		}
	}
	user := a.authorizedUser()
	if user == nil {
		a.serveAJAXUnauthorized()
		a.StopRun()
		return
	}
	a.currentUser = user
}

// serveAJAXSuccessMeta serve success AJAX as described in ServeSuccess plus some external meta data.
func (a *ApplicationController) serveAJAXSuccessMeta(data interface{}, meta map[string]interface{}) {
	a.serveAJAX(false, data, meta, "")
}

// serveAJAXSuccess response success ajax with the given data into json format
func (a *ApplicationController) serveAJAXSuccess(data interface{}) {
	a.setStatusCode(http.StatusOK)
	a.serveAJAXSuccessMeta(data, nil)
}

// serveAJAXErrorMeta serve error AJAX as described in ServeError plus some external meta data
func (a *ApplicationController) serveAJAXErrorMeta(data interface{}, meta map[string]interface{}, errors ...string) {
	a.serveAJAX(true, data, meta, errors...)
}

// serveAJAXError response AJAX error with true "has-error" and "errors" equals to the all occurred errors
// The specified data is passed directly to responseAJAX.
func (a *ApplicationController) serveAJAXError(data interface{}, code int, errors ...string) {
	a.setStatusCode(code)
	a.serveAJAXErrorMeta(data, nil, errors...)
}

// serveAJAXBadRequest is a wrapper, which also sets the status code to 400
func (a *ApplicationController) serveAJAXBadRequest(errors ...string) {
	errors = append(errors, "bad-request")
	a.serveAJAXError(nil, http.StatusBadRequest, errors...)
}

// setStatusCode sets status code for the current response
func (a *ApplicationController) setStatusCode(code int) {
	a.Ctx.ResponseWriter.WriteHeader(code)
}

// serveAJAXMethodNotAllowed is a wrapper to serve "method-not-allowed" error
func (a *ApplicationController) serveAJAXMethodNotAllowed() {
	a.serveAJAXError(nil, http.StatusMethodNotAllowed, "method-not-allowed")
}

// serveAJAXUnauthorized is a wrapper to serve "unauthorized" error
func (a *ApplicationController) serveAJAXUnauthorized() {
	a.serveAJAXError(nil, http.StatusUnauthorized, "unauthorized")
}

// serveAJAXInternalServerError is a wrapper to serve "internal-error" error
func (a *ApplicationController) serveAJAXInternalServerError() {
	a.serveAJAXError(nil, http.StatusInternalServerError, "internal-error")
}

// serveAJAXInternalServerError is a wrapper to serve "internal-error" error
func (a *ApplicationController) serveAJAXForbidden() {
	a.serveAJAXError(nil, http.StatusForbidden, "forbidden")
}

// serveAJAX response with the json with two keys: "meta" and "data".
// result meta consists of argument 'meta' map  merged with error and errorType into one map.
// if 'meta' argument contains key-value pairs which may override specified error and errorType, these values
// are skipped so the original arguments are primer.
func (a *ApplicationController) serveAJAX(hasError bool, data interface{}, meta map[string]interface{}, errors ...string) {
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

// getID retrieve an id from the url (not as a url parameter). If the id is of an invalid
// value, the "invalid-id" error is served and the caller action  terminates with a StopRun method.
func (a *ApplicationController) getID() int {
	id, err := a.getUrlIntValue(URLParamID)
	if err != nil {
		a.serveAJAXBadRequest("invalid-id")
		a.StopRun()
	}
	return id
}

// getUrlIntValue retrieve an int value by the given key(e.g. :id) from the url
func (a *ApplicationController) getUrlIntValue(key string) (int, error) {
	return strconv.Atoi(a.Ctx.Input.Param(key))
}
