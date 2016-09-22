package database

import (
	"testing"

	"bitbucket.org/SummerCampDev/summercamp/models"
	"bitbucket.org/SummerCampDev/summercamp/tests/setup"
	. "github.com/smartystreets/goconvey/convey"
)

func TestProjectModel(t *testing.T) {
	Convey("Test project model", t, func() {
		project := models.Project{
			Description: "hello",
			Budget:      10,
			ClientID:    1,
		}

		Convey("Test valid save", func() {
			ok := project.Save()

			So(ok, ShouldBeTrue)
			So(project.ID, ShouldNotEqual, 0)

			ok = project.Delete()
			So(ok, ShouldBeTrue)
		})

		Convey("Test negative budget", func() {
			project.Budget = -1

			ok := project.Save()
			So(ok, ShouldBeFalse)
		})
	})
}

func TestProjectAPI(t *testing.T) {
	Convey("Test projects model", t, func() {
		Convey("Test fetch by id", func() {
			project, ok := models.Projects.FetchByID(1)
			So(ok, ShouldBeTrue)
			So(project, ShouldNotBeNil)
		})

		Convey("Test fetch by client", func() {
			projects, ok := models.Projects.FetchAllByClient(1)
			So(ok, ShouldBeTrue)
			So(projects, ShouldHaveLength,
				setup.GetFixture("projects").Filter("client_id", setup.Equal, "1").Count())
		})
	})
}
