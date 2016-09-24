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
	var bnn beegoNewNamespace
	bnn = beego.NewNamespace

	goto swaggerAfterFucked
	// Here is a place where beego's swagger being extremely FUCKED UP
	// Add your fucking things (other controllers) below
	_ = beego.NewNamespace("/v1/users",
		beego.NSInclude(&controllers.Users{}))
	swaggerAfterFucked:

	usersNS := bnn("/v1/users",
		beego.NSRouter("", &controllers.Users{}, "post:Register"),
		beego.NSRouter("/login", &controllers.Users{}, "post:Login"),
		beego.NSRouter("/logout", &controllers.Users{}, "post:Logout"),
	)

	beego.AddNamespace(usersNS)
}
