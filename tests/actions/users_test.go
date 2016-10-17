package actions

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "bitbucket.org/SummerCampDev/summercamp/routers"
	_ "bitbucket.org/SummerCampDev/summercamp/tests/setup"

	"strconv"

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

			body_ := bytes.NewReader([]byte(`{"field":"  email   ", "value":" olehgol260@gmail.com  "}`))
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

func TestUpdatePassword(t *testing.T) {
	cookie := login()

	Convey("Test update password", t, func() {
		w := httptest.NewRecorder()

		Convey("Test valid password change", func() {
			body := bytes.NewReader([]byte(`{"current_password":"1235~", "password":"1234~", "password_confirm":"1234~"}`))
			r, _ := http.NewRequest("POST", "/v1/users/update_password", body)
			r.AddCookie(cookie)
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, http.StatusOK)
			response, err := ReadResponse(w.Body)
			So(err, ShouldBeNil)

			So(response.Meta.HasError, ShouldBeFalse)
			So(response.Data, ShouldNotBeNil)

			body = bytes.NewReader([]byte(`{"current_password":"1234~", "password":"1235~", "password_confirm":"1235~"}`))
			r, _ = http.NewRequest("POST", "/v1/users/update_password", body)
			r.AddCookie(cookie)
			beego.BeeApp.Handlers.ServeHTTP(w, r)
		})

		Convey("Test invalid password change: different pass and confirm", func() {
			body := bytes.NewReader([]byte(`{"current_password":"1235~", "password":"1234~", "password_confirm":"1231~"}`))
			r, _ := http.NewRequest("POST", "/v1/users/update_password", body)
			r.AddCookie(cookie)
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, http.StatusBadRequest)
			response, err := ReadResponse(w.Body)
			So(err, ShouldBeNil)

			So(response.Meta.HasError, ShouldBeTrue)
			So(response.Data, ShouldBeNil)
		})

		Convey("Test invalid password change: weak pass", func() {
			body := bytes.NewReader([]byte(`{"current_password":"1235~", "password":"12345", "password_confirm":"12345"}`))
			r, _ := http.NewRequest("POST", "/v1/users/update_password", body)
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

func TestGetUser(t *testing.T) {
	cookie := login()

	Convey("Test get valid user", t, func() {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/v1/users/1", nil)
		r.AddCookie(cookie)
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		response, err := ReadResponse(w.Body)
		So(err, ShouldBeNil)

		So(response.Meta.HasError, ShouldBeFalse)
		So(response.Data, ShouldNotBeNil)

	})

	Convey("Test get invalid user", t, func() {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/v1/users/100500", nil)
		r.AddCookie(cookie)
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
		response, err := ReadResponse(w.Body)
		So(err, ShouldBeNil)

		So(response.Meta.HasError, ShouldBeTrue)
		So(response.Data, ShouldBeNil)

	})
}

func TestAddAndRemoveSkill(t *testing.T) {
	cookie := login()

	Convey("Test valid add and remove", t, func() {
		w := httptest.NewRecorder()
		body := bytes.NewReader([]byte(`{"skill_id":1}`))
		r, _ := http.NewRequest("POST", "/v1/users/skills", body)
		r.AddCookie(cookie)

		beego.BeeApp.Handlers.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		response, err := ReadResponse(w.Body)

		So(err, ShouldBeNil)
		So(response.Meta.HasError, ShouldBeFalse)
		So(response.Data, ShouldNotBeNil)

		responseMap, ok := response.Data.(map[string]interface{})
		So(ok, ShouldBeTrue)
		id := int(responseMap["id"].(float64))
		r, _ = http.NewRequest("DELETE", "/v1/users/skills/"+strconv.FormatInt(int64(id), 10), body)
		w_ := httptest.NewRecorder()
		r.AddCookie(cookie)

		beego.BeeApp.Handlers.ServeHTTP(w_, r)
		So(w_.Code, ShouldEqual, http.StatusOK)
		response_, err := ReadResponse(w_.Body)
		So(err, ShouldBeNil)

		So(response_.Meta.HasError, ShouldBeFalse)
		So(response_.Data, ShouldBeNil)
	})
}

func TestAddAndRemoveSphere(t *testing.T) {
	cookie := login()

	Convey("Test valid add and remove", t, func() {
		w := httptest.NewRecorder()
		body := bytes.NewReader([]byte(`{"sphere_id":1}`))
		r, _ := http.NewRequest("POST", "/v1/users/spheres", body)
		r.AddCookie(cookie)

		beego.BeeApp.Handlers.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		response, err := ReadResponse(w.Body)

		So(err, ShouldBeNil)
		So(response.Meta.HasError, ShouldBeFalse)
		So(response.Data, ShouldNotBeNil)

		responseMap, ok := response.Data.(map[string]interface{})
		So(ok, ShouldBeTrue)
		id := int(responseMap["id"].(float64))
		r, _ = http.NewRequest("DELETE", "/v1/users/spheres/"+strconv.FormatInt(int64(id), 10), body)
		w_ := httptest.NewRecorder()
		r.AddCookie(cookie)

		beego.BeeApp.Handlers.ServeHTTP(w_, r)
		So(w_.Code, ShouldEqual, http.StatusOK)
		response_, err := ReadResponse(w_.Body)
		So(err, ShouldBeNil)

		So(response_.Meta.HasError, ShouldBeFalse)
		So(response_.Data, ShouldBeNil)
	})
}
