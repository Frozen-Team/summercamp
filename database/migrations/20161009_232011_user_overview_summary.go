package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type UserOverviewSummary_20161009_232011 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserOverviewSummary_20161009_232011{}
	m.Created = "20161009_232011"
	migration.Register("UserOverviewSummary_20161009_232011", m)
}

// Run the migrations
func (m *UserOverviewSummary_20161009_232011) Up() {
	m.SQL(`
	ALTER TABLE public.users ADD summary TEXT NULL;
	ALTER TABLE public.users ADD overview TEXT NULL;
`)
}

// Reverse the migrations
func (m *UserOverviewSummary_20161009_232011) Down() {
	m.SQL(`
	ALTER TABLE public.users DROP COLUMN IF EXISTS summary;
	ALTER TABLE public.users DROP COLUMN IF EXISTS overview;
`)
}
