package database

import (
	"testing"

	"bitbucket.org/SummerCampDev/summercamp/models"
	"bitbucket.org/SummerCampDev/summercamp/tests/setup"
	. "github.com/smartystreets/goconvey/convey"
)

func TestProjectSphere(t *testing.T) {
	Convey("Test project sphere model", t, func() {
		projectSphere := models.ProjectSphere{
			ProjectID: 2,
			SphereID:  1,
		}

		ok := projectSphere.Save()

		So(ok, ShouldBeTrue)
		So(projectSphere.ID, ShouldNotEqual, 0)

		ok = projectSphere.Delete()

		So(ok, ShouldBeTrue)
	})
}

func TestProjectSpheresAPI(t *testing.T) {
	Convey("Test project spheres api", t, func() {
		Convey("save spheres for project", func() {
			spheres := []int{2}
			ok := models.ProjectSpheres.SaveSpheresForProject(2, spheres...)

			So(ok, ShouldBeTrue)
		})

		Convey("Fetch spheres by project", func() {
			spheres, ok := models.ProjectSpheres.FetchSpheresByProject(1)

			So(ok, ShouldBeTrue)
			So(spheres, ShouldHaveLength,
				setup.GetFixture("project_spheres").Filter("project_id", setup.Equal, "1").Count())
		})
		Convey("Fetch projects by sphere", func() {
			projects, ok := models.ProjectSpheres.FetchProjectsBySphere(1)

			So(ok, ShouldBeTrue)
			So(projects, ShouldHaveLength,
				setup.GetFixture("project_spheres").Filter("sphere_id", setup.Equal, "1").Count())
		})
	})
}
