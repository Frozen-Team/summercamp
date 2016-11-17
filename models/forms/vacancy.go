package forms

import (
	"github.com/Frozen-Team/summercamp/models"
	"github.com/astaxie/beego/validation"
	"regexp"
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
		Status:      models.VacancyStatusActive,
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
	Field string `json:"field" valid:"Required"`
	Value string `json:"value" valid:"Required"`
}

func (vu *VacancyUpdate) Valid(v *validation.Validation) {
	switch vu.Field {
	case "status":
		r, err := regexp.Compile("(archived|active)")
		if err != nil {
			panic("internal Golang error: failed to compile regexp")
		}

		v.Match(vu.Value, r, "value")
	case "name":
		v.MaxSize(vu.Value, 100, "value")
	case "description":
		v.MaxSize(vu.Value, 1000, "value")
	default:
		v.SetError("field", "Invalid field name")
	}
}

func (vu *VacancyUpdate) Update(id int) bool {
	v := models.Vacancies.NewByID(id)

	switch vu.Field {
	case "status":
		switch models.VacancyStatus(vu.Value) {
		case models.VacancyStatusActive:
			return v.Activate()
		case models.VacancyStatusArchived:
			return v.Archive()
		}
	case "name":
		v.Name = vu.Field
	case "description":
		v.Description = vu.Field
	}

	return v.Save(vu.Field)
}
