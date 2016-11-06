package actions

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

// login loges the user in the system and returns the cookie after the request.
// the cookie is then can be added to the request so, the app know the request is being done
// with some user logged in the system
func loginExec() *http.Cookie {
	return login(`{"email":"olehgol260@gmail.com", "password":"1235~"}`)
}

func loginClient() *http.Cookie {
	return login(`{"email":"olehgol260_@gmail.com", "password":"1235~"}`)
}

func login(credentials string) *http.Cookie {
	body := bytes.NewReader([]byte(credentials))
	r, _ := http.NewRequest("POST", "/v1/users/login", body)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	if len(r.Cookies()) == 0 || r.Cookies()[0] == nil {
		panic("invalid cookies after login")
	}
	return r.Cookies()[0]
}

func checkOK(w *httptest.ResponseRecorder) {
	So(w.Code, ShouldEqual, http.StatusOK)
	response, err := ReadResponse(w.Body)
	So(err, ShouldBeNil)
	So(response.Meta.HasError, ShouldBeFalse)
	So(response.Data, ShouldNotBeNil)
}

func checkOKNoData(w *httptest.ResponseRecorder) {
	So(w.Code, ShouldEqual, http.StatusOK)
	response, err := ReadResponse(w.Body)
	So(err, ShouldBeNil)
	So(response.Meta.HasError, ShouldBeFalse)
	So(response.Data, ShouldBeNil)
}

func checkBad(w *httptest.ResponseRecorder, code int) {
	So(w.Code, ShouldEqual, code)
	response, err := ReadResponse(w.Body)
	So(err, ShouldBeNil)
	So(response.Meta.HasError, ShouldBeTrue)
}
