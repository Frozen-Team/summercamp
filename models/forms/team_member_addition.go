package forms

import "bitbucket.org/SummerCampDev/summercamp/models"

type TeamMemberAddition struct {
	FormModel
	UserID int                `json:"user_id" valid:"Required"`
	Access models.AccessLevel `json:"access" valid:"Required"`
}

// AddMember validates the input data and if everything is OK, adds new member to the team
func (tma *TeamMemberAddition) AddMember(team *models.Team) (*models.TeamMember, bool) {
	return team.AddMember(tma.UserID, tma.Access)
}
