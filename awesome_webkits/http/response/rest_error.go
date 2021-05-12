package response

import (
	"fmt"
	"net/http"
)

func NewBadRequestErr(message interface{}, error interface{}) *Data {
	data := GetSimpleResponse(http.StatusBadRequest)
	if message != nil {
		data.Message = fmt.Sprintf("%v", message)
	}
	if error != nil {
		data.Error = error
	}
	return data
}

func NewUnauthorizedError(message interface{}, error interface{}) *Data {
	data := GetSimpleResponse(http.StatusUnauthorized)
	if message != nil {
		data.Message = fmt.Sprintf("%v", message)
	}
	if error != nil {
		data.Error = error
	}
	return data
}

func NewForbiddenError(message interface{}, error interface{}) *Data {
	data := GetSimpleResponse(http.StatusForbidden)
	if message != nil {
		data.Message = fmt.Sprintf("%v", message)
	}
	if error != nil {
		data.Error = error
	}
	return data
}

func NewNotFoundRequestErr(message interface{}, error interface{}) *Data {
	data := GetSimpleResponse(http.StatusNotFound)
	if message != nil {
		data.Message = fmt.Sprintf("%v", message)
	}
	if error != nil {
		data.Error = error
	}
	return data
}
func NewInternalServerError(message interface{}, error interface{}) *Data {
	data := GetSimpleResponse(http.StatusInternalServerError)
	if message != nil {
		data.Message = fmt.Sprintf("%v", message)
	}
	if error != nil {
		data.Error = error
	}
	return data
}
