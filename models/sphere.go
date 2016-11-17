package models

import "github.com/Frozen-Team/summercamp/models/utils"

type Sphere struct {
	ID   int    `json:"id" orm:"column(id)"`
	Name string `json:"name" orm:"column(name)"`
}

func (s *Sphere) TableName() string {
	return "spheres"
}

// Save inserts a new or updates an existing sphere record in the DB.
func (s *Sphere) Save() bool {
	_, err := DB.InsertOrUpdate(s)
	return utils.ProcessError(err, "insert or update sphere")
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
