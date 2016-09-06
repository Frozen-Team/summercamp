package setup

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gopkg.in/testfixtures.v1"
)

// the path from the project root folder
const FixturesPath = "tests/fixtures"

func PrepareTestDB() {
	db, err := orm.GetDB(beego.AppConfig.String("db.alias_name"))
	if err != nil {
		log.Fatal("GetDB error: " + err.Error())
	}
	err = testfixtures.LoadFixtures(FixturesPath, db, &testfixtures.PostgreSQLHelper{})
	if err != nil {
		log.Fatal("LoadFixtures error: " + err.Error())
	}
}
