package forms

import (
	"bitbucket.org/SummerCampDev/summercamp/models"
)

// TODO: consider MaxSize values
type Vacancy struct {
	FormModel
	Name        string `json:"name" valid:"Required; MaxSize(100)"`
	Description string `json:"description" valid:"Required; MaxSize(1000)"`
	Skills      []int  `json:"skills" valid:"Required; MaxSize(3)"`
	Spheres     []int  `json:"spheres" valid:"Required; MaxSize(3)"`
	TeamID      int    `json:"team_id" valid:"Required;"`
}

func (v *Vacancy) Save() (*models.Vacancy, bool) {
	vacancy := &models.Vacancy{
		Name:        v.Name,
		Description: v.Description,
		TeamID:      v.TeamID,
	}

	ok := vacancy.Save()
	if !ok {
		return nil, false
	}
	ok = models.VacancySkills.SaveSkillsForVacancy(vacancy.ID, v.Skills...)
	if !ok {
		return nil, false
	}
	// TODO: vacancy spheres

	return vacancy, ok
}

type VacancyUpdate struct {
	FormModel
	Field string
	Value string
}

func (vu *VacancyUpdate) Update() bool {
	switch vu.Field {
	case "skill":
	case "sphere":
	default:

	}
	return true
}
