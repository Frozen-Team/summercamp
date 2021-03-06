package forms

import (
	"strings"

	"github.com/Frozen-Team/summercamp/models"
	"github.com/astaxie/beego/validation"
)

type UserUpdate struct {
	FormModel
	Field string `json:"field" valid:"Required"`
	Value string `json:"value" valid:"Required"`
}

func (u *UserUpdate) Valid(v *validation.Validation) {
	u.Field = strings.TrimSpace(u.Field)
	u.Value = strings.TrimSpace(u.Value)
	if u.Field == "email" {
		v.Email(u, "value")
	}
}

// Update changes a value of the user field with value from the UserUpdate.Value
func (u *UserUpdate) Update(cu *models.User) (*models.User, bool) {
	switch u.Field {
	case "email":
		cu.Email = u.Value
	case "first_name":
		cu.FirstName = u.Value
	case "last_name":
		cu.LastName = u.Value
	case "country":
		cu.Country = u.Value
	case "city":
		cu.City = u.Value
	case "summary":
		cu.Summary = u.Value
	case "overview":
		cu.Overview = u.Value
	case "timezone":
	// TODO: implement timezone
	default:
		return nil, false
	}

	ok := cu.Save()
	if !ok {
		u.addError("user-save-failed")
		return nil, false
	}
	return cu, true
}
