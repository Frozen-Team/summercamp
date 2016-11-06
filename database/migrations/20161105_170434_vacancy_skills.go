package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type VacancySkills_20161105_170434 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &VacancySkills_20161105_170434{}
	m.Created = "20161105_170434"
	migration.Register("VacancySkills_20161105_170434", m)
}

// Run the migrations
func (m *VacancySkills_20161105_170434) Up() {
	m.SQL(`CREATE TABLE public.vacancy_skills
(
    id SERIAL PRIMARY KEY NOT NULL,
    vacancy_id INT NOT NULL,
    skill_id INT NOT NULL,
    CONSTRAINT vacancy_skills_skills_id_fk FOREIGN KEY (skill_id) REFERENCES skills (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT vacancy_skills_vacancies_id_fk FOREIGN KEY (vacancy_id) REFERENCES vacancies (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE UNIQUE INDEX vacancy_skills_id_uindex ON public.vacancy_skills (id);
CREATE UNIQUE INDEX vacancy_skills_vacancy_id_skill_id_uindex ON public.vacancy_skills (vacancy_id, skill_id);`)
}

// Reverse the migrations
func (m *VacancySkills_20161105_170434) Down() {
	m.SQL(`DROP TABLE IF EXISTS public.vacancy_skills;`)
}
