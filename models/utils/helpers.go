package utils

import "github.com/astaxie/beego"

//TODO: maybe add additional meta data, which may be useful in the log
func ProcessError(err error, action string) bool {
	if err != nil {
		beego.BeeLogger.Error("failed to "+action+". Error: %s", err)
		return false
	}
	return true
}
