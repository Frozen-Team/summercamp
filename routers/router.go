package routers

import (
	"bitbucket.org/SummerCampDev/summercamp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	usersNS := beego.NewNamespace("/users",
		beego.NSRouter("", &controllers.UsersController{}, "post:Register"),
	)

	beego.AddNamespace(usersNS)
}
