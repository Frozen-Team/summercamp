package models

import (
	"time"
	"bitbucket.org/SummerCampDev/summercamp/models/utils"
)

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

func (v *Vacancy) Save(columnToUpdate ...string) string {
	_, err := DB.InsertOrUpdate(v,columnToUpdate...)
	return utils.ProcessError(err, "insert of update a vacancy")
}
