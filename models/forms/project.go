package forms

import "bitbucket.org/SummerCampDev/summercamp/models"

type Project struct {
	FormModel
	ClientID     int            `json:"client_id" valid:"Required"`
	Description  string         `json:"description" valid:"Required"`
	Budget       int            `json:"budget" valid:"Required"`
	SphereSkills []SphereSkills `json:"sphere_skills" valid:"Required"`
}

func (p *Project) Save() (*models.Project, bool) {
	project := &models.Project{
		ClientID:    p.ClientID,
		Description: p.Description,
		Budget:      p.Budget,
	}

	if !project.Save() {
		return nil, false
	}

	for _, sphereSkills := range p.SphereSkills {
		projectSphere := models.ProjectSphere{
			ProjectID: project.ID,
			SphereID:  sphereSkills.Sphere,
		}

		if !projectSphere.Save() {
			return nil, false
		}

		for _, skill := range sphereSkills.Skills {
			projectSkill := models.ProjectSkill{
				ProjectID: project.ID,
				SkillID:   skill,
			}
			if !projectSkill.Save() {
				return nil, false
			}
		}
	}

	return project, true
}
