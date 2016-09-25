package models

// Speciality type defines possible variety of specialities of a user
type Speciality string

const (
	SpecTypeManager  Speciality = "manager"
	SpecTypeClient   Speciality = "client"
	SpecTypeExecutor Speciality = "executor"
)

func (s Speciality) Valid() bool {
	return s == SpecTypeManager || s == SpecTypeClient || s == SpecTypeExecutor
}
