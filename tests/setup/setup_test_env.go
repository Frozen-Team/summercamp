package setup

import (
	"path/filepath"
	"runtime"

	"log"

	"github.com/astaxie/beego"
)

func init() {
	_, file, _, _ := runtime.Caller(1)

	apppath := filepath.Join(filepath.Dir(file), "../..")
	beego.TestBeegoInit(apppath)
	err := beego.LoadAppConfig("ini", filepath.Join(apppath, "conf", "db.conf"))
	if err != nil {
		log.Fatal(err)
	}
}
