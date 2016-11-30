package database

import (
	"testing"

	"github.com/Frozen-Team/summercamp/models"
	"github.com/Frozen-Team/summercamp/tests/setup"
	. "github.com/smartystreets/goconvey/convey"
)

//func TestSkillModel(t *testing.T) {
//	Convey("Test skill model", t, func() {
//		skill := models.Skills.NewSkill("Golang", 1)
//
//		ok := skill.Save()
//
//		So(ok, ShouldBeTrue)
//		So(skill.ID, ShouldNotEqual, 0)
//
//		ok = skill.Delete()
//		So(ok, ShouldBeTrue)
//	})
//}

func TestSkillsAPI(t *testing.T) {
	Convey("Test skills api", t, func() {
		Convey("test fetch by id", func() {
			skill, ok := models.Skills.FetchByID(1)

			So(ok, ShouldBeTrue)
			So(skill, ShouldNotBeNil)
		})

		Convey("test fetch skills by their names", func() {
			Convey("all match list", func() {
				skillNames := []string{"Go", "C"}
				skills, ok := models.Skills.FetchAllByNames(skillNames...)
				So(ok, ShouldBeTrue)
				So(skills, ShouldHaveLength, len(skillNames))
			})

			Convey("1 match", func() {
				skillNames := []string{"Go", "PHP"}
				skills, ok := models.Skills.FetchAllByNames(skillNames...)
				So(ok, ShouldBeTrue)
				So(skills, ShouldHaveLength, 1)
			})

			Convey("no match", func() {
				skillNames := []string{"JS", "PHP"}
				skills, ok := models.Skills.FetchAllByNames(skillNames...)
				So(ok, ShouldBeTrue)
				So(skills, ShouldHaveLength, 0)

			})

			Convey("invalid list: empty list", func() {
				skillNames := []string{}
				skills, ok := models.Skills.FetchAllByNames(skillNames...)
				So(ok, ShouldBeFalse)
				So(skills, ShouldBeNil)
			})
		})

		Convey("test fetch skills by sphere", func() {
			skills, ok := models.Skills.FetchAllBySphere(1)
			So(ok, ShouldBeTrue)
			So(skills, ShouldHaveLength,
				setup.GetFixture("skills").Filter("sphere_id", setup.Equal, "1").Count())
		})
	})
}
