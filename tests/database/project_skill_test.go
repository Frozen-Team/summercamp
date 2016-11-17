package database

import (
	"testing"

	"github.com/Frozen-Team/summercamp/models"
	"github.com/Frozen-Team/summercamp/tests/setup"
	. "github.com/smartystreets/goconvey/convey"
)

func ProjectSkillModelTest(t *testing.T) {
	Convey("Test project skill model", t, func() {
		ps := models.ProjectSkill{
			ProjectID: 2,
			SkillID:   1,
		}

		ok := ps.Save()

		So(ok, ShouldBeTrue)
		So(ps.ID, ShouldNotEqual, 0)

		ok = ps.Delete()
		So(ok, ShouldBeTrue)
	})
}

func ProjectSkillsAPITest(t *testing.T) {
	Convey("Test project skills api", t, func() {
		Convey("save skills for project", func() {
			skills := []int{3, 4}
			ok := models.ProjectSkills.SaveSkillsForProject(2, skills...)

			So(ok, ShouldBeTrue)
		})

		Convey("Fetch skills by project", func() {
			skills, ok := models.ProjectSkills.FetchSkillsByProject(1)

			So(ok, ShouldBeTrue)
			So(skills, ShouldHaveLength,
				setup.GetFixture("project_skills").Filter("project_id", setup.Equal, "1").Count())
		})
		Convey("Fetch projects by skill", func() {
			projects, ok := models.ProjectSkills.FetchProjectsBySkill(1)

			So(ok, ShouldBeTrue)
			So(projects, ShouldHaveLength,
				setup.GetFixture("project_skills").Filter("skill_id", setup.Equal, "1").Count())
		})
	})
}
