package models

import (
	"time"

	"bitbucket.org/SummerCampDev/summercamp/models/utils"
	"errors"
)

type Team struct {
	ID          int       `json:"id" orm:"column(id)"`
	Name        string    `json:"name" orm:"column(name)"`
	Description string    `json:"description" orm:"column(description)"`
	CreateTime  time.Time `json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
}

// TableName specify the table name for Team model. This name is used in the orm RegisterModel
func (t *Team) TableName() string {
	return "teams"
}

func (t *Team) Save() bool {
	var err error
	var action string

	if t.ID == 0 {
		_, err = DB.Insert(t)
		action = "create"
	} else {
		_, err = DB.Update(t)
		action = "update"
	}

	return utils.ProcessError(err, action + " team")
}

func (t *Team) Delete() bool {
	if t.ID == 0 {
		return utils.ProcessError(errors.New("Remove unexisting team."), "delete team")
	} else {
		_, err := DB.Delete(t)
		return utils.ProcessError(err, "delete team")
	}
}

// Members return all members` ids for the current team. If there is no error, the second
// return value is true, false - otherwise
func (t *Team) Members() ([]User, bool) {
	return TeamMembers.FetchMembersByTeam(t.ID)
}

// AddMember returns true if there is no errors, false - otherwise
func (t *Team) AddMember(u *User, l AccessLevel) (bool) {
	if t.ID == 0 {
		return utils.ProcessError(errors.New("Cannot add member to unexisting team"), "add member to the team")
	}
	teamMember := TeamMember{
		TeamID: t.ID,
		UserID: u.ID,
		Access: l}
	return teamMember.Save()
}

// DeleteMember returns true if there is no errors, false - otherwise
func (t *Team) DeleteMember(u *User, l AccessLevel) (bool) {
	teamMember := TeamMember{
		TeamID: t.ID,
		UserID: u.ID,
		Access: l}
	return teamMember.Save()
}

type teamsAPI struct{}

var Teams *teamsAPI

// FetchByID fetch a team from the teams table by id
func (t *teamsAPI) FetchByID(id int) (*Team, bool) {
	team := Team{ID: id}
	err := DB.Read(&team)
	return &team, utils.ProcessError(err, "fetch the team by id")
}

// FetchAll fetches all teams from the teams table
func (t *teamsAPI) FetchAll() ([]Team, bool) {
	var teams []Team
	_, err := DB.QueryTable(TeamModel).All(&teams)
	return teams, utils.ProcessError(err, "fetch all teams")
}
