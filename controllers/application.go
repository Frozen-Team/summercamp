package controllers

import (
	"github.com/astaxie/beego"
)

// ApplicationController is a base controller for all controllers in the project.
// It contains common helper methods.
type ApplicationController struct {
	beego.Controller
}

// ResponseSuccess response AJAX success with false "error" and empty "error-type"
// The specified data is passed directly to responseAJAX.
func (a *ApplicationController) ResponseSuccess(data interface{}) {
	meta := map[string]interface{}{
		"error" : false,
		"error-type" : "",
	}
	a.responseAJAX(meta, data)
}

// ResponseError response AJAX error with true "error" and "error-type" equals to specified error argument
// The specified data is passed directly to responseAJAX.
func (a *ApplicationController) ResponseError(error interface{}, data interface{}) {
	meta := map[string]interface{}{
		"error" : true,
		"error-type" : error,
	}
	a.responseAJAX(meta, data)
}

// responseAJAX response with the json with two keys: "meta" and "data" value of which corresponds
// to the specified arguments.
func (a *ApplicationController) responseAJAX(meta map[string]interface{}, data interface{}) {
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