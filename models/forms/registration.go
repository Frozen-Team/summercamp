package forms

type UserReg struct {
	Email           string `json:"email" valid:"Required; Email"`
	Type            string `json:"type" valid:"Required"`
	FirstName       string `json:"first_name" valid:"Required"`
	LastName        string `json:"last_name" valid:"Required"`
	Password        string `json:"password" valid:"Required"` //TODO: think about password restrictions
	PasswordConfirm string `json:"password_confirm" valid:"Required"`
	Country         string `json:"country" valid:"Required"`
	City            string `json:"city" valid:"Required"`
}
