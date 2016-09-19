package routers

import (
	"bitbucket.org/SummerCampDev/summercamp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	usersNS := beego.NewNamespace("/users",
		beego.NSRouter("", &controllers.Users{}, "post:Register"),
		beego.NSRouter("/login", &controllers.Users{}, "post:Login"),
		beego.NSRouter("/logout", &controllers.Users{}, "post:Logout"),
	)

	beego.AddNamespace(usersNS)
}
