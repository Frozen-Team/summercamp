package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type UserSkills_20161009_203354 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserSkills_20161009_203354{}
	m.Created = "20161009_203354"
	migration.Register("UserSkills_20161009_203354", m)
}

// Run the migrations
func (m *UserSkills_20161009_203354) Up() {
	m.SQL(`CREATE TABLE public.user_skills
(
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INT NOT NULL,
    skill_id INT NOT NULL,
    CONSTRAINT user_skills_users_id_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT user_skills_skills_id_fk FOREIGN KEY (skill_id) REFERENCES skills (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE UNIQUE INDEX user_skills_id_uindex ON public.user_skills (id);
CREATE UNIQUE INDEX user_skills_user_id_skill_id_uindex ON public.user_skills (user_id, skill_id);`)
}

// Reverse the migrations
func (m *UserSkills_20161009_203354) Down() {
	m.SQL(`DROP TABLE IF EXISTS public.user_skills;`)
}
