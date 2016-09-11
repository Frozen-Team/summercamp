package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Teams_20160911_201938 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Teams_20160911_201938{}
	m.Created = "20160911_201938"
	migration.Register("Teams_20160911_201938", m)
}

// Run the migrations
func (m *Teams_20160911_201938) Up() {
	m.SQL(`CREATE TABLE public.teams
	(
	    id SERIAL PRIMARY KEY NOT NULL,
	    name TEXT NOT NULL,
	    description TEXT,
	    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_DATE
	);
	CREATE UNIQUE INDEX teams_id_uindex ON public.teams (id);
	CREATE UNIQUE INDEX teams_name_uindex ON public.teams (name);`)

}

// Reverse the migrations
func (m *Teams_20160911_201938) Down() {
	m.SQL("DROP TABLE IF EXISTS public.teams;")
}
