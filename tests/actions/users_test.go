package actions

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "bitbucket.org/SummerCampDev/summercamp/routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRegistrationAction(t *testing.T) {
	Convey("Test registration action", t, func() {
		body := bytes.NewReader([]byte(`{"email":"mail@mail.com", "type":"manager", "first_name":"oleh", "last_name":"gol", "password":"1235~", "password_confirm":"1235~", "country":"ukraine", "city":"kyiv"}`))
		r, _ := http.NewRequest("POST", "/users", body)
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)
		beego.Trace("testing", "TestRegistrationAction", "Code[%d]\n%s", w.Code, w.Body.String())

		So(w.Code, ShouldEqual, 200)
		So(w.Body.Len(), ShouldBeGreaterThan, 0)

	})
}
