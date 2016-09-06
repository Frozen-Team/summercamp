package database

import (
	"os"
	"testing"

	_ "bitbucket.org/SummerCampDev/summercamp/tests/setup"
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"bitbucket.org/SummerCampDev/summercamp/models"
)

var DB orm.Ormer

func TestMain(m *testing.M) {
	DB = models.InitDB()
	os.Exit(m.Run())
}
