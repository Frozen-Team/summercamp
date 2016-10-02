package models

import (
	"bitbucket.org/SummerCampDev/summercamp/models/utils"
	"github.com/astaxie/beego"
)

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

type UserSpheresAPI struct{}

var UserSpheres *UserSpheresAPI

// SaveSpheresForUser create a new UserSphere record for each sphereID from sphereIDs and userID pair.
// If each record is successfully saved to the db, the func return false
func (us *UserSpheresAPI) SaveSpheresForUser(userID int, sphereIDs ...int) bool {
	if len(sphereIDs) == 0 {
		beego.BeeLogger.Warning("Empty spheres list is passed to SaveSpheresForUser")
		return false
	}

	var failedSpheres []int
	for _, sphereID := range sphereIDs {
		userSphere := UserSphere{
			UserID:   userID,
			SphereID: sphereID,
		}
		if ok := userSphere.Save(); !ok {
			failedSpheres = append(failedSpheres, sphereID)
		}
	}
	ok := len(failedSpheres) == 0
	if !ok {
		beego.BeeLogger.Warning("Failed to save user spheres for spheres with ids: '%v'", failedSpheres)
	}
	return ok
}

// FetchSpheresByUser fetch all spheres for a given user
func (us *UserSpheresAPI) FetchSpheresByUser(userID int) ([]Sphere, bool) {
	var spheres []Sphere
	_, err := DB.Raw(`
	SELECT spheres.id,
	       spheres.name
	FROM user_spheres us
	LEFT OUTER JOIN spheres ON spheres.id=us.sphere_id
	WHERE us.user_id=$1;`, userID).QueryRows(&spheres)
	return spheres, utils.ProcessError(err, " fetch spheres by a user id")
}

// FetchUsersBySpheres fetch all users for a given sphere id
func (us *UserSpheresAPI) FetchUsersBySphere(sphereID int) ([]User, bool) {
	var users []User
	_, err := DB.Raw(`
	SELECT users.id,
	       users.type,
	       users.first_name,
	       users.last_name,
	       users.balance,
	       users.bid,
	       users.braintree_id,
	       users.country,
	       users.city,
	       users.timezone,
	       users.create_time,
	       users.update_time
	FROM user_spheres us
	LEFT OUTER JOIN users ON users.id=us.user_id
	WHERE us.sphere_id=$1;`, sphereID).QueryRows(&users)
	return users, utils.ProcessError(err, " fetch users by a sphere id")
}
