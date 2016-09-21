package models

import (
	"fmt"
	"time"

	"bitbucket.org/SummerCampDev/summercamp/models/utils"
	"golang.org/x/crypto/bcrypt"
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
	ID          int        `orm:"column(id)"`
	Type        Speciality `orm:"column(type)"`
	FirstName   string     `orm:"column(first_name)"`
	LastName    string     `orm:"column(last_name)"`
	Email       string     `orm:"column(email)"`
	Password    string     `orm:"column(password)" json:"-"`
	Balance     int        `orm:"column(balance)"`
	Bid         int        `orm:"column(bid)"`
	BraintreeID string     `orm:"column(braintree_id)"`
	Country     string     `orm:"column(country)"`
	City        string     `orm:"column(city)"`
	Timezone    int        `orm:"column(timezone)"`
	CreateTime  time.Time  `orm:"column(create_time);auto_now_add;type(datetime)"`
	UpdateTime  time.Time  `orm:"column(update_time);auto_now;type(datetime)"`
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

// SetPassword generate the encrypted password based on the given string
func (u *User) SetPassword(password string) bool {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.Password = string(encryptedPassword)
	return utils.ProcessError(err, "generate bcrypt password")
}

// CheckPassword checks if given plain password matches hashed password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return utils.ProcessError(err, "check bcrypt password")
}

// Delete deletes the user record from the db
func (u *User) Delete() bool {
	_, err := DB.Delete(u)
	return utils.ProcessError(err, "delete user")
}

type usersAPI struct{}

var Users *usersAPI

// FetchByID fetch the user from users table by id
func (u *usersAPI) FetchByID(id int) (*User, bool) {
	user := User{ID: id}
	err := DB.Read(&user)
	return &user, utils.ProcessError(err, "fetch the user by id")
}

// FetchByEmail fetches the user from users table by email
func (u *usersAPI) FetchByEmail(email string) (*User, bool) {
	var user User
	err := DB.QueryTable(UserObj).Filter("email", email).One(&user)
	return &user, utils.ProcessError(err, "fetch the user by email")
}

// FetchAll fetches all users from the users table
func (u *usersAPI) FetchAll() ([]User, bool) {
	var users []User
	_, err := DB.QueryTable(UserObj).All(&users)
	return users, utils.ProcessError(err, "fetch all users")
}

// FetchAllByType fetch users from users table by speciality
func (u *usersAPI) FetchAllByType(s Speciality) ([]User, bool) {
	if !s.Valid() {
		return nil, utils.ProcessError(fmt.Errorf("Not valid Speciality '%s'", s), "fetch users by type")
	}
	result := []User{}
	_, err := DB.QueryTable(UserObj).Filter("Type__exact", s).All(&result)
	return result, utils.ProcessError(err, "fetch users by type")
}
