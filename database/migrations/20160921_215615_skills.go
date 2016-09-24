package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Skills_20160921_215615 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Skills_20160921_215615{}
	m.Created = "20160921_215615"
	migration.Register("Skills_20160921_215615", m)
}

// Run the migrations
func (m *Skills_20160921_215615) Up() {
	m.SQL(`CREATE TABLE public.skills
(
    id SERIAL PRIMARY KEY NOT NULL,
    name CITEXT NOT NULL
);
CREATE UNIQUE INDEX skills_id_uindex ON public.skills (id);
CREATE UNIQUE INDEX skills_name_uindex ON public.skills (name);
`)
}

// Reverse the migrations
func (m *Skills_20160921_215615) Down() {
	m.SQL(`DROP TABLE IF EXISTS public.skills;`)
}
