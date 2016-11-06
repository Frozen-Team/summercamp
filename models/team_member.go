package models

import (
	"time"

	"bitbucket.org/SummerCampDev/summercamp/models/utils"
)

type AccessLevel int

//TODO: add default value for DB column as readonly
const (
	AccessCreator  AccessLevel = 0
	AccessReadOnly AccessLevel = 100
)

// TeamMember represents a member with UserID who is in the team with TeamID.
type TeamMember struct {
	ID         int         `json:"id" orm:"column(id)"`
	TeamID     int         `json:"team_id" orm:"column(team_id)"`
	UserID     int         `json:"user_id" orm:"column(user_id)"`
	Access     AccessLevel `json:"access" orm:"column(access)"`
	CreateTime time.Time   `json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
}

// TableName specify the table name for Team model. This name is used in the orm RegisterModel.
func (tm *TeamMember) TableName() string {
	return "team_members"
}

// Save insert a new record to the db if ID field is of default value. Otherwise an existing
// record is updated.
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

// Delete deletes a team member record from the db. If the record is successfully deleted, the return value
// is true, false - otherwise.
func (tm *TeamMember) Delete() bool {
	_, err := DB.Delete(tm)
	return utils.ProcessError(err, " delete team member")
}

// IsCreator is a helper to check if the current team member is a creator of the team.
func (tm *TeamMember) IsCreator() bool {
	return tm.Access == AccessCreator
}

// teamsMembersAPI is an empty struct which is a receiver of helper methods
// which can be useful while working with TeamMember model and are not directly relate to it
type teamsMembersAPI struct{}

// TeamMembers is an object via which we can access helper methods for the TeamMember model
var TeamMembers *teamsMembersAPI

// FetchByUserId searches for TeamMember by user id
func (tm *TeamMember) FetchByUserId(userId int) (*TeamMember, bool) {
	m := TeamMember{UserID: userId}
	err := DB.Read(&m)

	return &m, utils.ProcessError(err, " fetch team member by user id")
}

// FetchByID fetch a team from the teams table by id
func (tm *teamsMembersAPI) FetchByID(id int) (*TeamMember, bool) {
	teamMember := TeamMember{ID: id}
	err := DB.Read(&teamMember)
	return &teamMember, utils.ProcessError(err, "fetch the team member by id")
}

// FetchAll fetches all teams from the team_members table
func (tm *teamsMembersAPI) FetchAll() ([]TeamMember, bool) {
	var teamMembers []TeamMember
	_, err := DB.QueryTable(TeamMemberModel).All(&teamMembers)
	return teamMembers, utils.ProcessError(err, "fetch all team members")
}

// FetchTeamsByMember fetch all teams in which the user with the teamMemberID is a member
func (tm *teamsMembersAPI) FetchTeamsByMember(teamMemberID int) ([]Team, bool) {
	var teamMembers []TeamMember
	_, err := DB.QueryTable(TeamMemberModel).Filter("user_id", teamMemberID).All(&teamMembers)
	if err != nil {
		return nil, utils.ProcessError(err, "fetch teamMembers by member id")
	}
	teamIDs := make([]int, len(teamMembers))
	for _, teamMember := range teamMembers {
		teamIDs = append(teamIDs, teamMember.TeamID)
	}
	var teams []Team
	_, err = DB.QueryTable(TeamModel).Filter("id__in", teamIDs).All(&teams)
	return teams, utils.ProcessError(err, "fetch teams by teamIDs")
}

// FetchMembersByTeam fetch all members of the team with the given teamID
func (tm *teamsMembersAPI) FetchUsersByTeam(teamID int) ([]User, bool) {
	var teamMembers []TeamMember
	_, err := DB.QueryTable(TeamMemberModel).Filter("team_id", teamID).All(&teamMembers)
	if err != nil {
		return nil, utils.ProcessError(err, "fetch members by team")
	}
	userIDs := make([]int, len(teamMembers))
	for _, teamMember := range teamMembers {
		userIDs = append(userIDs, teamMember.UserID)
	}
	var users []User
	_, err = DB.QueryTable(UserModel).Filter("id__in", userIDs).All(&users)
	return users, utils.ProcessError(err, "fetch members by userIDs")
}

// FetchMembersByTeam fetch all members of the team with the given teamID
func (tm *teamsMembersAPI) FetchTeamMembersByTeam(teamID int) ([]TeamMember, bool) {
	var teamMembers []TeamMember
	_, err := DB.QueryTable(TeamMemberModel).Filter("team_id", teamID).All(&teamMembers)
	if err != nil {
		return nil, utils.ProcessError(err, "fetch members by team")
	}
	return teamMembers, true
}
