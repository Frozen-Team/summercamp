package database

import (
	"os"
	"testing"

	"bitbucket.org/SummerCampDev/summercamp/tests/setup"
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "github.com/smartystreets/goconvey/convey"

	"bitbucket.org/SummerCampDev/summercamp/models"
)

var DB orm.Ormer

func TestMain(m *testing.M) {
	DB = models.InitDB()
	os.Exit(m.Run())
}

func TestDefault(t *testing.T) {
	setup.PrepareTestDB()
	Convey("adsgsdga", t, func() {
		So(1, ShouldEqual, 2)
	})
}
