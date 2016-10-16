package actions

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/astaxie/beego"
)

func init() {
	body := bytes.NewReader([]byte(`{"email":"my_mail@mail.com", "type":"manager", "first_name":"oleh",
		 "last_name":"gol", "password":"1235~", "password_confirm":"1235~",
		  "country":"ukraine", "city":"kyiv"}`))
	r, _ := http.NewRequest("POST", "/v1/users", body)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	response, err := ReadResponse(w.Body)
	if err != nil || response.Meta.HasError {
		panic("Failed to register user on tests start")
	}

}

// login loges the user in the system and returns the cookie after the request.
// the cookie is then can be added to the request so, the app know the request is being done
// with some user logged in the system
func login() *http.Cookie {
	body := bytes.NewReader([]byte(`{"email":"my_mail@mail.com", "password":"1235~"}`))
	r, _ := http.NewRequest("POST", "/v1/users/login", body)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	if len(r.Cookies()) == 0 || r.Cookies()[0] == nil {
		panic("invalid cookies after login")
	}
	return r.Cookies()[0]
}
