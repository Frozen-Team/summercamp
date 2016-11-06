package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type SpecialityNew_20161023_112231 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &SpecialityNew_20161023_112231{}
	m.Created = "20161023_112231"
	migration.Register("SpecialityNew_20161023_112231", m)
}

// Run the migrations
func (m *SpecialityNew_20161023_112231) Up() {
	m.SQL(`DROP TYPE IF EXISTS speciality_new;

CREATE TYPE speciality_new AS ENUM ('client', 'executor');

DELETE FROM users u WHERE u.type = 'manager';

ALTER TABLE users
  ALTER COLUMN "type" TYPE speciality_new
  USING ("type"::text::speciality_new);

DROP TYPE speciality;

ALTER TYPE speciality_new RENAME TO speciality;`)
}

// Reverse the migrations
func (m *SpecialityNew_20161023_112231) Down() {
	m.SQL(`DROP TYPE IF EXISTS speciality_new;

CREATE TYPE speciality_new AS ENUM ('client', 'executor', 'manager');

ALTER TABLE users
  ALTER COLUMN "type" TYPE speciality_new
  USING ("type"::text::speciality_new);

DROP TYPE speciality;

ALTER TYPE speciality_new RENAME TO speciality;`)
}
