package actions

import (
	"testing"

	_ "github.com/Frozen-Team/summercamp/routers"
	_ "github.com/Frozen-Team/summercamp/tests/setup"

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

		checkOK(w)
	})

	Convey("invalid save: missing field", t, func() {
		body := bytes.NewReader([]byte(`{"description":"a good team"}`))
		r, _ := http.NewRequest("POST", "/v1/teams", body)
		r.AddCookie(loginExec())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		checkBad(w, http.StatusBadRequest)
	})

	Convey("invalid save: unauthorized", t, func() {
		body := bytes.NewReader([]byte(`{"name":"bithces", "description":"a good team"}`))
		r, _ := http.NewRequest("POST", "/v1/teams", body)
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		checkBad(w, http.StatusUnauthorized)
	})
}

func TestAddMember(t *testing.T) {
	Convey("valid addition", t, func() {
		body := bytes.NewReader([]byte(`{"user_id":1, "access":100}`))
		r, _ := http.NewRequest("POST", "/v1/teams/1/members", body)
		r.AddCookie(loginExec())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		checkOK(w)
	})

	Convey("invalid addition:invalid team", t, func() {
		body := bytes.NewReader([]byte(`{"user_id":1, "access":100}`))
		r, _ := http.NewRequest("POST", "/v1/teams/100/members", body)
		r.AddCookie(loginExec())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		checkBad(w, http.StatusBadRequest)
	})
}

func TestDeleteTeam(t *testing.T) {
	Convey("valid deletion", t, func() {
		r, _ := http.NewRequest("DELETE", "/v1/teams/3", nil)
		r.AddCookie(loginExec())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		checkOKNoData(w)
	})
}

func TestGetTeam(t *testing.T) {
	Convey("valid get", t, func() {
		r, _ := http.NewRequest("GET", "/v1/teams/1", nil)
		r.AddCookie(loginExec())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		checkOK(w)
	})

	Convey("invalid get", t, func() {
		r, _ := http.NewRequest("GET", "/v1/teams/100", nil)
		r.AddCookie(loginExec())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		checkBad(w, http.StatusBadRequest)
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

	Convey("Invalid addition: bad team", t, func() {
		body := bytes.NewReader([]byte(`{"name":"front-end", "description":"we need a front-end dev", "skills":[1,2], "spheres":[1] }`))
		r, _ := http.NewRequest("POST", "/v1/teams/100500/vacancies", body)
		r.AddCookie(loginExec())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		checkBad(w, http.StatusBadRequest)
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

func TestUpdateVacancy(t *testing.T) {
	Convey("valid archive", t, func() {
		body := bytes.NewReader([]byte(`{"field":"status", "value":"archived"}`))
		r, _ := http.NewRequest("PUT", "/v1/teams/1/vacancies/1", body)
		r.AddCookie(loginExec())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		checkOKNoData(w)
	})
}
