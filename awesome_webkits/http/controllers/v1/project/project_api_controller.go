package project

import (
	"awesome_webkits/cache"
	"awesome_webkits/database/models"
	auth_user "awesome_webkits/http/auth"
	"awesome_webkits/http/requests/v1"
	"awesome_webkits/http/response"
	"awesome_webkits/utils/validation"
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type projectApiControllerInterface interface {
	Save(ctx *gin.Context)
	Show(ctx *gin.Context)
	Edit(ctx *gin.Context)
}

type projectApiController struct{}

var ProjectApiController projectApiControllerInterface = &projectApiController{}

func (projectApiController) Save(ctx *gin.Context) {

	var request requests.ProjectApiSaveRequest
	{
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validatorErrors := validation.GetValidationErrors(err)
		response.Json.Error(ctx, http.StatusBadRequest, "ShouldBindJSON Error", validatorErrors)
		return
	}

	//Authenticate user
	user := auth_user.GetAuthUser()

	var project models.Project
	projectTemp, err := project.GetAuthProject(ctx.GetHeader("Project-Token"))
	if err != nil {
		response.Json.Object(ctx, *err)
		return
	}
	project = *projectTemp

	//Check if its user has access to this project
	//TODO : make it middleware
	if project.UserID != user.ID {
		response.Json.Object(ctx, *response.NewUnauthorizedError("You don't have the permission", nil))
		return
	}

	//we are implementing type= static for Now
	if request.Type != "static" {
		response.Json.Simple(ctx, http.StatusNotImplemented)
		return
	}

	var projectApi models.ProjectApi
	if err := copier.Copy(&projectApi, &request); err != nil {
		response.Json.Error(ctx, http.StatusBadRequest, "Invalid json body to cast", err)
		return
	}
	projectApi.Project = project
	projectApi.ProjectID = project.ID
	//convert result to base_64
	projectApi.Result = b64.StdEncoding.EncodeToString([]byte(projectApi.Result))

	if err := projectApi.Save(); err != nil {
		response.Json.Object(ctx, *err)
		return
	}

	response.Json.Success(
		ctx,
		nil,
	)

}

func (projectApiController) Edit(ctx *gin.Context) {

	var request requests.ProjectApiEditRequest
	{
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validatorErrors := validation.GetValidationErrors(err)
		response.Json.Error(ctx, http.StatusBadRequest, "ShouldBindJSON Error", validatorErrors)
		return
	}

	//Authenticate user
	user := auth_user.GetAuthUser()

	var project models.Project
	projectTemp, err := project.GetAuthProject(ctx.GetHeader("Project-Token"))
	if err != nil {
		response.Json.Object(ctx, *err)
		return
	}
	project = *projectTemp

	//Check if its user has access to this project
	//TODO : make it middleware
	if project.UserID != user.ID {
		response.Json.Object(ctx, *response.NewUnauthorizedError("You don't have the permission", nil))
		return
	}

	var projectApi models.ProjectApi
	//if err := copier.Copy(&projectApi, &request); err != nil {
	//	response.Json.Error(ctx, http.StatusBadRequest, "Invalid json body to cast", err)
	//	return
	//}
	projectApi.ID = request.ProjectApiID
	projectApi.Project = project
	projectApi.ProjectID = project.ID

	resultTemp := request.Result
	if resultTemp != "" {
		resultTemp = b64.StdEncoding.EncodeToString([]byte(resultTemp))
	}

	if err := projectApi.Update(request.Name, request.Path, request.Type, resultTemp); err != nil {
		response.Json.Object(ctx, *err)
		return
	}

	go func() {
		//insert result to redis
		_, err2 := cache.Cache.Get(context.Background(), fmt.Sprintf("%d:%s", projectApi.ProjectID, projectApi.Path), nil)
		if err2 == nil {
			if err := cache.Cache.Save(context.Background(), time.Hour*24, fmt.Sprintf("%d:%s", projectApi.ProjectID, projectApi.Path), projectApi.Result); err != nil {
				logrus.Error("Error in redis saving", err)
			}
		}
	}()

	response.Json.Success(
		ctx,
		nil,
	)

}

func (projectApiController) Show(ctx *gin.Context) {

	var project models.Project

	projectID, err := project.GetAuthProjectId(ctx.GetHeader("Project-Token"))
	if err != nil {
		response.Json.Object(ctx, *err)
		return
	}
	var projectApi models.ProjectApi
	path := ctx.Param("path")
	if path == "" {
		response.Json.Error(ctx, http.StatusBadRequest, "Parameter path doesn't exist", nil)
		return
	}
	projectApi.Path = path
	projectApi.ProjectID = projectID

	//get result from redis or db
	redisResult, err2 := cache.Cache.Get(context.Background(), fmt.Sprintf("%d:%s", projectApi.ProjectID, projectApi.Path), nil)
	if err2 == nil {
		projectApi.Result = fmt.Sprintf("%s", redisResult)
	} else {
		if err := projectApi.FindWithPath(); err != nil {
			response.Json.Object(ctx, *err)
			return
		}
	}

	result, err2 := b64.StdEncoding.DecodeString(projectApi.Result)
	if err2 != nil {
		response.Json.Error(ctx, http.StatusBadRequest, "Error in decode", err2)
		return
	}

	var jsonResult interface{}

	dataType := "json"
	if err3 := json.Unmarshal(result, &jsonResult); err3 != nil {
		jsonResult = string(result)
		dataType = "string"
	}

	go func() {
		//insert result to redis
		if err := cache.Cache.Save(context.Background(), time.Hour*24, fmt.Sprintf("%d:%s", projectApi.ProjectID, projectApi.Path), projectApi.Result); err != nil {
			logrus.Error("Error in redis saving", err)
		}
	}()

	if dataType == "string" {

		ctx.String(http.StatusOK, string(result))
		return
	}

	ctx.JSON(http.StatusOK, jsonResult)
	return

}
