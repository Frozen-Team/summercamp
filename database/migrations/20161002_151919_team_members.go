package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type TeamMember_20161002_151919 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &TeamMember_20161002_151919{}
	m.Created = "20161002_151919"
	migration.Register("TeamMember_20161002_151919", m)
}

// Run the migrations
func (m *TeamMember_20161002_151919) Up() {
	m.SQL(`ALTER TABLE public.team_members
    ADD COLUMN access integer;`)
}

// Reverse the migrations
func (m *TeamMember_20161002_151919) Down() {
	m.SQL(`ALTER TABLE public.team_members
    DROP COLUMN access integer;`)
}
