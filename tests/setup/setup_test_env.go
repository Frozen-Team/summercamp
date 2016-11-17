package setup

import (
	"path/filepath"
	"runtime"

	"os"

	"github.com/Frozen-Team/summercamp/models"
	"github.com/astaxie/beego"
	"strconv"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath := filepath.Join(filepath.Dir(file), "../..")

	isTravisBuild, err := strconv.ParseBool(os.Getenv("TRAVIS_BUILD"))
	if err != nil {
		isTravisBuild = false
	}
	if isTravisBuild {
		os.Setenv("BEEGO_RUNMODE", "travis")
	}else{
		os.Setenv("BEEGO_RUNMODE", "test")
	}
	err = beego.LoadAppConfig("ini", filepath.Join(apppath, "conf", "app.conf"))
	if err != nil {
		beego.BeeLogger.Error("failed to load app config \"db.conf\". Error: %v", err)
		os.Exit(1)
	}
	beego.TestBeegoInit(apppath)
	beego.BConfig.WebConfig.EnableXSRF = false

	beego.AddAPPStartHook(func() error {
		models.InitDB()
		return nil
	})
	PrepareTestDB()
}
