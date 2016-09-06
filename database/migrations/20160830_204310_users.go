package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Users_20160830_204310 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20160830_204310{}
	m.Created = "20160830_204310"
	migration.Register("Users_20160830_204310", m)
}

// Run the migrations
func (m *Users_20160830_204310) Up() {
	m.SQL(`
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'speciality') THEN
			CREATE TYPE speciality AS ENUM ('manager', 'client', 'executor');
		END IF;
	END$$;
	CREATE EXTENSION IF NOT EXISTS citext;
	CREATE TABLE public.users
	(
		id SERIAL PRIMARY KEY NOT NULL,
		type SPECIALITY NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email CITEXT NOT NULL,
		password TEXT NOT NULL,
		password_salt TEXT NOT NULL,
		balance INT NOT NULL,
		bid INT NOT NULL,
		braintree_id TEXT NOT NULL,
		country TEXT NOT NULL,
		city TEXT NOT NULL,
		timezone INT NOT NULL,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_DATE,
		update_time TIMESTAMP NOT NULL
	);
	CREATE UNIQUE INDEX users_id_uindex ON public.users (id);
	CREATE UNIQUE INDEX users_email_uindex ON public.users (email);`)
}

// Reverse the migrations
func (m *Users_20160830_204310) Down() {
	m.SQL(`DROP TABLE public.users;
	DROP TYPE IF EXISTS SPECIALITY;`)
}
