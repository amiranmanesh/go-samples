package project

import (
	"awesome_webkits/database/models"
	auth_user "awesome_webkits/http/auth"
	"awesome_webkits/http/requests/v1"
	"awesome_webkits/http/response"
	"awesome_webkits/utils/validation"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
)

type projectControllerInterface interface {
	Save(ctx *gin.Context)
	NameEdit(ctx *gin.Context)
}

type projectController struct{}

var ProjectController projectControllerInterface = &projectController{}

func (projectController) Save(ctx *gin.Context) {

	var request requests.NewProjectRequest
	{
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validatorErrors := validation.GetValidationErrors(err)
		response.Json.Error(ctx, http.StatusBadRequest, "ShouldBindJSON Error", validatorErrors)
		return // exit on first error
	}

	user := auth_user.GetAuthUser()

	var project models.Project
	if err := copier.Copy(&project, &request); err != nil {
		response.Json.Error(ctx, http.StatusBadRequest, "Invalid json body to cast", err)
		return
	}
	project.User = user
	project.UserID = user.ID

	pToken, err := project.Save()
	if err != nil {
		response.Json.Object(ctx, *err)
		return
	}

	response.Json.Success(
		ctx,
		gin.H{
			"project_token": pToken,
		},
	)

}

func (projectController) NameEdit(ctx *gin.Context) {

	var request requests.EditProjectRequest
	{
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validatorErrors := validation.GetValidationErrors(err)
		response.Json.Error(ctx, http.StatusBadRequest, "ShouldBindJSON Error", validatorErrors)
		return
	}

	user := auth_user.GetAuthUser()

	var project models.Project
	projectTemp, err := project.GetAuthProject(ctx.GetHeader("Project-Token"))
	if err != nil {
		response.Json.Object(ctx, *err)
		return
	}
	project = *projectTemp

	if project.UserID != user.ID {
		response.Json.Object(ctx, *response.NewUnauthorizedError("You don't have the permission", nil))
		return
	}

	if err := project.UpdateName(request.Name); err != nil {
		response.Json.Object(ctx, *err)
		return
	}

	response.Json.Success(
		ctx,
		nil,
	)

}
