package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type VacancySpheres_20161107_000336 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &VacancySpheres_20161107_000336{}
	m.Created = "20161107_000336"
	migration.Register("VacancySpheres_20161107_000336", m)
}

// Run the migrations
func (m *VacancySpheres_20161107_000336) Up() {
	m.SQL(`
	CREATE TABLE public.vacancy_spheres
(
    id SERIAL PRIMARY KEY NOT NULL,
    vacancy_id INT NOT NULL,
    sphere_id INT NOT NULL,
    CONSTRAINT vacancy_spheres_vacancies_id_fk FOREIGN KEY (vacancy_id) REFERENCES vacancies (id) ON DELETE CASCADE ON UPDATE CASCADE DEFERRABLE,
    CONSTRAINT vacancy_spheres_spheres_id_fk FOREIGN KEY (sphere_id) REFERENCES spheres (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE UNIQUE INDEX vacancy_spheres_id_uindex ON public.vacancy_spheres (id);
CREATE UNIQUE INDEX "vacancy_spheres_vacancy_id_sphere_Id_uindex" ON public.vacancy_spheres (vacancy_id, sphere_id);`)
}

// Reverse the migrations
func (m *VacancySpheres_20161107_000336) Down() {
	m.SQL(`DROP TABLE IF EXISTS public.vacancy_spheres;`)
}
