package models

import "bitbucket.org/SummerCampDev/summercamp/models/utils"

type Sphere struct {
	ID   int    `json:"id" orm:"column(id)"`
	Name string `json:"name" orm:"column(name)"`
}

func (s *Sphere) TableName() string {
	return "spheres"
}

// Save insert a new record to the db if ID field is of default value. Otherwise an existing
// record is updated.
func (s *Sphere) Save() bool {
	var err error
	var action string

	if s.ID == 0 {
		_, err = DB.Insert(s)
		action = "create"
	} else {
		_, err = DB.Update(s)
		action = "update"
	}

	return utils.ProcessError(err, action+" sphere")
}

// Delete deletes a record from the db. If the record is successfully deleted, the return value
// is true, false - otherwise.
func (s *Sphere) Delete() bool {
	if s.ID == 0 {
		return false
	}
	_, err := DB.Delete(s)

	return utils.ProcessError(err, " delete sphere")
}
