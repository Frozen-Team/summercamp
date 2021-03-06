package database

import (
	"testing"

	"github.com/Frozen-Team/summercamp/models"
	"github.com/Frozen-Team/summercamp/tests/setup"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTeamModel(t *testing.T) {
	Convey("Test fetch by id", t, func() {
		team, ok := models.Teams.FetchByID(1)
		So(ok, ShouldBeTrue)
		So(team, ShouldNotBeNil)
		So(team.Description, ShouldBeEmpty)
	})

	Convey("Test fetch all", t, func() {
		teams, ok := models.Teams.FetchAll()
		So(ok, ShouldBeTrue)
		So(teams, ShouldHaveLength, setup.GetFixture("teams").Count())
	})
}
