package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type UsersRemovePassSalt_20160909_200408 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UsersRemovePassSalt_20160909_200408{}
	m.Created = "20160909_200408"
	migration.Register("UsersRemovePassSalt_20160909_200408", m)
}

// Run the migrations
func (m *UsersRemovePassSalt_20160909_200408) Up() {
	m.SQL("ALTER TABLE public.users DROP password_salt;")

}

// Reverse the migrations
func (m *UsersRemovePassSalt_20160909_200408) Down() {
	m.SQL("ALTER TABLE public.users ADD password_salt TEXT NOT NULL;")

}
