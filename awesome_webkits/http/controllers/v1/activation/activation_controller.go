package auth

import (
	"awesome_webkits/database/models"
	"awesome_webkits/http/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type activationControllerInterface interface {
	Email(ctx *gin.Context)
	Password(ctx *gin.Context)
}

type activationController struct{}

var ActivationController activationControllerInterface = &activationController{}

func (activationController) Email(ctx *gin.Context) {
	token := ctx.Param("token")
	if token == "" {
		response.Json.Error(ctx, http.StatusBadRequest, "Parameter token doesn't exist", nil)
		return
	}

	emailModel := models.Email{}
	emailModel.Token = token
	emailModel.Type = models.EmailTypeRegister
	emailModel.Status = models.EmailStatusSent
	emailErr := emailModel.FindWithToken()
	if emailErr != nil {
		response.Json.Object(ctx, *emailErr)
		return
	}

	if time.Now().Unix() > emailModel.ExpireAt.Unix() {
		response.Json.Error(ctx, http.StatusBadRequest, "Token was expired", nil)
		return
	}

	userModel := models.User{}
	userModel.Email = emailModel.Email
	userErr := userModel.FindWithEmail()
	if userErr != nil {
		response.Json.Object(ctx, *userErr)
		return
	}

	verifyErr := userModel.Verified()
	if verifyErr != nil {
		response.Json.Object(ctx, *verifyErr)
		return
	}

	statusErr := models.UpdateEmailStatus(emailModel.Email, models.EmailTypeRegister, models.EmailStatusUsed)
	if statusErr != nil {
		response.Json.Error(ctx, http.StatusBadRequest, "Error in update status", emailErr)
		return
	}

	response.Json.Success(ctx, "Email successfully verified")
}

func (activationController) Password(ctx *gin.Context) {
	//todo
	response.Json.Simple(ctx, http.StatusNotImplemented)
}
