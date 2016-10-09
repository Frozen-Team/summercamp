package database

import (
	"testing"

	"bitbucket.org/SummerCampDev/summercamp/models"
	"bitbucket.org/SummerCampDev/summercamp/tests/setup"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTeamMembersAPI(t *testing.T) {
	Convey("Test team members api", t, func() {
		Convey("Test fetch teams by id", func() {
			teams, ok := models.TeamMembers.FetchTeamsByMember(1)
			So(ok, ShouldBeTrue)
			So(teams, ShouldHaveLength, setup.GetFixture("team_members").Filter("team_id", setup.Equal, "1").Count())
		})

		Convey("Test fetch all", func() {
			teams, ok := models.TeamMembers.FetchAll()
			So(ok, ShouldBeTrue)
			So(teams, ShouldHaveLength, setup.GetFixture("team_members").Count())
		})

		Convey("Test fetch members by team", func() {
			members, ok := models.TeamMembers.FetchUsersByTeam(1)
			So(ok, ShouldBeTrue)
			So(members, ShouldHaveLength, setup.GetFixture("team_members").Filter("user_id", setup.Equal, "1").Count())
		})

		Convey("Test fetch by id", func() {
			teamMember, ok := models.TeamMembers.FetchByID(1)
			So(ok, ShouldBeTrue)
			So(teamMember, ShouldNotBeNil)
		})
	})
}

func TestTeamMemberSave(t *testing.T) {
	Convey("Test team member model", t, func() {
		teamMember := models.TeamMember{UserID: 3, TeamID: 3}

		Convey("Test Save and Delete", func() {
			ok := teamMember.Save()

			So(ok, ShouldBeTrue)
			So(teamMember.ID, ShouldNotEqual, 0)

			ok = teamMember.Delete()
			So(ok, ShouldBeTrue)
		})
	})
}
