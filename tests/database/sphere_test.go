package database

import (
	"testing"

	"bitbucket.org/SummerCampDev/summercamp/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSphereModel(t *testing.T) {
	Convey("Test sphere model", t, func() {
		sphere := models.Spheres.NewSphere("Desktop")

		ok := sphere.Save()

		So(ok, ShouldBeTrue)
		So(sphere.ID, ShouldNotEqual, 0)

		ok = sphere.Delete()
		So(ok, ShouldBeTrue)
	})
}

func TestSpheresAPI(t *testing.T) {
	Convey("Test spheres api", t, func() {
		Convey("test fetch by id", func() {
			sphere, ok := models.Spheres.FetchByID(1)

			So(ok, ShouldBeTrue)
			So(sphere, ShouldNotBeNil)
		})
	})
}
