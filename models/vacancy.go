package models

import (
	"time"
	"bitbucket.org/SummerCampDev/summercamp/models/utils"
)

// Vacancy is a model to represent a vacancy of a team
type Vacancy struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	TeamID int `json:"team_id"`
	Published time.Time `json:"published"`
}

func (v *Vacancy) TableName() string {
	return "vacancy"
}

// Save inserts a new element to the db or update columns columnToUpdate of an existing record.
func (v *Vacancy) Save(columnToUpdate ...string) string {
	_, err := DB.InsertOrUpdate(v, columnToUpdate...)
	return utils.ProcessError(err, "insert of update a vacancy")
}
