package actions

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "bitbucket.org/SummerCampDev/summercamp/routers"
	_ "bitbucket.org/SummerCampDev/summercamp/tests/setup"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRegistrationAction(t *testing.T) {
	Convey("Test registration action", t, func() {
		body := bytes.NewReader([]byte(`{"email":"mail@mail.com", "type":"manager", "first_name":"oleh",
		 "last_name":"gol", "password":"1235~", "password_confirm":"1235~",
		  "country":"ukraine", "city":"kyiv"}`))
		r, _ := http.NewRequest("POST", "/v1/users", body)
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)
		beego.Trace("testing", "TestRegistrationAction", "Code[%d]\n%s", w.Code, w.Body.String())
		So(w.Code, ShouldEqual, http.StatusOK)
		response, err := ReadResponse(w.Body)
		So(err, ShouldBeNil)
		So(response.Meta.HasError, ShouldBeFalse)
	})
}

func TestCurrentAction(t *testing.T) {
	Convey("Test current action", t, func() {
		cookie := login()

		Convey("With login: Test current action", func() {
			r, _ := http.NewRequest("GET", "/v1/users/current", nil)
			r.AddCookie(cookie)
			w := httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(w, r)

			So(w.Code, ShouldEqual, http.StatusOK)
			So(w.Body.Len(), ShouldBeGreaterThan, 0)

			response, err := ReadResponse(w.Body)
			So(err, ShouldBeNil)

			So(response.Meta.HasError, ShouldBeFalse)
			So(response.Data, ShouldNotBeNil)
		})

		Convey("Without login: Test current action", func() {
			r, _ := http.NewRequest("GET", "/v1/users/current", nil)
			w := httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, http.StatusUnauthorized)
			response, err := ReadResponse(w.Body)
			So(err, ShouldBeNil)

			So(response.Meta.HasError, ShouldBeTrue)
			So(response.Data, ShouldBeNil)
		})
	})
}
