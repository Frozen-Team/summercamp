package setup

import (
	"path/filepath"
	"runtime"

	"os"

	"bitbucket.org/SummerCampDev/summercamp/models"
	"github.com/astaxie/beego"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath := filepath.Join(filepath.Dir(file), "../..")
	os.Setenv("BEEGO_RUNMODE", "test")
	err := beego.LoadAppConfig("ini", filepath.Join(apppath, "conf", "app.conf"))
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
