package users

import (
	"back-src/model/existence"
	"back-src/view"
	"back-src/view/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func RespondEmployerEdit(context *gin.Context, token string, err error) {
	if err == nil {
		context.Header("Token", token)
		context.JSON(http.StatusOK, responses.Response{Message: "Successful"})
	} else {
		if !view.RespondTokenErrors(context, err) {
			context.Header("Token", token)
			if !view.RespondDataValidationErrors(context, err) {
				var status int = http.StatusInternalServerError
				context.JSON(status, responses.Response{Message: err.Error()})
			}
		}
	}
}

func RespondEmployerGetProfile(context *gin.Context, token string, emp existence.Employer, err error) {
	if err == nil {
		context.Header("Token", token)
		context.JSON(http.StatusOK, emp)
	} else {
		if !view.RespondTokenErrors(context, err) {
			context.Header("Token", token)
			//TODO : add switch cases if there are other types of error
			var status int = http.StatusInternalServerError
			context.JSON(status, responses.Response{Message: err.Error()})
		}
	}
}

func RespondEmployerAddProject(context *gin.Context, token string, err error) {
	if err == nil {
		context.Header("Token", token)
		context.JSON(http.StatusOK, responses.Response{Message: "Successful"})
	} else {
		if !view.RespondTokenErrors(context, err) {
			context.Header("Token", token)
			if !view.RespondDataValidationErrors(context, err) {
				var status int
				switch {
				case strings.Contains(err.Error(), "project fields not valid"):
					status = http.StatusBadRequest
				default:
					status = http.StatusInternalServerError
				}
				context.JSON(status, responses.Response{Message: err.Error()})
			}
		}
	}
}

func RespondEmployerEditProject(context *gin.Context, token string, err error) {
	if err == nil {
		context.Header("Token", token)
		context.JSON(http.StatusOK, responses.Response{Message: "Successful"})
	} else {
		if !view.RespondTokenErrors(context, err) {
			context.Header("Token", token)
			if !view.RespondDataValidationErrors(context, err) {
				var status int
				switch err.Error() {
				case "project access denied":
					status = http.StatusForbidden
				case "project not open":
					status = http.StatusBadRequest
				default:
					status = http.StatusInternalServerError
				}
				context.JSON(status, responses.Response{Message: err.Error()})
			}
		}
	}
}

func RespondEmployerAssignProject(context *gin.Context, token string, err error) {
	if err == nil {
		context.Header("Token", token)
		context.JSON(http.StatusOK, responses.Response{Message: "Successful"})
	} else {
		if !view.RespondTokenErrors(context, err) {
			context.Header("Token", token)
			if !view.RespondDataValidationErrors(context, err) {
				var status int = http.StatusInternalServerError
				context.JSON(status, responses.Response{Message: err.Error()})
			}
		}
	}
}

func RespondEmployerGetProjects(context *gin.Context, projects []existence.Project, err error) {
	if err == nil {
		context.JSON(http.StatusOK, struct {
			projects []existence.Project `json:"projects"`
		}{projects: projects})
	} else {
		var status int
		switch {
		case strings.Contains(err.Error(), "no user with such username :"):
			status = http.StatusBadRequest
		default:
			status = http.StatusInternalServerError
		}
		context.JSON(status, responses.Response{Message: err.Error()})
	}
}
