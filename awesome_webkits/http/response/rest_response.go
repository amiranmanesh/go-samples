package response

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type iJson interface {
	Simple(c *gin.Context, status int)
	Success(c *gin.Context, meta interface{})
	Object(c *gin.Context, response Data)
	Full(c *gin.Context, status int, success bool, message string, meta interface{})
	Error(c *gin.Context, status int, message string, error interface{})
}
type jsonResponse struct{}

var Json iJson = &jsonResponse{}

type Data struct {
	Status  int         `json:"-"`
	Success bool        `json:"success" default:"false"`
	Message string      `json:"message"`
	Meta    interface{} `json:"meta"`
	Error   interface{} `json:"error"`
}

func (jsonResponse) Simple(c *gin.Context, status int) {
	response := GetSimpleResponse(status)
	log("RespondJson Simple", *response)
	c.Header("Content-Type", "application/json")
	c.JSON(status, response)
}

func (jsonResponse) Success(c *gin.Context, meta interface{}) {
	status := 200
	response := GetSimpleResponse(status)
	if meta != nil {
		response.Meta = meta
	}
	log("RespondJson Success", *response)
	c.Header("Content-Type", "application/json")
	c.JSON(status, response)
}

func (jsonResponse) Object(c *gin.Context, response Data) {
	log("RespondJson Object", response)
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)
}

func (jsonResponse) Full(c *gin.Context, status int, success bool, message string, meta interface{}) {
	response := GetSimpleResponse(status)
	response.Success = success
	if message != "" {
		response.Message = message
	}
	response.Meta = meta
	log("RespondJson Full", *response)
	c.Header("Content-Type", "application/json")
	c.JSON(status, response)
}

func (jsonResponse) Error(c *gin.Context, status int, message string, error interface{}) {
	response := GetSimpleResponse(status)
	if message != "" {
		response.Message = message
	}
	response.Error = error
	log("RespondJson Error", *response)
	c.Header("Content-Type", "application/json")
	c.JSON(status, response)
}

func GetSimpleResponse(status int) *Data {
	switch status {
	case 200:
		{
			return &Data{
				Status:  200,
				Success: true,
				Message: "OK",
			}
		}
	case 201:
		{
			return &Data{
				Status:  201,
				Success: true,
				Message: "CREATED",
			}
		}
	case 202:
		{
			return &Data{
				Status:  202,
				Success: true,
				Message: "ACCEPTED",
			}
		}
	case 204:
		{
			return &Data{
				Status:  204,
				Success: true,
				Message: "NO CONTENT",
			}
		}
	case 400:
		{
			return &Data{
				Status:  400,
				Success: false,
				Message: "INVALID REQUEST",
			}
		}
	case 401:
		{
			return &Data{
				Status:  401,
				Success: false,
				Message: "UNAUTHORIZED",
			}
		}
	case 404:
		{
			return &Data{
				Status:  404,
				Success: false,
				Message: "NOT FOUND",
			}
		}
	case 500:
		{
			return &Data{
				Status:  500,
				Success: false,
				Message: "INTERNAL SERVER ERROR",
			}
		}
	case 501:
		{
			return &Data{
				Status:  501,
				Success: false,
				Message: "NOT IMPLEMENTED",
			}
		}
	default:
		{
			return &Data{
				Status:  500,
				Success: false,
				Message: "INTERNAL SERVER ERROR",
			}
		}
	}
}

func log(TAG string, response Data) {
	logrus.WithFields(logrus.Fields{
		"status": response.Status,
		"response": logrus.Fields{
			"Success": response.Success,
			"Message": response.Message,
			"Meta":    response.Meta,
			"Error":   response.Error,
		},
	}).Info(TAG)
}
