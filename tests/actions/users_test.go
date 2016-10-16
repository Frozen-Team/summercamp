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

func TestUpdateField(t *testing.T) {
	cookie := login()

	Convey("Test update field", t, func() {
		w := httptest.NewRecorder()
		Convey("first_name update", func() {
			body := bytes.NewReader([]byte(`{"field":"first_name", "value":"petro"}`))
			r, _ := http.NewRequest("PUT", "/v1/users", body)
			r.AddCookie(cookie)
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, http.StatusOK)
			response, err := ReadResponse(w.Body)
			So(err, ShouldBeNil)

			So(response.Meta.HasError, ShouldBeFalse)
			So(response.Data, ShouldNotBeNil)
		})

		Convey("last_name update", func() {
			body := bytes.NewReader([]byte(`{"field":"last_name", "value":"petrenko"}`))
			r, _ := http.NewRequest("PUT", "/v1/users", body)
			r.AddCookie(cookie)
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, http.StatusOK)
			response, err := ReadResponse(w.Body)
			So(err, ShouldBeNil)

			So(response.Meta.HasError, ShouldBeFalse)
			So(response.Data, ShouldNotBeNil)
		})

		Convey("country update", func() {
			body := bytes.NewReader([]byte(`{"field":"country", "value":"USA"}`))
			r, _ := http.NewRequest("PUT", "/v1/users", body)
			r.AddCookie(cookie)
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, http.StatusOK)
			response, err := ReadResponse(w.Body)
			So(err, ShouldBeNil)

			So(response.Meta.HasError, ShouldBeFalse)
			So(response.Data, ShouldNotBeNil)
		})

		Convey("city update", func() {
			body := bytes.NewReader([]byte(`{"field":"city", "value":"New-York"}`))
			r, _ := http.NewRequest("PUT", "/v1/users", body)
			r.AddCookie(cookie)
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, http.StatusOK)
			response, err := ReadResponse(w.Body)
			So(err, ShouldBeNil)

			So(response.Meta.HasError, ShouldBeFalse)
			So(response.Data, ShouldNotBeNil)
		})

		Convey("overview update", func() {
			body := bytes.NewReader([]byte(`{"field":"overview", "value":"software engineer"}`))
			r, _ := http.NewRequest("PUT", "/v1/users", body)
			r.AddCookie(cookie)
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, http.StatusOK)
			response, err := ReadResponse(w.Body)
			So(err, ShouldBeNil)

			So(response.Meta.HasError, ShouldBeFalse)
			So(response.Data, ShouldNotBeNil)
		})

		Convey("summary update", func() {
			body := bytes.NewReader([]byte(`{"field":"  summary   ", "value":"   i'm the best of the best   "}`))
			r, _ := http.NewRequest("PUT", "/v1/users", body)
			r.AddCookie(cookie)
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, http.StatusOK)
			response, err := ReadResponse(w.Body)
			So(err, ShouldBeNil)

			So(response.Meta.HasError, ShouldBeFalse)
			So(response.Data, ShouldNotBeNil)
		})

		Convey("email update:valid email", func() {
			body := bytes.NewReader([]byte(`{"field":"  email   ", "value":" hello@mail.com  "}`))
			r, _ := http.NewRequest("PUT", "/v1/users", body)
			r.AddCookie(cookie)
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, http.StatusOK)
			response, err := ReadResponse(w.Body)
			So(err, ShouldBeNil)

			So(response.Meta.HasError, ShouldBeFalse)
			So(response.Data, ShouldNotBeNil)

			body_ := bytes.NewReader([]byte(`{"field":"  email   ", "value":" my_mail@mail.com  "}`))
			r_, _ := http.NewRequest("PUT", "/v1/users", body_)
			r_.AddCookie(cookie)
			beego.BeeApp.Handlers.ServeHTTP(w, r_)
		})

		Convey("email update:invalid email", func() {
			body := bytes.NewReader([]byte(`{"field":"  email   ", "value":" hello_mail.com  "}`))
			r, _ := http.NewRequest("PUT", "/v1/users", body)
			r.AddCookie(cookie)
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, http.StatusBadRequest)
			response, err := ReadResponse(w.Body)
			So(err, ShouldBeNil)

			So(response.Meta.HasError, ShouldBeTrue)
			So(response.Data, ShouldBeNil)

		})
	})
}
