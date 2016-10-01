// @APIVersion 1.0.0
// @Title Summer Camp
// @Description Summer Camp is a collaboration platform for freelancers
// @Contact admin@frozen-team.com
// @TermsOfServiceUrl http://frozen-team.com/
// @License Proprietary
// @LicenseUrl http://frozen-team.com/license/
package routers

import (
	"bitbucket.org/SummerCampDev/summercamp/controllers"
	"github.com/astaxie/beego"
)

// Fuck beego swagger
type beegoNewNamespace func(prefix string, params ...beego.LinkNamespace) *beego.Namespace

func init() {
	goto swaggerAfterFucked
	// Here is a place where beego's swagger being extremely FUCKED UP
	// Add your fucking things (other controllers) below
	_ = beego.NewNamespace(
		"/v1",
		beego.NSInclude(&controllers.Users{}),
		beego.NSInclude(&controllers.Teams{}),
	)
	swaggerAfterFucked:

	beego.AddNamespace(beego.NewNamespace(
		"/v1/users",
		beego.NSRouter("", &controllers.Users{}, "post:Register"),
		beego.NSRouter("/current", &controllers.Users{}, "get:Current"),
		beego.NSRouter("/login", &controllers.Users{}, "post:Login"),
		beego.NSRouter("/logout", &controllers.Users{}, "post:Logout"),
		beego.NSRouter("/:id", &controllers.Users{}, "get:GetUser"),
		beego.NSRouter("/update_field", &controllers.Users{}, "post:UpdateField"),
	))

	beego.AddNamespace(beego.NewNamespace(
		"/v1/teams",
		beego.NSRouter("", &controllers.Teams{}, "post:Register"),
		beego.NSRouter("/:id", &controllers.Users{}, "delete:Delete"),
	))
}
