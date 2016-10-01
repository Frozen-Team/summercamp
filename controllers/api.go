package controllers

// Operations about internal API
type Api struct {
	ApplicationController
}

func (a *Api) Prepare() {
	a.SkipAuthorizationActions("Ping")
	a.Prepare()
}

// @Title Ping
// @Description Service ping-pong
// @Success 200 message Pong
// @router /ping/ [get]
func (a *Api) Ping() {
	a.serveAJAXSuccess("OK`")
}
