package models

import (
	"github.com/Frozen-Team/summercamp/models/utils"
	"github.com/astaxie/beego"
)

type VacancySphere struct {
	ID        int `json:"id" orm:"column(id)"`
	VacancyID int `json:"vacancy_id" orm:"column(vacancy_id)"`
	SphereID  int `json:"sphere_id" orm:"column(sphere_id)"`
}

// TableName specify the table name for VacancySphere model. This name is used in the orm RegisterModel
func (us *VacancySphere) TableName() string {
	return "vacancy_spheres"
}

// Save inserts a new or updates an existing sphere record in the DB.
func (us *VacancySphere) Save() bool {
	_, err := DB.InsertOrUpdate(us)
	return utils.ProcessError(err, "insert or update vacancy sphere")
}

// Delete deletes a record from the db. If the record is successfully deleted, the return value
// is true, false - otherwise.
func (us *VacancySphere) Delete() bool {
	if us.ID == 0 {
		return false
	}
	_, err := DB.Delete(us)

	return utils.ProcessError(err, " delete a vacancy`s sphere")
}

type VacancySpheresAPI struct{}

var VacancySpheres *VacancySpheresAPI

// SaveSpheresForVacancy create a new VacancySphere record for each sphereID from sphereIDs and vacancyID pair.
// If the DB inserter successfully closes, the function returns true, false - otherwise.
func (us *VacancySpheresAPI) SaveSpheresForVacancy(vacancyID int, sphereIDs ...int) bool {
	if len(sphereIDs) == 0 {
		beego.BeeLogger.Warning("Empty spheres list is passed to SaveSpheresForVacancy")
		return false
	}

	i, err := DB.QueryTable(VacancySphereModel).PrepareInsert()
	if err != nil {
		return utils.ProcessError(err, " create an inserter")
	}

	var failedSpheres []int

	for _, sphereID := range sphereIDs {
		_, err = i.Insert(&VacancySphere{
			VacancyID: vacancyID,
			SphereID:  sphereID,
		})
		if err != nil {
			failedSpheres = append(failedSpheres, sphereID)
		}
	}

	ok := len(failedSpheres) == 0
	if !ok {
		beego.BeeLogger.Warning("Failed to save vacancy spheres for spheres with ids: '%v'", failedSpheres)
	}

	err = i.Close()
	return utils.ProcessError(err, " insert multiple vacancy spheres")
}
