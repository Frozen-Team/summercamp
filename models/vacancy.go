package models

import (
	"time"

	"bitbucket.org/SummerCampDev/summercamp/models/utils"
)

type VacancyStatus string

const (
	VacancyStatusActive VacancyStatus = "active"
	VacancyStatusArchived VacancyStatus = "archived"
)

// Vacancy is a model to represent a vacancy of a team
type Vacancy struct {
	ID          int       `json:"id" orm:"column(id)"`
	Name        string    `json:"name" orm:"column(name)"`
	Description string    `json:"description" orm:"column(description)"`
	TeamID      int       `json:"team_id" orm:"column(team_id)"`
	Status VacancyStatus `json:"status" orm:"column(status);"`
	Published   time.Time `json:"published" orm:"column(published);auto_now;type(datetime)"`
}

func (v *Vacancy) TableName() string {
	return "vacancies"
}

// Save inserts a new or updates an existing vacancy record in the DB.
func (v *Vacancy) Save(columnToUpdate ...string) bool {
	var err error
	var action string

	if v.ID == 0 {
		_, err = DB.Insert(v)
		action = "create"
	} else {
		_, err = DB.Update(v, columnToUpdate...)
		action = "update"
	}

	return utils.ProcessError(err, action+" a vacancy")
}

// Activate set the status of the vacancy as "archive".
func (v *Vacancy) Archive() bool {
	v.Status = VacancyStatusArchived
	return v.Save("status")
}

// Activate set the status of the vacancy as "active" and update the "published" column to the current time.
func (v *Vacancy) Activate() bool {
	v.Status = VacancyStatusActive
	v.Published = time.Now()
	return v.Save("status", "published")
}

// Delete deletes the vacancy record from the db
func (v *Vacancy) Delete() bool {
	_, err := DB.Delete(v)
	return utils.ProcessError(err, " delete vacancy")
}

type vacanciesAPI struct{}

var Vacancies *vacanciesAPI

// NewByID is a helper to create a new vacancy with the specified id.
func (v *vacanciesAPI) NewByID(id int) *Vacancy {
	return &Vacancy{
		ID: id,
	}
}

// Archive is a wrapper to archive the vacancy specified with the id.
func (v *vacanciesAPI) Archive(id int) bool {
	return v.NewByID(id).Archive()
}

// Activate is a wrapper to activate the vacancy specified with the id.
func (v *vacanciesAPI) Activate(id int) bool {
	return v.NewByID(id).Activate()
}