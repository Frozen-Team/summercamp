package controllers

import (
	"github.com/astaxie/beego"
)

// keys which must be in meta map (not in external meta maps. External meta maps contains additional data, which is then
// merge with data under keys described below)
const (
	errorKey     = "error"
	errorTypeKey = "error-type"
)

// ApplicationController is a base controller for all controllers in the project.
// It contains common helper methods.
type ApplicationController struct {
	beego.Controller
}

// buildMeta make a result meta map by merging error, errorType and externalMeta data into one map.
// if externalMeta contains key-value pairs which may override specified error and errorType, these values
// are skipped so the original arguments are primer.
func buildMeta(error bool, errorType interface{}, externalMeta map[string]interface{}) map[string]interface{} {
	resMeta := map[string]interface{}{
		errorKey:     error,
		errorTypeKey: errorType,
	}
	if externalMeta != nil {
		for key, value := range externalMeta {
			if key != errorKey && key != errorTypeKey {
				resMeta[key] = value
			}
		}
	}
	return resMeta
}

// ServeSuccessMeta serve success AJAX as described in ServeSuccess plus some external meta data
func (a *ApplicationController) ServeSuccessMeta(data interface{}, meta map[string]interface{}) {
	a.serveAJAX(data, buildMeta(false, "", meta))
}

// ResponseSuccess response AJAX success with false "error" and empty "error-type"
// The specified data is passed directly to responseAJAX.
func (a *ApplicationController) ServeSuccess(data interface{}) {
	a.ServeSuccessMeta(data, nil)
}

// ServeErrorMeta serve error AJAX as described in ServeError plus some external meta data
func (a *ApplicationController) ServeErrorMeta(error interface{}, data interface{}, meta map[string]interface{}) {
	a.serveAJAX(data, buildMeta(true, error, meta))
}

// ResponseError response AJAX error with true "error" and "error-type" equals to specified error argument
// The specified data is passed directly to responseAJAX.
func (a *ApplicationController) ServeError(error interface{}, data interface{}) {
	a.ServeErrorMeta(error, data, nil)
}

// responseAJAX response with the json with two keys: "meta" and "data" value of which corresponds
// to the specified arguments.
func (a *ApplicationController) serveAJAX(data interface{}, meta map[string]interface{}) {
	response := struct {
		Meta map[string]interface{} `json:"meta"`
		Data interface{}            `json:"data"`
	}{
		Meta: meta,
		Data: data,
	}
	a.Data["json"] = &response
	a.ServeJSON()
}
