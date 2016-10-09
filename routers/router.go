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
		beego.NSNamespace("/users",
			beego.NSInclude(&controllers.Users{}),
		),
		beego.NSNamespace("/teams",
			beego.NSInclude(&controllers.Teams{}),
		),
		beego.NSNamespace("/api",
			beego.NSInclude(&controllers.Api{}),
		),
		beego.NSNamespace("/projects",
			beego.NSInclude(&controllers.Projects{}),
		),
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
		beego.NSRouter("/update_password", &controllers.Users{}, "post:UpdatePassword"),
		beego.NSRouter("/update_email", &controllers.Users{}, "post:UpdateEmail"),
		beego.NSRouter("/:id/skills", &controllers.Users{}, "get:GetSkills"),
	))

	beego.AddNamespace(beego.NewNamespace(
		"/v1/teams",
		beego.NSRouter("", &controllers.Teams{}, "post:Register"),
		beego.NSRouter("/:id", &controllers.Teams{}, "delete:Delete"),
		beego.NSRouter("/:id", &controllers.Teams{}, "get:GetTeam"),
		beego.NSRouter("/:id/members", &controllers.Teams{}, "post:AddMember"),
	))

	beego.AddNamespace(beego.NewNamespace(
		"/v1/api",
		beego.NSRouter("/ping", &controllers.Api{}, "get:Ping"),
	))
	beego.AddNamespace(beego.NewNamespace(
		"/v1/api/projects",
		beego.NSRouter("/", &controllers.Projects{}, "post:Save"),
	))
	beego.AddNamespace(beego.NewNamespace(
		"/v1/user_spheres",
		beego.NSRouter("", &controllers.UserSpheres{}, "post:Save"),
	))
}
