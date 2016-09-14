package database

import (
	"testing"

	"bytes"
	"net/http"
	"net/http/httptest"

	"bitbucket.org/SummerCampDev/summercamp/models"
	"bitbucket.org/SummerCampDev/summercamp/models/forms"
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
		Convey("Subject: Test Station Endpoint\n", func() {
			Convey("Status Code Should Be 200", func() {
				So(w.Code, ShouldEqual, 200)
			})
			Convey("The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
		})
	})
}

func TestRegistrationForm(t *testing.T) {
	Convey("Test registration form", t, func() {
		ur := forms.UserRegistration{
			Email:           "valid@mail.com",
			Type:            models.SpecTypeExecutor,
			FirstName:       "oleh",
			LastName:        "holovko",
			Password:        "1234%",
			PasswordConfirm: "1234%",
			Country:         "country",
			City:            "city",
		}

		Convey("Everything is okay", func() {
			ur.Type = models.SpecTypeExecutor

			user, ok := ur.Register()
			So(user, ShouldNotBeNil)
			So(ok, ShouldBeTrue)
			So(ur.Errors, ShouldHaveLength, 0)
		})

		Convey("Got invalid email", func() {
			ur.Email = "bad_email"
			user, ok := ur.Register()

			So(user, ShouldBeNil)
			So(ok, ShouldBeFalse)

			So(len(ur.Errors), ShouldNotEqual, 0)
		})

		Convey("Password and password confirmation mismatch", func() {
			ur.Password = "1"
			user, ok := ur.Register()

			So(user, ShouldBeNil)
			So(ok, ShouldBeFalse)
			So(len(ur.Errors), ShouldNotEqual, 0)
		})

		Convey("Invalid speciality", func() {
			ur.Type = "bad_spec"
			spec := models.Speciality(ur.Type)

			So(spec.Valid(), ShouldBeFalse)

			user, ok := ur.Register()

			So(user, ShouldBeNil)
			So(ok, ShouldBeFalse)
			So(len(ur.Errors), ShouldNotEqual, 0)
		})

		Convey("Very long password", func() {
			ur.Password = "%1111111111111111111111111111111111111111111111111111111111111111111111111111"
			ur.PasswordConfirm = "%1111111111111111111111111111111111111111111111111111111111111111111111111111"

			user, ok := ur.Register()

			So(user, ShouldBeNil)
			So(ok, ShouldBeFalse)
			So(len(ur.Errors), ShouldNotEqual, 0)
		})
	})

}
