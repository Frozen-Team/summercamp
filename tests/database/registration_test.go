package database

import (
	"testing"

	"bitbucket.org/SummerCampDev/summercamp/models"
	"bitbucket.org/SummerCampDev/summercamp/models/forms"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRegistrationForm(t *testing.T) {
	Convey("Test registration form", t, func() {
		ur := forms.UserReg{
			Email:           "valid@mail.com",
			Type:            models.SpecTypeExecutor,
			FirstName:       "oleh",
			LastName:        "holovko",
			Password:        "12345",
			PasswordConfirm: "12345",
			Country:         "country",
			City:            "city",
		}

		Convey("Everything is okay", func() {
			ur.Type = models.SpecTypeExecutor

			user, ok := ur.Register()
			So(user, ShouldNotBeNil)
			So(ok, ShouldBeTrue)
			So(len(ur.Errors), ShouldBeZeroValue)
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
	})

}
