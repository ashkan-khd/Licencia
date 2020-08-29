package control

import (
	"back-src/controller/control/utils/users"
	"back-src/model/database"
	"back-src/model/existence"
	"errors"
	"github.com/gin-gonic/gin"
)

type Control struct {
}

var DB *database.Database

func NewControl() *Control {
	DB = database.NewDb()
	err := DB.Initialize()
	if err != nil {
		panic(err)
	}
	return &Control{}
}

/*
json{
username
email
password
}
QUERIES : "account-type"
*/
func (controller *Control) Register(ctx *gin.Context) error {

	switch accountType := ctx.Query("account-type"); accountType {

	case "employer":
		employer := existence.Employer{}
		if err := ctx.ShouldBindJSON(&employer); err != nil {
			return err
		}
		return users.RegisterEmployer(employer, DB)

	case "freelancer":
		freelancer := existence.Freelancer{}
		if err := ctx.ShouldBindJSON(&freelancer); err != nil {
			return err
		}
		return users.RegisterFreelancer(freelancer, DB)

	default:
		return errors.New("invalid query: " + accountType)
	}

}

func (controller *Control) EditEmployerProfile(ctx *gin.Context) error {
	emp := existence.Employer{}
	if err := ctx.ShouldBindJSON(&emp); err != nil {
		return err
	}
	return users.EditEmployerProfile(emp, DB)
}

func (controller *Control) GetEmployerProfile(ctx *gin.Context) (existence.Employer, error) {
	user := struct {
		username string `json:"username" binding:"required"`
	}{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return existence.Employer{}, err
	}
	return users.GetEmployerProfile(user.username, DB)
}

func (controller *Control) GetEmployerProjects(ctx *gin.Context) ([]existence.Project, error) {
	user := struct {
		username string `json:"username" binding:"required"`
	}{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil, err
	}
	return users.GetEmployerProjects(user.username, DB)
}
