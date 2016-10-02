package models

import "bitbucket.org/SummerCampDev/summercamp/models/utils"

type UserSphere struct {
	ID       int `json:"id" orm:"id"`
	UserID   int `json:"user_id" orm:"user_id"`
	SphereID int `json:"sphere_id" orm:"sphere_id"`
}

// TableName specify the table name for UserSphere model. This name is used in the orm RegisterModel
func (us *UserSphere) TableName() string {
	return "user_spheres"
}

// Save insert a new record to the db if ID field is of default value. Otherwise an existing
// record is updated.
func (us *UserSphere) Save() bool {
	var err error
	var action string

	if us.ID == 0 {
		_, err = DB.Insert(us)
		action = "create"
	} else {
		_, err = DB.Update(us)
		action = "update"
	}

	return utils.ProcessError(err, action+" a user`s sphere")
}

// Delete deletes a record from the db. If the record is successfully deleted, the return value
// is true, false - otherwise.
func (us *UserSphere) Delete() bool {
	if us.ID == 0 {
		return false
	}
	_, err := DB.Delete(us)

	return utils.ProcessError(err, " delete a user`s sphere")
}
