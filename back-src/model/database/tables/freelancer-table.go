package tables

import (
	"back-src/controller/utils/libs"
	"back-src/model/existence"
	"errors"
	"github.com/go-pg/pg"
)

type FreelancerTable struct {
	*pg.DB
}

func (table *FreelancerTable) DoesFreelancerExistWithUsername(username string) bool {
	resultSet := &[]existence.Freelancer{}
	_ = table.Model(resultSet).Where("username = ?", username).Select()
	return len(*resultSet) != 0
}

func (table *FreelancerTable) DoesFreelancerExistWithEmail(email string) bool {
	resultSet := &[]existence.Freelancer{}
	_ = table.Model(resultSet).Where("email = ?", email).Select()
	return len(*resultSet) != 0
}

func (table *FreelancerTable) InsertFreelancer(frl existence.Freelancer) error {
	_, err := table.Model(&frl).Insert()
	return err
}

func (table *FreelancerTable) AddFreelancerSkills(username string, fieldId string, skills []string) error {
	var frl existence.Freelancer
	if err := table.Model(&frl). /*.Column("chosen_field_with_skills")*/ Where("username = ?", username).Select(); err != nil {
		return err
	}
	if frl.ChosenFieldWithSkills == nil {
		frl.ChosenFieldWithSkills = map[string][]string{}
	}
	frl.ChosenFieldWithSkills[fieldId] = skills
	_, err := table.Model(&frl).Column("chosen_field_with_skills").Where("username = ?", username).Update()
	return err
}

func (table *FreelancerTable) UpdateFreelancerProfile(username string, frl existence.Freelancer) error {
	if _, err := table.Model(&frl).Column("shown_name", "description", "first_name", "last_name", "phone_number", "address").Where("username = ?", username).Update(); err != nil {
		return err
	}
	return nil
}

func (table *FreelancerTable) UpdateFreelancerPassword(username string, oldPass string, newPass string) error {
	frl, _ := table.GetFreelancer(username)
	if frl.Password != oldPass {
		return errors.New("password mismatch")
	}

	frl.Password = newPass
	if _, err := table.Model(&frl).Column("password").Where("username = ?", username).Update(); err != nil {
		return err
	}
	return nil
}

func (table *FreelancerTable) UpdateFreelancerLinks(username string, frl existence.Freelancer) error {
	if _, err := table.Model(&frl).Column("website", "github", "github_repos").Where("username = ?", username).Update(); err != nil {
		return err
	}
	return nil
}

func (table *FreelancerTable) GetFreelancerByUsername(username string) (existence.Freelancer, error) {
	var frl existence.Freelancer
	if err := table.Model(&frl).Where("username = ?", username).Select(); err != nil {
		return frl, err
	}
	return frl, nil
}

func (table *FreelancerTable) GetFreelancerPasswordByUsername(username string) (string, error) {
	freelancer := existence.Freelancer{}
	err := table.Model(&freelancer).Where("username = ?", username).Column("password").Select()
	return freelancer.Password, err
}

func (table *FreelancerTable) GetFreelancerUsernameByEmail(email string) (string, error) {
	freelancer := existence.Freelancer{}
	err := table.Model(&freelancer).Where("email = ?", email).Column("username").Select()
	return freelancer.Username, err
}

func (table *FreelancerTable) GetFreelancer(username string) (existence.Freelancer, error) {
	frl := new(existence.Freelancer)
	err := table.Model(frl).Where("username = ?", username).Select()
	return *frl, err
}

func (table *FreelancerTable) DeleteFreelancerRequestedProject(username string, projectId string) error {
	frl, err := table.GetFreelancer(username)
	if err != nil {
		return err
	}
	var index int
	for i, id := range frl.RequestedProjectIds {
		if id == projectId {
			index = i
			break
		}
	}
	frl.RequestedProjectIds = libs.RemoveStringElement(frl.RequestedProjectIds, index)
	if _, err = table.Model(&frl).Column("requested_project_ids").Where("username = ?", username).Update(); err != nil {
		return err
	}
	return nil
}

func (table *FreelancerTable) AddFreelancerProjectId(username string, projectId string) error {
	frl, err := table.GetFreelancer(username)
	if err != nil {
		return err
	}
	found := false
	for _, id := range frl.ProjectIds {
		if id == projectId {
			found = true
		}
	}
	if found {
		return nil
	}
	frl.ProjectIds = append(frl.ProjectIds, projectId)
	if _, err = table.Model(&frl).Column("project_ids").Where("username = ?", username).Update(); err != nil {
		return err
	}
	return nil
}