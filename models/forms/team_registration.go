package forms

import (
	"bitbucket.org/SummerCampDev/summercamp/models"
)

type TeamRegistration struct {
	FormModel
	Name        string	`json:"name" valid:"Required"`
	Description string	`json:"description" valid:"Required"`
}

// Register validates the input data and if everything is OK, initialize the models.Team struct with
// the data from Registration struct and save the record to the db.
func (tr *TeamRegistration) Register(current *models.User) (*models.Team, bool) {
	if ok := tr.validate(tr); !ok {
		return nil, false
	}
	team := &models.Team{
		Name:     	tr.Name,
		Description:	tr.Description,
	}
	ok := team.Save()
	if !ok {
		tr.addError("team-save-failed")
		return nil, false
	}
	ok = team.AddMember(current, models.LevelAdmin)
	if !ok {
		return nil, team.Delete()
	}
	return team, true
}
