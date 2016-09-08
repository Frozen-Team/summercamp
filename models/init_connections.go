package models

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/lib/pq" // this import is used by orm.RegisterDataBase
)

var DB orm.Ormer // postgres orm

func init() {
	beego.AddAPPStartHook(func() error {
		InitDB()
		return nil
	})
}

func InitDB(){
	dbDriverName := beego.AppConfig.String("db.driver_name")
	dbAliasName := beego.AppConfig.String("db.alias_name")
	dbUser := beego.AppConfig.String("db.user")
	dbPassword := beego.AppConfig.String("db.password")
	dbHost := beego.AppConfig.String("db.host")
	dbName := beego.AppConfig.String("db.name")
	sslMode := beego.AppConfig.String("db.ssl_mode")

	dbInfo := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbName, sslMode)

	err := orm.RegisterDriver(dbDriverName, orm.DRPostgres)
	if err != nil {
		beego.BeeLogger.Error("failed to register "+dbDriverName+" driver. Error: %s", err)
		os.Exit(1)
	}

	err = orm.RegisterDataBase(dbAliasName, dbDriverName, dbInfo)
	if err != nil {
		beego.BeeLogger.Error("failed to register "+dbDriverName+" database. Error: %s", err)
		os.Exit(1)
	}

	registerModels()
	DB = orm.NewOrm()
}

// all models will be registered here
func registerModels() {
	orm.RegisterModel(UserObj)
}
