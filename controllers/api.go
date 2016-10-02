package controllers

// Operations about internal API
type Api struct {
	ApplicationController
}

func (a *Api) Prepare() {
	a.SkipAuthorizationActions("Ping")
	a.ApplicationController.Prepare()
}

// @Title Ping
// @Description Platform ping-pong
// @Success 200 OK
// @router /ping [get]
func (a *Api) Ping() {
	a.serveAJAXSuccess("OK")
}
