package models

import (
	"bitbucket.org/SummerCampDev/summercamp/models/utils"
	"github.com/astaxie/beego"
)

type ProjectSphere struct {
	ID        int `json:"id" orm:"column(id)"`
	ProjectID int `json:"project_id" orm:"column(project_id)"`
	SphereID  int `json:"sphere_id" orm:"column(sphere_id)"`
}

// TableName specify the table name for ProjectSphere model. This name is used in the orm RegisterModel
func (ps *ProjectSphere) TableName() string {
	return "project_spheres"
}

// Save inserts a new or updates an existing project's sphere record in the DB.
func (ps *ProjectSphere) Save() bool {
	_, err := DB.InsertOrUpdate(ps)
	return utils.ProcessError(err, "insert or update project's sphere")
}

// Delete deletes a record from the db. If the record is successfully deleted, the return value
// is true, false - otherwise.
func (ps *ProjectSphere) Delete() bool {
	if ps.ID == 0 {
		return false
	}
	_, err := DB.Delete(ps)

	return utils.ProcessError(err, "delete a project`s sphere")
}

type ProjectSpheresAPI struct{}

var ProjectSpheres *ProjectSpheresAPI

// SaveSpheresForProject create a new ProjectSphere record for each sphereID from sphereIDs and projectID pair.
// If each record is successfully saved to the db, the func return false
func (ps *ProjectSpheresAPI) SaveSpheresForProject(projectID int, sphereIDs ...int) bool {
	if len(sphereIDs) == 0 {
		beego.BeeLogger.Warning("Empty spheres list is passed to SaveSpheresForProject")
		return false
	}

	var failedSpheres []int
	for _, sphereID := range sphereIDs {
		projectSphere := ProjectSphere{
			ProjectID: projectID,
			SphereID:  sphereID,
		}
		if ok := projectSphere.Save(); !ok {
			failedSpheres = append(failedSpheres, sphereID)
		}
	}
	ok := len(failedSpheres) == 0
	if !ok {
		beego.BeeLogger.Warning("Failed to save project spheres for spheres with ids: '%v'", failedSpheres)
	}
	return ok
}

// FetchSpheresByProject fetch all spheres for a given project
func (ps *ProjectSpheresAPI) FetchSpheresByProject(projectID int) ([]Sphere, bool) {
	var spheres []Sphere
	_, err := DB.Raw(`
	SELECT spheres.id,
	       spheres.name
	FROM project_spheres ps
	LEFT OUTER JOIN spheres ON spheres.id=ps.sphere_id
	WHERE ps.project_id=$1;`, projectID).QueryRows(&spheres)
	return spheres, utils.ProcessError(err, " fetch spheres by a project id")
}

// FetchProjectsBySpheres fetch all projects for a given sphere id
func (ps *ProjectSpheresAPI) FetchProjectsBySphere(sphereID int) ([]Project, bool) {
	var projects []Project
	_, err := DB.Raw(`
	SELECT projects.id,
	       projects.description,
	       projects.budget,
	       projects.client_id,
	       projects.create_time,
	       projects.update_time
	FROM project_spheres ps
	LEFT OUTER JOIN projects ON projects.id=ps.project_id
	WHERE ps.sphere_id=$1;`, sphereID).QueryRows(&projects)
	return projects, utils.ProcessError(err, " fetch projects by a sphere id")
}
