package auth

import (
	"awesome_webkits/database/models"
	"awesome_webkits/http/requests/v1"
	"awesome_webkits/http/response"
	"awesome_webkits/utils/validation"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
)

type authControllerInterface interface {
	Signup(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authController struct{}

var AuthController authControllerInterface = &authController{}

func (authController) Signup(ctx *gin.Context) {

	var request requests.UserSignupRequest
	{
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		validatorErrors := validation.GetValidationErrors(err)
		response.Json.Error(ctx, http.StatusBadRequest, "ShouldBindJSON Error", validatorErrors)
		return // exit on first error
	}

	var user models.User
	if err := copier.Copy(&user, &request); err != nil {
		response.Json.Error(ctx, http.StatusBadRequest, "Invalid json body to cast", err)
		return
	}
	var email models.Email
	email.Email = user.Email

	token, err := user.Save()
	if err != nil {
		response.Json.Object(ctx, *err)
		return
	}

	if err := email.Register(); err != nil {
		response.Json.Object(ctx, *err)
		return
	}

	response.Json.Success(
		ctx,
		gin.H{
			"user_token": token,
		},
	)

}

func (authController) Login(ctx *gin.Context) {
	var request requests.UserLoginRequest
	{
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validatorErrors := validation.GetValidationErrors(err)
		response.Json.Error(ctx, http.StatusBadRequest, "ShouldBindJSON Error", validatorErrors)
		return
	}

	user := models.User{}
	token, err := user.Login(request.Email, request.Password)
	if err != nil {
		response.Json.Object(ctx, *err)
		return
	}

	response.Json.Success(
		ctx,
		gin.H{
			"user_token": token,
		},
	)
}
