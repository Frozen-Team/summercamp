package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

// keys which must be in meta map (not in external meta maps. External meta maps contains additional data, which is then
// merge with data under keys described below)
const (
	AJAXKeyHasError = "has-error"
	AJAXKeyErrors   = "errors"
)

// ApplicationController is a base controller for all controllers in the project.
// It contains common helper methods.
type ApplicationController struct {
	beego.Controller
}

// ServeAJAXSuccessMeta serve success AJAX as described in ServeSuccess plus some external meta data
func (a *ApplicationController) ServeAJAXSuccessMeta(data interface{}, meta map[string]interface{}) {
	a.serveAJAX(false, data, meta, "")
}

// ServeAJAXSuccess response success ajax with the given data into json format
func (a *ApplicationController) ServeAJAXSuccess(data interface{}) {
	a.ServeAJAXSuccessMeta(data, nil)
}

// ServeAJAXErrorMeta serve error AJAX as described in ServeError plus some external meta data
func (a *ApplicationController) ServeAJAXErrorMeta(data interface{}, meta map[string]interface{}, errors ...interface{}) {
	a.serveAJAX(true, data, meta, errors)
}

// ServeAJAXError response AJAX error with true "has-error" and "errors" equals to the all occured errors
// The specified data is passed directly to responseAJAX.
func (a *ApplicationController) ServeAJAXError(data interface{}, errors ...interface{}) {
	a.ServeAJAXErrorMeta(data, nil, errors)
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
