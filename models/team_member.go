package models

import (
	"time"

	"bitbucket.org/SummerCampDev/summercamp/models/utils"
)

type TeamMember struct {
	ID         int       `json:"id" orm:"column(id)"`
	TeamID     int       `json:"team_id" orm:"column(team_id)"`
	UserID     int       `json:"user_id" orm:"column(user_id)"`
	CreateTime time.Time `json:"create_time" orm:"column(create_time) auto_now_add;type(datetime)"`
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
