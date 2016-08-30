package models

import (
	"github.com/astaxie/beego"
	"time"
)

type Speciality string

const (
	SpecManager  Speciality = "manager"
	SpecClient   Speciality = "client"
	SpecExecutor Speciality = "executor"
)

type User struct {
	Id           int        `orm:"column(id)"`
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
	CreateTime   time.Time  `orm:"column(create_time)"`
	UpdateTime   time.Time  `orm:"column(update_time)"`
	Timezone     int        `orm:"column(timezone)"`
}

// TableName specify the table name for User model. This name is used in the orm RegisterModel
func (u *User) TableName() string {
	return "users"
}

// Save creates a user record in the db
func (u *User) Save() bool {
	_, err := DB.Insert(u)
	return u.processError(err, "save")
}

// Delete deletes the user record from the db
func (u *User) Delete() bool {
	_, err := DB.Delete(u)
	return u.processError(err, "delete")
}

func (u *User) processError(err error, action string) bool {
	if err != nil {
		beego.BeeLogger.Error("failed to "+action+" user `%v` to db. Error: %s", *u, err)
		return false
	}
	return true
}

// UserGetByID fetch the user from users table by id
func UserGetByID(id int) (*User, bool) {
	user := User{Id: id}
	err := DB.Read(&user)
	if err != nil {
		beego.BeeLogger.Error("failed to fetch the user by id: %v error: %s", id, err)
		return nil, false
	}
	return &user, true
}

// UserGetAll fetches all users from the users table
func UserGetAll() ([]User, bool) {
	var users []User
	_, err := DB.QueryTable(UserObj).All(&users)
	if err != nil {
		beego.BeeLogger.Error("failed to fetch all users. Error: %s", err)
		return nil, false
	}
	return users, true
}
