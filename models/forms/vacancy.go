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
		Status:models.VacancyStatusActive,
	}

	ok := vacancy.Save()
	if !ok {
		return nil, false
	}
	ok = models.VacancySkills.SaveSkillsForVacancy(vacancy.ID, v.Skills...)
	if !ok {
		return nil, false
	}
	ok = models.VacancySpheres.SaveSpheresForVacancy(vacancy.ID, v.Spheres...)
	if !ok {
		return nil, false
	}

	return vacancy, ok
}

type VacancyUpdate struct {
	FormModel
	Field string
	Value string
}

func (vu *VacancyUpdate) Update(id int) bool {
	switch vu.Field {
	case "status":
		switch vu.Value {
		case models.VacancyStatusActive:
			return models.Vacancies.Activate(id)
		case models.VacancyStatusArchived:
			return models.Vacancies.Archive(id)
		}
	case "skill":
	case "sphere":
	default:

	}
	return true
}
