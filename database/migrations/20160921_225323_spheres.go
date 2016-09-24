package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Spheres_20160921_225323 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Spheres_20160921_225323{}
	m.Created = "20160921_225323"
	migration.Register("Spheres_20160921_225323", m)
}

// Run the migrations
func (m *Spheres_20160921_225323) Up() {
	m.SQL(`CREATE TABLE public.spheres
(
    id SERIAL PRIMARY KEY NOT NULL,
    name CITEXT NOT NULL
);
CREATE UNIQUE INDEX spheres_id_uindex ON public.spheres (id);
CREATE UNIQUE INDEX spheres_name_uindex ON public.spheres (name);
`)

}

// Reverse the migrations
func (m *Spheres_20160921_225323) Down() {
	m.SQL(`DROP TABLE IF EXISTS public.spheres;`)
}
