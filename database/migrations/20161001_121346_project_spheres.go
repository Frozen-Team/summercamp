package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ProjectSpheres_20161001_121346 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ProjectSpheres_20161001_121346{}
	m.Created = "20161001_121346"
	migration.Register("ProjectSpheres_20161001_121346", m)
}

// Run the migrations
func (m *ProjectSpheres_20161001_121346) Up() {
	m.SQL(`
	CREATE TABLE public.project_spheres
(
    id SERIAL PRIMARY KEY NOT NULL,
    project_id INT NOT NULL,
    sphere_id INT NOT NULL,
    CONSTRAINT project_spheres_projects_id_fk FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT project_spheres_spheres_id_fk FOREIGN KEY (sphere_id) REFERENCES spheres (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE UNIQUE INDEX project_spheres_id_uindex ON public.project_spheres (id);
CREATE UNIQUE INDEX project_spheres_project_id_sphere_id_uindex ON public.project_spheres (project_id, sphere_id);
`)
}

// Reverse the migrations
func (m *ProjectSpheres_20161001_121346) Down() {
	m.SQL(`DROP TABLE IF EXISTS public.project_spheres;`)
}
