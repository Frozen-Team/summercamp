package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Projects_20160920_212034 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Projects_20160920_212034{}
	m.Created = "20160920_212034"
	migration.Register("Projects_20160920_212034", m)
}

// Run the migrations
func (m *Projects_20160920_212034) Up() {
	m.SQL(`CREATE TABLE public.projects
(
    id SERIAL PRIMARY KEY NOT NULL,
    description TEXT NOT NULL,
    budget INT NOT NULL CHECK (budget > 0),
    client_id INT NOT NULL,
    create_time TIMESTAMP DEFAULT now() NOT NULL,
    update_time TIMESTAMP DEFAULT now() NOT NULL,
    CONSTRAINT projects_users_id_fk FOREIGN KEY (client_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE UNIQUE INDEX projects_id_uindex ON public.projects (id);`)
}

// Reverse the migrations
func (m *Projects_20160920_212034) Down() {
	m.SQL("DROP TABLE IF EXISTS public.projects;")
}
