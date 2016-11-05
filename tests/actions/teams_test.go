package actions

import (
	"testing"

	_ "bitbucket.org/SummerCampDev/summercamp/routers"
	_ "bitbucket.org/SummerCampDev/summercamp/tests/setup"

	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTeamSave(t *testing.T) {
	Convey("valid save", t, func() {
		body := bytes.NewReader([]byte(`{"name":"bithces", "description":"a good team"}`))
		r, _ := http.NewRequest("POST", "/v1/teams", body)
		r.AddCookie(loginExec())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		response, err := ReadResponse(w.Body)
		So(err, ShouldBeNil)
		So(response.Meta.HasError, ShouldBeFalse)
	})

	Convey("invalid save: missing field", t, func() {
		body := bytes.NewReader([]byte(`{"description":"a good team"}`))
		r, _ := http.NewRequest("POST", "/v1/teams", body)
		r.AddCookie(loginExec())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
		response, err := ReadResponse(w.Body)
		So(err, ShouldBeNil)
		So(response.Meta.HasError, ShouldBeTrue)
	})

	Convey("invalid save: unauthorized", t, func() {
		body := bytes.NewReader([]byte(`{"name":"bithces", "description":"a good team"}`))
		r, _ := http.NewRequest("POST", "/v1/teams", body)
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusUnauthorized)
		response, err := ReadResponse(w.Body)
		So(err, ShouldBeNil)
		So(response.Meta.HasError, ShouldBeTrue)
	})
}

func TestAddMember(t *testing.T) {
	Convey("valid addition", t, func() {
		body := bytes.NewReader([]byte(`{"user_id":1, "access":100}`))
		r, _ := http.NewRequest("POST", "/v1/teams/1/members", body)
		r.AddCookie(loginExec())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		response, err := ReadResponse(w.Body)
		So(err, ShouldBeNil)
		So(response.Meta.HasError, ShouldBeFalse)
	})

	Convey("invalid addition:invalid team", t, func() {
		body := bytes.NewReader([]byte(`{"user_id":1, "access":100}`))
		r, _ := http.NewRequest("POST", "/v1/teams/100/members", body)
		r.AddCookie(loginExec())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
		response, err := ReadResponse(w.Body)
		So(err, ShouldBeNil)
		So(response.Meta.HasError, ShouldBeTrue)
	})
}

func TestDeleteTeam(t *testing.T) {
	Convey("valid deletion", t, func() {
		r, _ := http.NewRequest("DELETE", "/v1/teams/3", nil)
		r.AddCookie(loginExec())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		response, err := ReadResponse(w.Body)
		So(err, ShouldBeNil)
		So(response.Meta.HasError, ShouldBeFalse)
	})
}

func TestGetTeam(t *testing.T) {
	Convey("valid get", t, func() {
		r, _ := http.NewRequest("GET", "/v1/teams/1", nil)
		r.AddCookie(loginExec())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		response, err := ReadResponse(w.Body)
		So(err, ShouldBeNil)
		So(response.Meta.HasError, ShouldBeFalse)
		So(response.Data, ShouldNotBeNil)
	})

	Convey("invalid get", t, func() {
		r, _ := http.NewRequest("GET", "/v1/teams/100", nil)
		r.AddCookie(loginExec())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
		response, err := ReadResponse(w.Body)
		So(err, ShouldBeNil)
		So(response.Meta.HasError, ShouldBeTrue)
	})
}

func TestAddVacancy(t *testing.T) {
	Convey("Valid addition", t, func() {
		body := bytes.NewReader([]byte(`{"name":"front-end", "description":"we need a front-end dev", "skills":[1,2], "spheres":[1] }`))
		r, _ := http.NewRequest("POST", "/v1/teams/1/vacancies", body)
		r.AddCookie(loginExec())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		checkOK(w)
	})

	Convey("Invalid addition: unauthorized", t, func() {
		body := bytes.NewReader([]byte(`{"name":"front-end", "description":"we need a front-end dev", "skills":[1,2], "spheres":[1] }`))
		r, _ := http.NewRequest("POST", "/v1/teams/1/vacancies", body)
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		checkBad(w, http.StatusUnauthorized)
	})

	Convey("Invalid addition: unauthorized", t, func() {
		body := bytes.NewReader([]byte(`{"name":"front-end", "description":"we need a front-end dev", "skills":[1,2], "spheres":[1] }`))
		r, _ := http.NewRequest("POST", "/v1/teams/1/vacancies", body)
		r.AddCookie(loginClient())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		checkBad(w, http.StatusForbidden)
	})
}
