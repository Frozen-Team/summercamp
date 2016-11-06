package controllers

import (
	"bitbucket.org/SummerCampDev/summercamp/models"
	"bitbucket.org/SummerCampDev/summercamp/models/forms"
)

// Operations about Users
type Users struct {
	ApplicationController
}

func (u *Users) Prepare() {
	u.SkipAuthorizationActions("Register", "LogIn")
	u.ApplicationController.Prepare()
}

// Register reads the data from the request body into forms.UserReg struct and attempts to save a user to db
// @Title Register
// @Description User registration. If successful, the user is also authorized
// @Param body body string true "Registration info"
// @Success 200 {object} models.User
// @Failure 200 Nil object and error tag
// @router / [post]
func (u *Users) Register() {
	regForm := new(forms.UserRegistration)

	if ok := u.unmarshalJSON(regForm); !ok {
		u.serveAJAXInternalServerError()
		return
	}

	if ok := forms.Validate(regForm); !ok {
		u.serveAJAXBadRequest(regForm.Errors...)
		return
	}

	if user, ok := regForm.Register(); ok {
		u.authorizeUser(user)
		u.serveAJAXSuccess(user)
	} else {
		u.serveAJAXInternalServerError()
	}
}

// LogIn reads the data from the request body into forms.UserLogin struct, attempts to query a user from the db
// by email and checks password. In case of success the user is authorized
// @Title LogIn
// @Description LogIn a user to the system
// @Param body body string true "Json body message with user credentials"
// @Success 200 {object} models.User
// @Failure 500 Internal server error
// @Failure 400 bad-request
// @router /login [post]
func (u *Users) LogIn() {
	loginForm := new(forms.UserLogin)

	if ok := u.unmarshalJSON(loginForm); !ok {
		u.serveAJAXInternalServerError()
		return
	}

	if ok := forms.Validate(loginForm); !ok {
		u.serveAJAXBadRequest(loginForm.Errors...)
		return
	}

	if user, ok := loginForm.LogIn(); ok {
		u.authorizeUser(user)
		u.serveAJAXSuccess(user)
	} else {
		u.serveAJAXInternalServerError()
	}
}

// LogOut deauthorizes logged-in User otherwise responses "bad-request"
// @Title LogOut
// @Description LogOut a user from the system
// @Success 200 OK
// @Failure 401 unauthorized
// @router /logout [post]
func (u *Users) LogOut() {
	u.deauthorizeUser()
	u.serveAJAXSuccess(nil)
}

// @Title Current
// @Description Get info about the currently logged-in user
// @Success 200 {object} models.User
// @Failure 401 unauthorized
// @router /current [get]
func (u *Users) Current() {
	u.serveAJAXSuccess(u.currentUser)
}

// @Title UpdateField
// @Description Updates user field
// @Param body body string true "A body that should contain a field name and new value"
// @Success 200 {object} models.User
// @Failure 400 bad-data
// @Failure 401 unauthorized
// @Failure 500 internal-error
// @router / [put]
func (u *Users) UpdateField() {
	form := new(forms.UserUpdate)

	if ok := u.unmarshalJSON(form); !ok {
		u.serveAJAXInternalServerError()
		return
	}

	if ok := forms.Validate(form); !ok {
		u.serveAJAXBadRequest(form.Errors...)
		return
	}

	if _, ok := form.Update(u.currentUser); !ok {
		u.serveAJAXInternalServerError()
		return
	}
	u.serveAJAXSuccess(u.currentUser)
}

// @Title UpdatePassword
// @Description Updates password of the user
// @Param body body string true "A body that should contain the current password, password and password_confirm fields"
// @Success 200 {object} models.User
// @Failure 401 unauthorized
// @Failure 400 bad-data
// @Failure 500 internal-error
// @router /update_password [post]
func (u *Users) UpdatePassword() {
	form := new(forms.UserPasswordUpdate)

	if ok := u.unmarshalJSON(form); !ok {
		u.serveAJAXInternalServerError()
		return
	}

	if ok := forms.Validate(form); !ok {
		u.serveAJAXBadRequest(form.Errors...)
		return
	}

	if _, ok := form.UpdatePassword(u.currentUser); !ok {
		u.serveAJAXInternalServerError()
		return
	}
	u.serveAJAXSuccess(u.currentUser)
}

// @Title GetUser
// @Description Get info about a user by its id
// @Param id path int true "An id of a user you want to get"
// @Success 200 {object} models.User
// @Failure 400 invalid-id or no-such-user
// @Failure 401 unauthorized
// @router /:id [get]
func (u *Users) GetUser() {
	// TODO: Check if the requested user can be seen (publicly or privately)
	id := u.getID()

	if user, ok := models.Users.FetchByID(id); ok {
		u.serveAJAXSuccess(user)
	} else {
		u.serveAJAXBadRequest("no-such-user")
	}
}

// @Title GetSkills
// @Description get skills for the user with id passed in the url
// @Param id path int true "An id of a user you want to get skills for"
// @Success 200 {array} models.Skill
// @Failure 400 bad-request
// @Failure 401 unauthorized
// @Failure 500 internal-error
// @router /:id/skills [get]
func (u *Users) GetSkills() {
	userID := u.currentUser.ID

	if skills, ok := models.UserSkills.FetchSkillsByUser(userID); ok {
		u.serveAJAXSuccess(skills)
	} else {
		u.serveAJAXInternalServerError()
	}
}

// @Title AddSkill
// @Description add skill for the currently logged in user
// @Param skill_id body int true "skill id of the skill, max 10 skills"
// @Success 200 {object} models.UserSKill
// @Failure 400 bad-request and validation error
// @Failure 401 unauthorized
// @Failure 500 internal-error
// @router /skills [post]
func (u *Users) AddSkill() {
	form := new(forms.UserSkill)

	if ok := u.unmarshalJSON(form); !ok {
		u.serveAJAXInternalServerError()
		return
	}

	form.UserID = u.currentUser.ID

	if ok := forms.Validate(form); !ok {
		u.serveAJAXBadRequest(form.Errors...)
		return
	}

	canAdd, ok := u.currentUser.CanAddSkill()
	if !ok {
		u.serveAJAXInternalServerError()
		return
	}
	if !canAdd {
		u.serveAJAXBadRequest("max-skills-count")
		return
	}

	if userSkill, ok := form.Save(); ok {
		u.serveAJAXSuccess(userSkill)
	} else {
		u.serveAJAXInternalServerError()
	}
}

// @Title RemoveSkill
// @Description remove skill for the user
// @Param id path int true "id of the userSkill to be removed"
// @Success 200 OK
// @Failure 400 invalid-id
// @Failure 401 unauthorized
// @Failure 500 internal-error
// @router /skills/:id [delete]
func (u *Users) RemoveSkill() {
	userSkillID := u.getID()

	userSkill := models.UserSkill{ID: userSkillID}

	if ok := userSkill.Delete(); ok {
		u.serveAJAXSuccess(nil)
	} else {
		u.serveAJAXInternalServerError()
	}
}

// @Title AddSphere
// @Description Create a new user sphere record
// @Param sphere_id body int true "sphere id to be added for the current user"
// @Success 200 {object} models.UserSphere
// @Failure 400 bad-request + validation errors
// @Failure 401 unauthorized
// @Failure 500 internal-error
// @router /spheres [post]
func (u *Users) AddSphere() {
	form := new(forms.UserSphere)

	if ok := u.unmarshalJSON(form); !ok {
		u.serveAJAXInternalServerError()
		return
	}

	form.UserID = u.currentUser.ID

	if ok := forms.Validate(form); !ok {
		u.serveAJAXBadRequest(form.Errors...)
		return
	}

	if userSphere, ok := form.Save(); ok {
		u.serveAJAXSuccess(userSphere)
	} else {
		u.serveAJAXInternalServerError()
	}
}

// @Title RemoveSphere
// @Description remove sphere for the user
// @Param id path int true "id of the userSphere to be removed"
// @Success 200 OK
// @Failure 400 invalid-id
// @Failure 401 unauthorized
// @Failure 500 internal-error
// @router /spheres/:id [delete]
func (u *Users) RemoveSphere() {
	userSphereID := u.getID()

	userSphere := models.UserSphere{ID: userSphereID}

	if ok := userSphere.Delete(); ok {
		u.serveAJAXSuccess(nil)
	} else {
		u.serveAJAXInternalServerError()
	}
}
