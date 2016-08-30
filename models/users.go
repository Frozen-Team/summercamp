package models

import (
	"bitbucket.org/SummerCampDev/summercamp/models/utils"
	"time"
)

type Speciality string

const (
	SpecManager  Speciality = "manager"
	SpecClient   Speciality = "client"
	SpecExecutor Speciality = "executor"
)

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
func UserGetByID(id int) (*User, bool) {
	user := User{ID: id}
	err := DB.Read(&user)
	return &user, utils.ProcessError(err, "fetch the user by id")
}

// UserGetAll fetches all users from the users table
func UserGetAll() ([]User, bool) {
	var users []User
	_, err := DB.QueryTable(UserObj).All(&users)
	return users, utils.ProcessError(err, "fetch all users")
}
