package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ProjectSkills_20160921_231744 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ProjectSkills_20160921_231744{}
	m.Created = "20160921_231744"
	migration.Register("ProjectSkills_20160921_231744", m)
}

// Run the migrations
func (m *ProjectSkills_20160921_231744) Up() {
	m.SQL(`
CREATE TABLE public.project_skills
(
    id SERIAL PRIMARY KEY NOT NULL,
    project_id INT NOT NULL,
    skill_id INT NOT NULL,
    CONSTRAINT project_skills_projects_id_fk FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT project_skills_skills_id_fk FOREIGN KEY (skill_id) REFERENCES skills (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE UNIQUE INDEX project_skills_id_uindex ON public.project_skills (id);
CREATE UNIQUE INDEX project_skills_project_id_skill_id_uindex ON public.project_skills (project_id, skill_id);
`)
}

// Reverse the migrations
func (m *ProjectSkills_20160921_231744) Down() {
	m.SQL(`DROP TABLE IF EXISTS public.project_skills;`)
}
