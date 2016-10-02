package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type UserSpheres_20161002_133303 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserSpheres_20161002_133303{}
	m.Created = "20161002_133303"
	migration.Register("UserSpheres_20161002_133303", m)
}

// Run the migrations
func (m *UserSpheres_20161002_133303) Up() {
	m.SQL(`
	CREATE TABLE public.user_spheres
(
    id SERIAL NOT NULL,
    user_id INT NOT NULL,
    sphere_id INT NOT NULL,
    CONSTRAINT user_spheres_users_id_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT user_spheres_spheres_id_fk FOREIGN KEY (sphere_id) REFERENCES spheres (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE UNIQUE INDEX user_spheres_id_uindex ON public.user_spheres (id);
CREATE UNIQUE INDEX user_spheres_user_id_sphere_id_uindex ON public.user_spheres (user_id, sphere_id);
`)
}

// Reverse the migrations
func (m *UserSpheres_20161002_133303) Down() {
	m.SQL(`DROP TABLE IF EXISTS public.user_spheres;`)
}
