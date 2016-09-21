package database

import (
	"testing"

	"bitbucket.org/SummerCampDev/summercamp/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSkillModel(t *testing.T) {
	Convey("Test skill model", t, func() {
		skill := models.Skills.NewSkill("Golang")

		ok := skill.Save()

		So(ok, ShouldBeTrue)
		So(skill.ID, ShouldNotEqual, 0)
	})
}

func TestSkillsAPI(t *testing.T) {
	Convey("Test skills api", t, func() {
		Convey("test fetch by id", func() {
			skill, ok := models.Skills.FetchByID(1)

			So(ok, ShouldBeTrue)
			So(skill, ShouldNotBeNil)
		})
	})
}
