package setup

import (
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gopkg.in/testfixtures.v1"
)

// FixturesPath is the path where fixtures files stores
var FixturesPath = "tests/fixtures"

func PrepareTestDB() {
	db, err := orm.GetDB(beego.AppConfig.String("db.alias_name"))
	if err != nil {
		beego.BeeLogger.Error("GetDB error: %v", err)
		os.Exit(1)
	}
	err = testfixtures.LoadFixtures(FixturesPath, db, &testfixtures.PostgreSQLHelper{})
	if err != nil {
		beego.BeeLogger.Error("LoadFixtures error: %v", err)
		os.Exit(1)
	}
}
