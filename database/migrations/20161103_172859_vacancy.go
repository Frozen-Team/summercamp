package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Vacancy_20161103_172859 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Vacancy_20161103_172859{}
	m.Created = "20161103_172859"
	migration.Register("Vacancy_20161103_172859", m)
}

// Run the migrations
func (m *Vacancy_20161103_172859) Up() {
	m.SQL(`CREATE TABLE public.vacancy
(
    id SERIAL PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    team_id INT NOT NULL,
    published TIMESTAMP DEFAULT now(),
    CONSTRAINT vacancy_teams_id_fk FOREIGN KEY (team_id) REFERENCES teams (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE UNIQUE INDEX vacancy_id_uindex ON public.vacancy (id);
CREATE UNIQUE INDEX vacancy_name_team_id_uindex ON public.vacancy (name, team_id);`)
}

// Reverse the migrations
func (m *Vacancy_20161103_172859) Down() {
	m.SQL(`DROP TABLE IF EXISTS public.vacancy;`)
}
