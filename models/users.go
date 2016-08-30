package models

import "time"

type Speciality int

const (
	SpecManager Speciality = iota
	SpecClient
	SpecExecutor
)

type User struct {
	Id           int        `db:"id"`
	Type         Speciality `db:"type"`
	FirstName    string     `db:"first_name"`
	LastName     string     `db:"last_name"`
	Email        string     `db:"email"`
	Password     string     `db:"password" json:"-"`
	PasswordSalt string     `db:"password_salt" json:"-"`
	Balance      int        `db:"balance"`
	Bid          int        `db:"bid"`
	BraintreeID  string     `db:"braintree_id"`
	Country      string     `db:"country"`
	City         string     `db:"city"`
	CreateTime   time.Time  `db:"create_time"`
	UpdateTime   time.Time  `db:"update_time"`
	Timezone     int        `db:"timezone"`
}
