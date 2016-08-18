package models

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/lib/pq" // this import is used by orm.RegisterDataBase
)

var o orm.Ormer // postgres orm

func init() {
	dbUser := beego.AppConfig.String("db_user")
	dbPassword := beego.AppConfig.String("db_password")
	dbHost := beego.AppConfig.String("db_host")
	dbName := beego.AppConfig.String("db_name")
	sslMode := beego.AppConfig.String("ssl_mode")

	dbInfo := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbName, sslMode)

	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		beego.BeeLogger.Error("failed to register postgres driver. Error: %s", err)
		os.Exit(1)
	}

	err = orm.RegisterDataBase("default", "postgres", dbInfo)
	if err != nil {
		beego.BeeLogger.Error("failed to register postgres database. Error: %s", err)
		os.Exit(1)
	}

	registerModels()

	o = orm.NewOrm()
}

// all models will be registered here
func registerModels() {

}
