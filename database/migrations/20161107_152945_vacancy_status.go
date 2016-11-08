package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type VacancyStatus_20161107_152945 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &VacancyStatus_20161107_152945{}
	m.Created = "20161107_152945"
	migration.Register("VacancyStatus_20161107_152945", m)
}

// Run the migrations
func (m *VacancyStatus_20161107_152945) Up() {
	m.SQL(`CREATE TYPE vacancy_status AS ENUM ('archived', 'active');
	ALTER TABLE public.vacancies ADD status VACANCY_STATUS DEFAULT 'active' NOT NULL;`)
}

// Reverse the migrations
func (m *VacancyStatus_20161107_152945) Down() {
	m.SQL(`DROP TYPE IF EXISTS vacancy_status;
	ALTER TABLE IF EXISTS vacancies DROP COLUMN IF EXISTS status;`)
}
