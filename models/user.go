package models

import (
	"fmt"
	"time"

	"bitbucket.org/SummerCampDev/summercamp/models/utils"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          int        `json:"id" orm:"column(id)"`
	Type        Speciality `json:"type" orm:"column(type)"`
	FirstName   string     `json:"first_name" orm:"column(first_name)"`
	LastName    string     `json:"last_name" orm:"column(last_name)"`
	Overview    string     `json:"overview" orm:"column(overview)"`
	Summary     string     `json:"summary" orm:"column(summary)"`
	Email       string     `json:"email" orm:"column(email)"`
	Password    string     `json:"-" orm:"column(password)"`
	Balance     int        `json:"balance" orm:"column(balance)"`
	Bid         int        `json:"bid" orm:"column(bid)"`
	BraintreeID string     `json:"braintree_id" orm:"column(braintree_id)"`
	Country     string     `json:"country" orm:"column(country)"`
	City        string     `json:"city" orm:"column(city)"`
	Timezone    int        `json:"timezone" orm:"column(timezone)"`
	CreateTime  time.Time  `json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
	UpdateTime  time.Time  `json:"update_time" orm:"column(update_time);auto_now;type(datetime)"`
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

// CanAddSkill checks if the user has less skills than the system allows.
func (u *User) CanAddSkill() (bool, bool) {
	currentSkillsCount, ok := UserSkills.SkillsCountByUser(u.ID)
	if !ok {
		return false, false
	}

	maxSkillsCount, err := beego.AppConfig.Int("UserSkillsMaxCount")
	if err != nil {
		return false, utils.ProcessError(err, " failed to get value for config `UserSkillsMaxCount`")
	}

	return currentSkillsCount < maxSkillsCount, true
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

// Teams returns teams for the current user. If everything is okay, the second
// returned value is true, false - otherwise.
func (u *User) Teams() ([]Team, bool) {
	return TeamMembers.FetchTeamsByMember(u.ID)
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
	err := DB.QueryTable(UserModel).Filter("email", email).One(&user)
	return &user, utils.ProcessError(err, "fetch the user by email")
}

// FetchAll fetches all users from the users table
func (u *usersAPI) FetchAll() ([]User, bool) {
	var users []User
	_, err := DB.QueryTable(UserModel).All(&users)
	return users, utils.ProcessError(err, "fetch all users")
}

// FetchAllByType fetch users from users table by speciality
func (u *usersAPI) FetchAllByType(s Speciality) ([]User, bool) {
	if !s.Valid() {
		return nil, utils.ProcessError(fmt.Errorf("Not valid Speciality '%s'", s), "fetch users by type")
	}
	result := []User{}
	_, err := DB.QueryTable(UserModel).Filter("Type__exact", s).All(&result)
	return result, utils.ProcessError(err, "fetch users by type")
}
