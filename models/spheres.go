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

// Projects is a wrapper for a method of ProjectSpheresAPI to fetch projects by a sphere id.
func (s *Sphere) Projects() ([]Project, bool) {
	return ProjectSpheres.FetchProjectsBySphere(s.ID)
}

// spheresAPI is an empty struct which is a receiver of helper methods
// which can be useful while working with Sphere model and are not directly relate to it
type spheresAPI struct{}

// Spheres is an object via which we can access helper methods for the Sphere model
var Spheres *spheresAPI

// FetchByID fetch a sphere from the spheres table by id
func (s *spheresAPI) FetchByID(id int) (*Sphere, bool) {
	sphere := Sphere{ID: id}
	err := DB.Read(&sphere)
	return &sphere, utils.ProcessError(err, "fetch the sphere by id")
}

// NewSphere is a wrapper to initialize a new sphere object
func (s *spheresAPI) NewSphere(name string) *Sphere {
	return &Sphere{
		Name: name,
	}
}
