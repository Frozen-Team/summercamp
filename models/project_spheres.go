package models

import "bitbucket.org/SummerCampDev/summercamp/models/utils"

type ProjectSphere struct {
	ID        int `json:"id" orm:"column(id)"`
	ProjectID int `json:"project_id" orm:"column(project_id)"`
	SphereID  int `json:"sphere_id" orm:"column(sphere_id)"`
}

// TableName specify the table name for ProjectSphere model. This name is used in the orm RegisterModel
func (ps *ProjectSphere) TableName() string {
	return "project_spheres"
}

// Save insert a new record to the db if ID field is of default value. Otherwise an existing
// record is updated.
func (ps *ProjectSphere) Save() bool {
	var err error
	var action string

	if ps.ID == 0 {
		_, err = DB.Insert(ps)
		action = "create"
	} else {
		_, err = DB.Update(ps)
		action = "update"
	}

	return utils.ProcessError(err, action+" a project`s sphere")
}

// Delete deletes a record from the db. If the record is successfully deleted, the return value
// is true, false - otherwise.
func (ps *ProjectSphere) Delete() bool {
	if ps.ID == 0 {
		return false
	}
	_, err := DB.Delete(ps)

	return utils.ProcessError(err, " delete a project`s sphere")
}
