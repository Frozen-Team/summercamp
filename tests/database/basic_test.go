package database

import (
	"os"
	"testing"

	_ "bitbucket.org/SummerCampDev/summercamp/tests/setup"
	_ "github.com/astaxie/beego"

	"bitbucket.org/SummerCampDev/summercamp/models"
	"bitbucket.org/SummerCampDev/summercamp/tests/setup"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	models.InitDB()
	setup.PrepareTestDB()
	os.Exit(m.Run())
}

func TestUserModel(t *testing.T) {
	Convey("Test User model", t, func() {
		Convey("Test FetchAll", func() {
			c := setup.GetFixture("users").Count()
			users, ok := models.Users.FetchAll()
			So(ok, ShouldBeTrue)
			So(users, ShouldHaveLength, c)
		})
	})
}
