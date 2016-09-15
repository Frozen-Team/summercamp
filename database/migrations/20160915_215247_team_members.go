package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type TeamMembers_20160915_215247 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &TeamMembers_20160915_215247{}
	m.Created = "20160915_215247"
	migration.Register("TeamMembers_20160915_215247", m)
}

// Run the migrations
func (m *TeamMembers_20160915_215247) Up() {
	m.SQL(`CREATE TABLE public.team_members
	(
	    id SERIAL PRIMARY KEY NOT NULL,
	    team_id INT NOT NULL,
	    user_id INT NOT NULL,
	    create_time TIMESTAMP DEFAULT now() NOT NULL,
	    CONSTRAINT team_members_teams_id_fk FOREIGN KEY (team_id) REFERENCES teams (id) ON DELETE CASCADE ON UPDATE CASCADE,
	    CONSTRAINT team_members_users_id_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
	);
	CREATE UNIQUE INDEX team_members_id_uindex ON public.team_members (id);
	CREATE UNIQUE INDEX team_members_user_team_ids_uindex ON public.team_members (team_id, user_id); `)
}

// Reverse the migrations
func (m *TeamMembers_20160915_215247) Down() {
	m.SQL(`DROP TABLE IF EXISTS public.team_members;`)

}
