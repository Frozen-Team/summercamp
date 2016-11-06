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

// for beego swagger.
type beegoNewNamespace func(prefix string, params ...beego.LinkNamespace) *beego.Namespace

func init() {
	goto swaggerAfterFucked
	// Beego's swagger internal implementation requires such a huck.
	// Add new controllers below.
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

	beego.AddNamespace(beego.NewNamespace("/v1",
		beego.NSNamespace("/users",
			beego.NSRouter("", &controllers.Users{}, "post:Register;put:UpdateField"),
			beego.NSRouter("/current", &controllers.Users{}, "get:Current"),
			beego.NSRouter("/login", &controllers.Users{}, "post:LogIn"),
			beego.NSRouter("/logout", &controllers.Users{}, "post:LogOut"),
			beego.NSRouter("/:id", &controllers.Users{}, "get:GetUser"),
			beego.NSRouter("/update_password", &controllers.Users{}, "post:UpdatePassword"),
			beego.NSRouter("/:id/skills", &controllers.Users{}, "get:GetSkills"),
			beego.NSRouter("/skills", &controllers.Users{}, "post:AddSkill"),
			beego.NSRouter("/skills/:id", &controllers.Users{}, "delete:RemoveSkill"),
			beego.NSRouter("/spheres", &controllers.Users{}, "post:AddSphere"),
			beego.NSRouter("/spheres/:id", &controllers.Users{}, "delete:RemoveSphere"),
		),

		beego.NSNamespace("/teams",
			beego.NSRouter("", &controllers.Teams{}, "post:Save"),
			beego.NSRouter("/:id", &controllers.Teams{}, "get:GetTeam;delete:Delete"),
			beego.NSRouter("/:id/members", &controllers.Teams{}, "post:AddMember"),
			beego.NSRouter("/:id/vacancies", &controllers.Teams{}, "post:AddVacancy"),
			beego.NSRouter("/:id/vacancies/:v_id", &controllers.Teams{}, "delete:RemoveVacancy"),
		),

		beego.NSNamespace(
			"/projects",
			beego.NSRouter("/", &controllers.Projects{}, "post:Save"),
		),

		beego.NSNamespace(
			"/api",
			beego.NSRouter("/ping", &controllers.Api{}, "get:Ping"),
		),
	))
}
