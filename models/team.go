package models

import (
	"time"

	"bitbucket.org/SummerCampDev/summercamp/models/utils"
)

type Team struct {
	ID          int       `json:"id" orm:"column(id)"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreateTime  time.Time `json:"create_time"`
}

// TableName specify the table name for Team model. This name is used in the orm RegisterModel
func (t *Team) TableName() string {
	return "teams"
}

func (t *Team) Save() bool {
	var err error
	var action string

	if t.ID == 0 {
		t.CreateTime = time.Now()
		_, err = DB.Insert(t)
		action = "create"
	} else {
		_, err = DB.Update(t)
		action = "update"
	}

	return utils.ProcessError(err, action+" team")
}

type teamsAPI struct{}

var Teams *teamsAPI

// FetchByID fetch a team from the teams table by id
func (t *teamsAPI) FetchByID(id int) (*Team, bool) {
	team := Team{ID: id}
	err := DB.Read(&team)
	return &team, utils.ProcessError(err, "fetch the team by id")
}

// FetchAll fetches all teams from the users table
func (t *teamsAPI) FetchAll() ([]Team, bool) {
	var teams []Team
	_, err := DB.QueryTable(TeamObj).All(&teams)
	return teams, utils.ProcessError(err, "fetch all teams")
}
