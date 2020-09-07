package users

import (
	"back-src/controller/control/media"
	"back-src/controller/control/projects/filters"
	"back-src/controller/utils/libs"
	"back-src/model/database"
	"back-src/model/existence"
	"back-src/view/data"
	"errors"
	"time"
)

const (
	ProjectIdSize = 15
)

func EditEmployerProfile(token string, emp existence.Employer, db *database.Database) error {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
		if err := db.EmployerTable.UpdateEmployerProfile(username, emp); err == nil {
			media.AddUpdateProfileEvent(username, false, db)
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func EditEmployerPassword(token string, emp data.ChangePassRequest, db *database.Database) error {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
		return db.EmployerTable.UpdateEmployerPassword(username, emp.OldPass, emp.NewPass)
	} else {
		return err
	}
}

func GetEmployer(token string, db *database.Database) (existence.Employer, error) {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
		if emp, err := db.EmployerTable.GetEmployer(username); err != nil {
			return existence.Employer{}, err
		} else {
			emp.Password = "N/A"
			return emp, nil
		}
	} else {
		return existence.Employer{}, err
	}
}

func GetEmployerProjects(username string, db *database.Database) ([]existence.Project, error) {
	if !db.EmployerTable.DoesEmployerExistWithUsername(username) {
		return nil, errors.New("no user with such username :" + username)
	}
	return db.EmployerTable.GetEmployerProjects(username)
}

func AddProjectToEmployer(token string, project existence.Project, db *database.Database) (e error) {
	e = nil
	if err := checkAddProjectFieldsValidity(project); err == nil {
		if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
			if emp, err := db.EmployerTable.GetEmployer(username); err == nil {
				project.EmployerUsername = username
				project.ProjectStatus = existence.Open
				if project.Id, err = makeNewProjectId(db); err == nil {
					if err := checkProjectSkills(project.Id, project.FieldsWithSkills, db); err == nil {
						db.ProjectTable.AddProject(project)
						emp.ProjectIds = append(emp.ProjectIds, project.Id)
						if err := db.EmployerTable.UpdateEmployerProjects(username, emp); err == nil {
							media.AddAddProjectEvent(username, project.Id, db)
							if e == nil {
								e = nil
							}
						} else {
							e = err
						}
					} else {
						e = err
					}
				} else {
					e = err
				}
			} else {
				e = err
			}
		} else {
			e = err
		}
	} else {
		e = err
	}
	return
}

func checkAddProjectFieldsValidity(project existence.Project) error {
	error := errors.New("project fields not valid")
	if project.MinBudget < project.MaxBudget {
		return error
	}
	return nil
}

func makeNewProjectId(db *database.Database) (id string, e error) {
	id = "p" + libs.GetRandomNumberAsString(ProjectIdSize-1, func(str string) bool {
		if isThere, err := db.ProjectTable.IsThereProjectWithId("p" + str); err != nil {
			e = err
			return false
		} else {
			return isThere
		}
	})
	return id, e
}

func checkProjectSkills(projectId string, fieldsWithSkills map[string][]string, db *database.Database) error {
	for field, skills := range fieldsWithSkills {
		oldSkills, err := db.FieldTable.GetFieldSkills(field)
		if err != nil {
			continue
		}
		for _, skill := range skills {
			if !libs.ContainsString(oldSkills, skill) {
				if err := db.FieldTable.AddSkillToField(field, skill); err != nil {
					return err
				}
			}
			filters.AddSkillToProject(skill, projectId)
		}

	}
	return nil
}

func EditEmployerProject(token string, project existence.Project, db *database.Database) error {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
		if _, err := db.EmployerTable.GetEmployer(username); err == nil {
			if realProject, err := db.ProjectTable.GetProject(project.Id); err == nil {
				if realProject.EmployerUsername == username {
					if realProject.ProjectStatus == existence.Open {
						return db.ProjectTable.EditProject(realProject.Id, project)
					} else {
						return errors.New("project not open")
					}
				} else {
					return errors.New("project access denied")
				}
			} else {
				return err
			}
		} else {
			return err
		}
	} else {
		return err
	}
}

func AssignProjectToFreelancer(token string, freelancer string, projectId string, db *database.Database) error {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
		if realUsername, err := db.ProjectTable.GetEmployerUsernameByProjectId(projectId); err != nil {
			return err
		} else if username != realUsername {
			return errors.New("not valid token for this project")
		} else {
			if requests, err := db.ProjectTable.GetProjectRequests(projectId); err == nil {
				for s := range requests {
					db.FreelancerTable.DeleteFreelancerRequestedProject(s, projectId)
				}
				if err := db.FreelancerTable.AddFreelancerProjectId(freelancer, projectId); err != nil {
					return err
				}
				if err := db.ProjectTable.SetProjectStatus(projectId, existence.OnGoing); err != nil {
					return err
				}
				if err := db.ProjectTable.AddFreelancerToProject(freelancer, projectId); err != nil {
					return err
				}
				if err := db.ProjectTable.DeleteProjectDescriptions(projectId); err != nil {
					return err
				}
				media.AddAssignProjectEvent(username, projectId, db)
				return nil
			} else {
				return err
			}
		}

	} else {
		return err
	}
}
func ExtendProject(token string, projectId string, finishDate time.Time, db *database.Database) error {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
		if realUsername, err := db.ProjectTable.GetEmployerUsernameByProjectId(projectId); err != nil {
			return err
		} else if username != realUsername {
			return errors.New("not valid token for this project")
		} else {
			if firstFinishDate, err := db.ProjectTable.GetProjectFinishDate(projectId); err != nil {
				return err
			} else if finishDate.Before(firstFinishDate) {
				return errors.New("not valid time")
			} else {
				if status, err := db.ProjectTable.GetProjectStatus(projectId); err != nil {
					return err
				} else if status != existence.Open {
					return errors.New("not valid project status")
				} else {
					media.AddExtendProjectEvent(username, projectId, db)
					return db.EmployerTable.ExtendProject(projectId, finishDate)
				}
			}
		}
	} else {
		return err
	}
}

func CloseProject(token string, projectId string, db *database.Database) error {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
		if realUsername, err := db.ProjectTable.GetEmployerUsernameByProjectId(projectId); err != nil {
			return err
		} else if username != realUsername {
			return errors.New("not valid token for this project")
		} else {
			if finishDate, err := db.ProjectTable.GetProjectFinishDate(projectId); err != nil {
				return err
			} else if finishDate.Before(time.Now()) {
				return errors.New("not valid time")
			} else {
				if status, err := db.ProjectTable.GetProjectStatus(projectId); err != nil {
					return err
				} else if status != existence.OnGoing {
					return errors.New("not valid project status")
				} else {
					media.AddCloseProjectEvent(username, projectId, db)
					panic(errors.New("not implemented function error"))
					//TODO(El Tipo, Close Project)
				}
			}
		}
	} else {
		return err
	}
}
