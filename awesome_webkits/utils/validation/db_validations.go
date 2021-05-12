package validation

import (
	"awesome_webkits/database/utils"
	"github.com/go-playground/validator/v10"
	"strings"
)

var dbUnique validator.Func = func(fl validator.FieldLevel) bool {

	//fmt.Println(fl.Field().String(), fl.Param())
	value := fl.Field().String()
	params := fl.Param()

	params = strings.ReplaceAll(params, "'", "")
	/**
	explodes[0] = table
	explodes[0] = field
	*/
	explodes := strings.Split(params, ":")

	return !utils.DbUtils.IsExists(explodes[0], explodes[1], value)
}

var dbExists validator.Func = func(fl validator.FieldLevel) bool {

	//fmt.Println(fl.Field().String(), fl.Param())
	value := fl.Field().String()
	params := fl.Param()

	params = strings.ReplaceAll(params, "'", "")
	/**
	explodes[0] = table
	explodes[0] = field
	*/
	explodes := strings.Split(params, ":")

	return utils.DbUtils.IsExists(explodes[0], explodes[1], value)
}
