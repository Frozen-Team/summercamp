package actions

import (
	"testing"

	_ "bitbucket.org/SummerCampDev/summercamp/routers"
	_ "bitbucket.org/SummerCampDev/summercamp/tests/setup"

	"bytes"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
)

func TestSave(t *testing.T) {
	Convey("Valid save", t, func() {
		body := bytes.NewReader([]byte(`{"description":"a good project", "budget":100500, "sphere_skills":[{"sphere":1, "skills":[1, 2]}]}`))
		r, _ := http.NewRequest("POST", "/v1/projects", body)
		r.AddCookie(loginClient())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		response, err := ReadResponse(w.Body)
		So(err, ShouldBeNil)
		So(response.Meta.HasError, ShouldBeFalse)
	})

	Convey("Invalid save: executor login", t, func() {
		body := bytes.NewReader([]byte(`{"description":"a good project", "budget":100500, "sphere_skills":[{"sphere":1, "skills":[1, 2]}]}`))
		r, _ := http.NewRequest("POST", "/v1/projects", body)
		r.AddCookie(loginExec())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusForbidden)
		response, err := ReadResponse(w.Body)
		So(err, ShouldBeNil)
		So(response.Meta.HasError, ShouldBeTrue)
	})

	Convey("Invalid save: missing field", t, func() {
		body := bytes.NewReader([]byte(`{"budget":100500, "sphere_skills":[{"sphere":1, "skills":[1, 2]}]}`))
		r, _ := http.NewRequest("POST", "/v1/projects", body)
		r.AddCookie(loginClient())
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
		response, err := ReadResponse(w.Body)
		So(err, ShouldBeNil)
		So(response.Meta.HasError, ShouldBeTrue)
	})
}
