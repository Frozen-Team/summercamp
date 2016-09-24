package routers

import (
	"bitbucket.org/SummerCampDev/summercamp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	usersNS := beego.NewNamespace("/users",
		beego.NSRouter("", &controllers.Users{}, "post:Register"),
		beego.NSRouter("/current", &controllers.Users{}, "get:Current"),
	)

	beego.AddNamespace(usersNS)
}
