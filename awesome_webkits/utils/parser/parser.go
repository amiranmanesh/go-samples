package parser

import (
	"awesome_webkits/http/response"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func ReadBearerToken(tokenHeader string) (string, *response.Data) {
	if !strings.Contains(tokenHeader, "Bearer ") {
		return "", response.NewUnauthorizedError("Bearer doesn't exist.", nil)
	}

	token := strings.Replace(tokenHeader, "Bearer", "", 1)
	token = strings.TrimSpace(token)

	return token, nil
}

func GetIntegerParam(param string) (uint, *response.Data) {

	res, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return 0, response.NewBadRequestErr("invalid param", err)
	}

	return uint(res), nil
}

func SwapJson(out interface{}, in interface{}) *response.Data {

	marshalledJson, err := json.Marshal(in)
	if err != nil {
		logrus.Error(err.Error(), err)
		return response.NewInternalServerError("Error in SwapJson", err.Error())
	}

	_ = json.Unmarshal(marshalledJson, &out)
	return nil
}
