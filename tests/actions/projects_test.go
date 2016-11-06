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

func TestProjectSave(t *testing.T) {
	Convey("Valid save", t, func() {
		body := bytes.NewReader([]byte(`{"description":"a good project", "budget":100500, "sphere_skills":[{"sphere":1, "skills":[1, 2]}]}`))
		r, _ := http.NewRequest("POST", "/v1/projects", body)
		r.AddCookie(loginClient())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		checkOK(w)
	})

	Convey("Invalid save: executor login", t, func() {
		body := bytes.NewReader([]byte(`{"description":"a good project", "budget":100500, "sphere_skills":[{"sphere":1, "skills":[1, 2]}]}`))
		r, _ := http.NewRequest("POST", "/v1/projects", body)
		r.AddCookie(loginExec())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		checkBad(w, http.StatusForbidden)
	})

	Convey("Invalid save: missing field", t, func() {
		body := bytes.NewReader([]byte(`{"budget":100500, "sphere_skills":[{"sphere":1, "skills":[1, 2]}]}`))
		r, _ := http.NewRequest("POST", "/v1/projects", body)
		r.AddCookie(loginClient())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		checkBad(w, http.StatusBadRequest)
	})
}
