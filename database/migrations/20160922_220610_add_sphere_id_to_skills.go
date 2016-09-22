package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddSphereIdToSkills_20160922_220610 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddSphereIdToSkills_20160922_220610{}
	m.Created = "20160922_220610"
	migration.Register("AddSphereIdToSkills_20160922_220610", m)
}

// Run the migrations
func (m *AddSphereIdToSkills_20160922_220610) Up() {
	m.SQL(`ALTER TABLE public.skills ADD sphere_id INT NOT NULL;
	CREATE UNIQUE INDEX skills_name_sphere_id_uindex ON public.skills (name, sphere_id);
	ALTER TABLE public.skills
	ADD CONSTRAINT skills_spheres_id_fk
	FOREIGN KEY (sphere_id) REFERENCES spheres (id) ON DELETE CASCADE ON UPDATE CASCADE;`)
}

// Reverse the migrations
func (m *AddSphereIdToSkills_20160922_220610) Down() {
	m.SQL(`
		ALTER TABLE public.skills DROP COLUMN IF EXISTS sphere_id;
		DROP INDEX IF EXISTS skills_name_sphere_id_uindex;
		ALTER TABLE public.skills DROP CONSTRAINT IF EXISTS skills_spheres_id_fk;
	`)
}
