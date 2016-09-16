package models

import (
	"time"

	"bitbucket.org/SummerCampDev/summercamp/models/utils"
)

type TeamMember struct {
	ID         int       `json:"id" orm:"column(id)"`
	TeamID     int       `json:"team_id" orm:"column(team_id)"`
	UserID     int       `json:"user_id" orm:"column(user_id)"`
	CreateTime time.Time `json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
}

// TableName specify the table name for Team model. This name is used in the orm RegisterModel
func (tm *TeamMember) TableName() string {
	return "team_members"
}

func (tm *TeamMember) Save() bool {
	var err error
	var action string

	if tm.ID == 0 {
		_, err = DB.Insert(tm)
		action = "create"
	} else {
		_, err = DB.Update(tm)
		action = "update"
	}

	return utils.ProcessError(err, action+" team member")
}

func (tm *TeamMember) Delete() bool {
	if tm.ID == 0 {
		return false
	}
	_, err := DB.Delete(tm)

	return utils.ProcessError(err, " delete team member")
}

type teamsMembersAPI struct{}

var TeamMembers *teamsMembersAPI

// FetchByID fetch a team from the teams table by id
func (tm *teamsMembersAPI) FetchByID(id int) (*TeamMember, bool) {
	teamMember := TeamMember{ID: id}
	err := DB.Read(&teamMember)
	return &teamMember, utils.ProcessError(err, "fetch the team member by id")
}

// FetchAll fetches all teams from the users table
func (tm *teamsMembersAPI) FetchAll() ([]TeamMember, bool) {
	var teamMembers []TeamMember
	_, err := DB.QueryTable(TeamMemberObj).All(&teamMembers)
	return teamMembers, utils.ProcessError(err, "fetch all team members")
}

// FetchTeamsByMember fetch all teams in which the user with the teamMemberID is a member
func (tm *teamsMembersAPI) FetchTeamsByMember(teamMemberID int) ([]int, bool) {
	var teamIDs []int
	_, err := DB.Raw("SELECT team_id FROM team_members WHERE user_id=$1;", teamMemberID).QueryRows(&teamIDs)
	return teamIDs, utils.ProcessError(err, "fetch teams by member")
}

// FetchMembersByTeam fetch all members of the team with the given teamID
func (tm *teamsMembersAPI) FetchMembersByTeam(teamID int) ([]int, bool) {
	var memberIDs []int
	_, err := DB.Raw("SELECT user_id FROM team_members WHERE team_id=$1", teamID).QueryRows(&memberIDs)
	return memberIDs, utils.ProcessError(err, "fetch members by team")
}
