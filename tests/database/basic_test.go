package database

import (
	"os"
	"testing"

	_ "bitbucket.org/SummerCampDev/summercamp/tests/setup"
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"bitbucket.org/SummerCampDev/summercamp/models"
	"bitbucket.org/SummerCampDev/summercamp/tests/setup"
	. "github.com/smartystreets/goconvey/convey"
)

var DB orm.Ormer

func TestMain(m *testing.M) {
	models.InitDB()
	setup.PrepareTestDB()
	os.Exit(m.Run())
}

func TestUserModel(t *testing.T) {
	Convey("Test User model", t, func() {
		Convey("Test FetchAll", func() {
			_, ok := models.Users.FetchByID(1)
			So(ok, ShouldBeTrue)
		})
	})
}
