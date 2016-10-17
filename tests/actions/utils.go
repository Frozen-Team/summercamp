package actions

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/astaxie/beego"
)

// login loges the user in the system and returns the cookie after the request.
// the cookie is then can be added to the request so, the app know the request is being done
// with some user logged in the system
func login() *http.Cookie {
	body := bytes.NewReader([]byte(`{"email":"olehgol260@gmail.com", "password":"1235~"}`))
	r, _ := http.NewRequest("POST", "/v1/users/login", body)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	if len(r.Cookies()) == 0 || r.Cookies()[0] == nil {
		panic("invalid cookies after login")
	}
	return r.Cookies()[0]
}
