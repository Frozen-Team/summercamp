package models

import (
	"time"

	"bitbucket.org/SummerCampDev/summercamp/models/utils"
)

// Vacancy is a model to represent a vacancy of a team
type Vacancy struct {
	ID          int       `json:"id" orm:"column(id)"`
	Name        string    `json:"name" orm:"column(name)"`
	Description string    `json:"description" orm:"column(description)"`
	TeamID      int       `json:"team_id" orm:"column(team_id)"`
	Published   time.Time `json:"published" orm:"column(published);auto_now;type(datetime)"`
}

func (v *Vacancy) TableName() string {
	return "vacancies"
}

// Save inserts a new or updates an existing vacancy record in the DB.
func (v *Vacancy) Save(columnToUpdate ...string) bool {
	_, err := DB.InsertOrUpdate(v)
	return utils.ProcessError(err, "insert or update vacancy")
}
