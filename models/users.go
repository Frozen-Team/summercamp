package models

import (
	"time"

	"fmt"

	"bitbucket.org/SummerCampDev/summercamp/models/utils"
)

type Speciality string

const (
	SpecTypeManager  Speciality = "manager"
	SpecTypeClient   Speciality = "client"
	SpecTypeExecutor Speciality = "executor"
)

func (s Speciality) Valid() bool {
	return s == SpecTypeManager || s == SpecTypeClient || s == SpecTypeExecutor
}

type User struct {
	ID           int        `orm:"column(id)"`
	Type         Speciality `orm:"column(type)"`
	FirstName    string     `orm:"column(first_name)"`
	LastName     string     `orm:"column(last_name)"`
	Email        string     `orm:"column(email)"`
	Password     string     `orm:"column(password)" json:"-"`
	PasswordSalt string     `orm:"column(password_salt)" json:"-"`
	Balance      int        `orm:"column(balance)"`
	Bid          int        `orm:"column(bid)"`
	BraintreeID  string     `orm:"column(braintree_id)"`
	Country      string     `orm:"column(country)"`
	City         string     `orm:"column(city)"`
	Timezone     int        `orm:"column(timezone)"`
	CreateTime   time.Time  `orm:"column(create_time)"`
	UpdateTime   time.Time  `orm:"column(update_time)"`
}

// TableName specify the table name for User model. This name is used in the orm RegisterModel
func (u *User) TableName() string {
	return "users"
}

// Save creates a user record in the db
func (u *User) Save() bool {
	var err error
	var action string

	if u.ID == 0 {
		_, err = DB.Insert(u)
		action = "create"
	} else {
		_, err = DB.Update(u)
		action = "update"
	}
	return utils.ProcessError(err, action+" user")
}

// Delete deletes the user record from the db
func (u *User) Delete() bool {
	_, err := DB.Delete(u)
	return utils.ProcessError(err, "delete user")
}

// UserGetByID fetch the user from users table by id
func (u *User) FetchByID(id int) (*User, bool) {
	user := User{ID: id}
	err := DB.Read(&user)
	return &user, utils.ProcessError(err, "fetch the user by id")
}

// UserGetAll fetches all users from the users table
func (u *User) FetchAll() ([]User, bool) {
	var users []User
	_, err := DB.QueryTable(UserObj).All(&users)
	return users, utils.ProcessError(err, "fetch all users")
}

// FetchAllByType fetch users from users table by speciality
func (u *User) FetchAllByType(s Speciality) ([]User, bool) {
	if !s.Valid() {
		return nil, utils.ProcessError(fmt.Errorf("Not valid Speciality '%s'", s), "fetch users by type")
	}
	result := []User{}
	_, err := DB.QueryTable(UserObj).Filter("User__Type__exact", s).All(&result)
	return result, utils.ProcessError(err, "fetch users by type")
}
