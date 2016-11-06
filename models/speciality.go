package models

// Speciality type defines possible variety of specialities of a user
type Speciality string

const (
	SpecTypeClient   Speciality = "client"
	SpecTypeExecutor Speciality = "executor"
)

func (s Speciality) Valid() bool {
	return s == SpecTypeClient || s == SpecTypeExecutor
}
